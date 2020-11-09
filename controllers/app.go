package controllers

import "github.com/astaxie/beego"

type VendorController struct {
	beego.Controller
}

func (vendorController *VendorController) GetVendors() {
	vendorController.Data["vendor"] = "Vendor1"
	vendorController.Data["service"] = "Transport"
	vendorController.Ctx.WriteString("test world")
}

func (vendorController *VendorController) Status() {
	vendorController.Ctx.WriteString("ping pong")
}
