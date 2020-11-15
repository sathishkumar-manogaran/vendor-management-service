package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/sathishkumar-manogaran/vendor-management-service/routers"
)

func main() {

	// Starting the Vendor Management Service portion
	beego.Run()
}
