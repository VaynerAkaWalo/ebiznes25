package main

import (
	"github.com/VaynerAkaWalo/ebiznes25/zadanie_04/internal/products"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	productsDao := products.NewMemoryDao()
	productsService := products.NewProductService(&productsDao)

	productsHandler := products.NewProductsHandler(productsService)
	productsHandler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
