package routers

import (
	"github.com/astaxie/beego"
	"github.com/sathishkumar-manogaran/vendor-management-service/controllers"
	"github.com/sathishkumar-manogaran/vendor-management-service/database"
	"log"
)

func init() {
	// Initiating DB Connections and Related tables
	createDatabaseConnection()

	beego.Router("/", &controllers.AppController{}, "get:Status")
	beego.Router("/booking", &controllers.BookingController{}, "post:GetBookingVendorDetails")
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
