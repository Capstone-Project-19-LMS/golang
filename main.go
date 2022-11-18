package main

import (
	"golang/app/routes"
	"golang/util"
	"log"

	_dbDriver "golang/drivers/mysql"
)

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_ADDRESS:  util.GetConfig("DB_ADDRESS"),
		DB_NAME:     util.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.DBMigrate(db)

	app := routes.New(db)

	log.Fatal(app.Start(util.GetConfig("APP_PORT")))
}
