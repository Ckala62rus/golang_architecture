package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/streadway/amqp"
)

// create folder if not exist
func CreateFolder(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirname, 0755)
		if errDir != nil {
			return errDir
		}
	}
	return nil
}

// string to hash
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// get timestamp
func TimeStamp() string {
	return (time.Now()).Format("20060102150405")
}

func ConsumerRabbitMQ()  {
	// Define RabbitMQ server URL.
    amqpServerURL := "amqp://guest:guest@localhost:5672/"

    // Create a new RabbitMQ connection.
    connectRabbitMQ, err := amqp.Dial(amqpServerURL)
    if err != nil {
        panic(err)
    }
    defer connectRabbitMQ.Close()

    // Opening a channel to our RabbitMQ instance over
    // the connection we have already established.
    channelRabbitMQ, err := connectRabbitMQ.Channel()
    if err != nil {
        panic(err)
    }
    defer channelRabbitMQ.Close()

    // Subscribing to QueueService1 for getting messages.
    messages, err := channelRabbitMQ.Consume(
        "users", // queue name
        "",              // consumer
        true,            // auto-ack
        false,           // exclusive
        false,           // no local
        false,           // no wait
        nil,             // arguments
    )
    if err != nil {
        log.Println(err)
    }

    // Build a welcome message.
    log.Println("Successfully connected to RabbitMQ")
    log.Println("Waiting for messages")

    // Make a channel to receive messages into infinite loop.
    forever := make(chan bool)

    go func() {
        for message := range messages {
            // For example, show received message in a console.
            log.Printf(" > Received message: %s\n", message.Body)
			SendEmailRabbit(message.Body)
        }
    }()

    <-forever
}

func SendEmailRabbit(message []byte){
	// Sender data.
    from := os.Getenv("EMAIL")
    password := os.Getenv("PASSWORD")

    // Receiver email address.
    to := []string{
        "agr.akyla@mail.ru",
    }

    // smtp server configuration.
    smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

    // Message.
    // message1 := []byte("This is a really unimaginative message, I know. Hello!")

	newUserFromLaravel := User{}
	json.Unmarshal(message, &newUserFromLaravel)

    // Authentication.
    auth := smtp.PlainAuth("", from, password, smtpServer.host)

    // Sending email.
    err := smtp.SendMail(smtpServer.Address(), auth, from, to, []byte("In system laraval was create new user with email: " + newUserFromLaravel.Email))
    if err != nil {
		log.Printf(err.Error())
        return
    }
}

type User struct {
	Id int
	Email string 
	Name string
}

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}
