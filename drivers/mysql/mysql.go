package mysql_driver

import (
	"fmt"
	"golang/models/model"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_ADDRESS  string
	DB_NAME     string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_ADDRESS,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	return db
}

func DBMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.Instructor{},
		model.Customer{},
		model.CustomerCode{},
		model.Category{},
		model.Course{},
		model.CustomerCourse{},
		model.Favorite{},
		model.Rating{},
		model.Module{},
		model.MediaModule{},
		model.Assignment{},
		model.CustomerAssignment{},
	)

	if err != nil {
		return err
	}

	return nil
}
