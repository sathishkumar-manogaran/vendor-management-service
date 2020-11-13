package controllers

import "github.com/astaxie/beego"

type AppController struct {
	beego.Controller
}

func (appController *AppController) Status() {
	appController.Ctx.WriteString("ping pong")
}
