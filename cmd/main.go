package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Ckala62rus/go/domain"
	"github.com/Ckala62rus/go/internal/handler"
	"github.com/Ckala62rus/go/internal/repositories"
	"github.com/Ckala62rus/go/internal/services"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// layout := "02:01:2006"
	// startTime := time.Now()

	// ch := make(chan bool)
	// ch2 := make(chan bool)

	// flag := false

	// go func(){
	// 	message := <- ch2

	// 	time.Sleep(5 * time.Second)

	// 	fmt.Println("ch2 отработал")
	// 	fmt.Println(message)
	// 	ch2 <- true
	// }()

	// ch2 <- false

	// go func(){
	// 	message := <- ch

	// 	time.Sleep(4 * time.Second)

	// 	fmt.Println("ch отработал")
	// 	fmt.Println(message)

	// 	ch <- true
	// }()

	// ch <- false

	// for i:= 0; i < 3; i++ {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("test test")
	// 	flag = true
	// }

	// outChan := <-ch
	// outChan2 := <-ch2

	// if flag && outChan && outChan2 {
	// 	fmt.Println(outChan)
	// 	fmt.Println(outChan2)
	// 	fmt.Println("обе программы отработали")
	// }

	// startTimeEnd := time.Now()

	// fmt.Println(startTimeEnd.Sub(startTime))

	// return

	// start time execute
	// startTime := time.Now()

	dir, _ := os.Getwd()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(dir + "\\cmd\\.env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

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
	for _, model := range domain.RegisterModel() {
		err := db.Debug().AutoMigrate(model.Model)

		if err != nil {
			logrus.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully!")
}
