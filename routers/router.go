package routers

import (
	"github.com/astaxie/beego"
	"github.com/sathishkumar-manogaran/vendor-management-service/controllers"
)

func init() {
	beego.Router("/", &controllers.VendorController{}, "get:Status")
	beego.Router("/vendor", &controllers.VendorController{}, "get:GetVendors")
}
