package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/sathishkumar-manogaran/vendor-management-service/routers"
)

func main() {

	// Starting the MVC portion
	beego.Run()
}
