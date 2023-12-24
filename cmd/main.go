package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swim0000/go_crud/handler"
)

func main() {
	app := echo.New()

	inventoryHandler := handler.InventoryHandler{}
	app.GET("/inventory", inventoryHandler.HandleInventoryShow)

	app.Start(":3000")
}
