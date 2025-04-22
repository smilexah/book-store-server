package controllers

import (
	"book-store-server/services"
)

type AppController struct {
	AppService *services.AppService
}

func NewController(appService *services.AppService) *AppController {
	return &AppController{AppService: appService}
}
