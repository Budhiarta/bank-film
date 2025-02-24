package database

import (
	"fmt"
	"time"

	"github.com/Budhiarta/bank-film-BE/pkg/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(dbHost string, dbPort string, dbUsername string, dbPassword string, dbName string, retries int) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	conString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}

	for {
		db, err = gorm.Open(mysql.Open(conString), gormConfig)
		if err == nil {
			break
		}

		if retries == 0 {
			return nil, err
		}
		retries--

		time.Sleep(5 * time.Second)
	}

	return db, err
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Sharing{},
		&entity.Movie{},
		&entity.ListMovie{},
	)
}
