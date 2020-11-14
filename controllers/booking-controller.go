package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/davecgh/go-spew/spew"
	"github.com/sathishkumar-manogaran/vendor-management-service/models"
	"github.com/sathishkumar-manogaran/vendor-management-service/services"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

//func init() {
//	validate = validator.New()
//	validate.RegisterValidation("json", isJSON)
//}

type BookingController struct {
	beego.Controller
}

func (bookingController *BookingController) GetBookingVendorDetails() {
	booking := models.Booking{}

	err := json.Unmarshal(bookingController.Ctx.Input.RequestBody, &booking)
	if err != nil {
		bookingController.Ctx.Output.ContentType("application/json")
		bookingController.Ctx.Output.SetStatus(400)
		spew.Dump("******** ERROR ********", err)
		spew.Dump("******** Request Payload ********", bookingController.Ctx.Input.RequestBody)
		bookingController.Ctx.Output.Body([]byte("Error Occurred while processing request"))
	}

	//var vendors interface{}
	if validErrs, vendors := services.GetBookingVendorDetails(&booking); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		bookingController.Ctx.Output.ContentType("application/json")
		bookingController.Ctx.Output.SetStatus(400)
		bookingController.Ctx.Output.IsForbidden()
		bookingController.Ctx.Output.JSON(err, true, true)
	} else {
		fmt.Println(vendors)
		bookingController.Data["json"] = vendors
		bookingController.ServeJSON()
	}

	//vendors := services.GetBookingVendorDetails(&booking)
	//fmt.Println(vendors)

	//bookingController.Data["json"] = vendors
	//bookingController.ServeJSON()
}
