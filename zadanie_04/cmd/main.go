package main

import (
	"github.com/VaynerAkaWalo/ebiznes25/zadanie_04/internal/cart"
	"github.com/VaynerAkaWalo/ebiznes25/zadanie_04/internal/products"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	db, err := gorm.Open(sqlite.Open("db/database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	productsDao := products.NewGormDao(db)
	productsService := products.NewProductService(&productsDao)

	productsHandler := products.NewProductsHandler(productsService)
	productsHandler.RegisterRoutes(e)

	cartDao := cart.NewGormDao(db)
	cartService := cart.NewCartService(&cartDao)

	cartHandler := cart.NewCartHandler(cartService)
	cartHandler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
