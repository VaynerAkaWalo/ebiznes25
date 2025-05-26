package products

import "github.com/labstack/echo/v4"

const (
	IdPath            = "/products/:id"
	IdIsRequiredError = "id is required"
)

type Handler struct {
	productService Service
}

func NewProductsHandler(service *Service) *Handler {
	return &Handler{
		productService: *service,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET(IdPath, h.getById)
	e.GET("/products", h.getAll)

	e.POST("/products", h.create)
	e.PUT(IdPath, h.update)
	e.DELETE(IdPath, h.delete)
}

func (h *Handler) getById(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return echo.NewHTTPError(400, IdIsRequiredError)
	}
	product, err := h.productService.getById(ctx, id)
	if err != nil {
		return err
	}

	return ctx.JSON(200, product)
}

func (h *Handler) getAll(ctx echo.Context) error {
	products, err := h.productService.getAll(ctx)
	if err != nil {
		return err
	}

	if products == nil {
		return ctx.JSON(200, []Product{})
	}

	return ctx.JSON(200, products)
}

func (h *Handler) delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return echo.NewHTTPError(400, IdIsRequiredError)
	}

	found, err := h.productService.delete(ctx, id)
	if err != nil {
		return err
	}
	if !found {
		return echo.NewHTTPError(404, "product not found")
	}

	return ctx.JSON(204, "")
}

func (h *Handler) create(ctx echo.Context) error {
	var req Product
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	product, err := h.productService.create(ctx, req.Name, req.Price)
	if err != nil {
		return err
	}

	return ctx.JSON(201, product)
}

func (h *Handler) update(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return echo.NewHTTPError(400, IdIsRequiredError)
	}

	var req Product
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	product, err := h.productService.update(ctx, id, req.Name, req.Price)
	if err != nil {
		return err
	}

	return ctx.JSON(201, product)
}
