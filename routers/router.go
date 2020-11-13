package routers

import (
	"github.com/astaxie/beego"
	"github.com/sathishkumar-manogaran/vendor-management-service/controllers"
)

func init() {
	beego.Router("/", &controllers.AppController{}, "get:Status")
	beego.Router("/booking", &controllers.BookingController{}, "post:GetBookingVendorDetails")
}
