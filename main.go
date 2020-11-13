package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sathishkumar-manogaran/vendor-management-service/database"
	_ "github.com/sathishkumar-manogaran/vendor-management-service/routers"
	"log"
)

func main() {

	// Initiating DB Connections and Related tables
	createDatabaseConnection()

	// Starting the MVC portion
	beego.Run()
}

func createDatabaseConnection() {
	db, err := database.DbConnection()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()
	log.Printf("Successfully connected to database")
	err = database.CreateVendorTable(db)
	if err != nil {
		log.Printf("Create product table failed with error %s", err)
		return
	}
}
