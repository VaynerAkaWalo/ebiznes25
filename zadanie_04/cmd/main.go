package main

import (
	"github.com/VaynerAkaWalo/ebiznes25/zadanie_04/internal/products"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	e := echo.New()

	db, err := gorm.Open(sqlite.Open("db/products.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	productsDao := products.NewGormDao(db)
	productsService := products.NewProductService(&productsDao)

	productsHandler := products.NewProductsHandler(productsService)
	productsHandler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
