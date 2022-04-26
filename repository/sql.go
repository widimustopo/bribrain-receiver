package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConnectionSQL struct {
	ConnGorm *gorm.DB
}

var ConnectionDB *gorm.DB

func (conn *ConnectionSQL) InitializationDB(retries int) *gorm.DB {
	init := ConnectionSQL{}
	if retries > 1 {
		logrus.Printf("Retrying connect to DB instance, Attempt %v", strconv.Itoa(retries))

		if retries > 5 {
			logrus.Printf("Cannot recovery situation retries > 5 attempt")
			os.Exit(1)
		}
	}

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	sqlDB, err := sql.Open("mysql", connString)

	if err != nil {
		logrus.Printf("error on creating connection sql database %v", err)
		init.InitializationDB(retries + 1)
		return nil
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	conn.ConnGorm, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		logrus.Println("error on creating gorm connection ", err)
		init.InitializationDB(retries + 1)
		return nil
	}

	if err != nil {
		logrus.Printf("error on creating connection database %v", err)

		init.InitializationDB(retries + 1)
		return nil
	}

	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
		})
	conn.ConnGorm.Session(&gorm.Session{Logger: newLogger})

	logrus.Println("database connection successfully")
	ConnectionDB = conn.ConnGorm

	return conn.ConnGorm
}
