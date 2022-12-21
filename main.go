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

	err := _dbDriver.DBMigrate(db)
	if err != nil {
		panic(err)
	}
	app := routes.New(db)

	log.Fatal(app.Start(util.GetConfig("APP_PORT")))
	// app.Logger.Fatal(app.StartAutoTLS(":443"))
}
