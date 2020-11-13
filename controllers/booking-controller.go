package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/davecgh/go-spew/spew"
	"github.com/sathishkumar-manogaran/vendor-management-service/models"
	"github.com/sathishkumar-manogaran/vendor-management-service/services"
)

type BookingController struct {
	beego.Controller
}

func (bookingController *BookingController) GetBookingVendorDetails() {
	booking := models.Booking{}

	err := json.Unmarshal(bookingController.Ctx.Input.RequestBody, &booking)
	if err != nil {
		bookingController.Ctx.Output.SetStatus(400)
		spew.Dump("******** ERROR ********", err)
		spew.Dump("******** Request Payload ********", bookingController.Ctx.Input.RequestBody)
		bookingController.Ctx.Output.Body([]byte("Error Occurred while processing request"))
	}

	vendors := services.GetBookingVendorDetails(&booking)
	fmt.Println(vendors)

	bookingController.Data["json"] = vendors
	bookingController.ServeJSON()
}
