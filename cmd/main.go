package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Ckala62rus/go/domain"
	"github.com/Ckala62rus/go/pkg/handler"
	"github.com/Ckala62rus/go/pkg/repositories"
	"github.com/Ckala62rus/go/pkg/services"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /api/

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	// start time execute
	// startTime := time.Now()

	// s := gocron.NewScheduler(time.UTC)
	// s.Cron("* * * * *").Do(cron)
	// s.StartAsync()

	// cron()

	dir, _ := os.Getwd()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(dir + "\\cmd\\.env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// go utils.ConsumerRabbitMQ()

	// github.com/mattn/go-sqlite3
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open(dir+"\\cmd\\gorm.db"), &gorm.Config{Logger: newLogger})
	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// Run migration initialize
	AutoMigrateInitialize(db)

	rep := repositories.NewUserRepository(db)
	service := services.NewService(rep)
	handlers := handler.NewHandler(service)

	logrus.Info("Server start on localhost:8081")

	srv := new(domain.Server)
	if err := srv.Run("8081", handlers.InitRoutes()); err != nil {
		logrus.Fatal("error occured while running http server: %s", err.Error())
	}

	// stop time execute
	// startTimeEnd := time.Now()
	// fmt.Println(startTimeEnd.Sub(startTime))
}

func AutoMigrateInitialize(db *gorm.DB) {
	// initialize auto migration
	for _, model := range repositories.RegisterModel() {
		err := db.Debug().AutoMigrate(model.Model)

		if err != nil {
			logrus.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully!")
}

func cron() {
	// time := (time.Now()).Format("2022-01-02")
	// fmt.Println("i just do it! \n time: " + time)
	files, err := ioutil.ReadDir("./images/1")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		os.Remove("./images/1" + "/" + file.Name())
		return
	}
}
