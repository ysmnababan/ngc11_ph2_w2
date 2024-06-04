package main

import (
	"ngc11/config"
	"ngc11/handler"
	"ngc11/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.Connect()

	h := &handler.Repo{DB: db}

	e := echo.New()
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)

	buy := e.Group("")
	buy.Use(middleware.Auth)
	{
		buy.GET("/products", h.GetProducts)
		buy.POST("/transactions", h.BuyProduct)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
