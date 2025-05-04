package cart

import "github.com/labstack/echo/v4"

type Handler struct {
	cartService Service
}

func NewCartHandler(service *Service) *Handler {
	return &Handler{
		cartService: *service,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("carts/current-cart", h.getCurrentCart)
	e.POST("carts/products/:product", h.addProduct)
	e.DELETE("carts/products/:product", h.removeProduct)
	e.DELETE("carts/products", h.clearCart)
}

func (h *Handler) getCurrentCart(ctx echo.Context) error {
	cart, err := h.cartService.getCurrentCart(ctx, ctx.QueryParam("userId"))
	if err != nil {
		return echo.NewHTTPError(500, "unknown error")
	}

	return ctx.JSON(200, cart)
}

func (h *Handler) addProduct(ctx echo.Context) error {
	cart, err := h.cartService.addProduct(ctx, ctx.QueryParam("userId"), ctx.Param("product"))
	if err != nil {
		return echo.NewHTTPError(500, "unknown error")
	}

	return ctx.JSON(200, cart)
}

func (h *Handler) removeProduct(ctx echo.Context) error {
	cart, err := h.cartService.removeProduct(ctx, ctx.QueryParam("userId"), ctx.Param("product"))
	if err != nil {
		return echo.NewHTTPError(500, "unknown error")
	}

	return ctx.JSON(200, cart)
}

func (h *Handler) clearCart(ctx echo.Context) error {
	cart, err := h.cartService.clearCart(ctx, ctx.QueryParam("userId"))
	if err != nil {
		return echo.NewHTTPError(500, "unknown error")
	}

	return ctx.JSON(200, cart)
}
