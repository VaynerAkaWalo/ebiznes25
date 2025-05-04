package cart

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"slices"
)

type Cart struct {
	Id         string   `json:"id"`
	Products   []string `json:"products"`
	CustomerId string   `json:"customerId"`
}

type Dao interface {
	getNewestByCustomerId(customerId string) (Cart, bool, error)
	create(customerId string) (Cart, error)
	updateProducts(cartId string, productsIds []string) (Cart, error)
}

type Service struct {
	cartDao Dao
}

func NewCartService(dao *Dao) *Service {
	return &Service{
		cartDao: *dao,
	}
}

func (s *Service) getCurrentCart(ctx echo.Context, customerId string) (Cart, error) {
	cart, found, err := s.cartDao.getNewestByCustomerId(customerId)
	if err != nil {
		return Cart{}, echo.NewHTTPError(500, "unknown error")
	}
	if !found {
		newCart, err := s.cartDao.create(customerId)
		if err != nil {
			return Cart{}, echo.NewHTTPError(500, "unknown error")
		}
		return newCart, nil
	}

	return cart, nil
}

func (s *Service) addProduct(ctx echo.Context, customerId string, productId string) (Cart, error) {
	cart, err := s.getCurrentCart(ctx, customerId)
	if err != nil {
		return Cart{}, err
	}

	fmt.Println(cart.Products)
	newProducts := append(cart.Products, productId)
	fmt.Println(newProducts)

	cart, err = s.cartDao.updateProducts(cart.Id, newProducts)
	if err != nil {
		return Cart{}, fmt.Errorf("unknown error")
	}

	return cart, nil
}

func (s *Service) removeProduct(ctx echo.Context, customerId string, productId string) (Cart, error) {
	cart, err := s.getCurrentCart(ctx, customerId)
	if err != nil {
		return Cart{}, err
	}

	newProducts := make([]string, len(cart.Products))
	copy(newProducts, cart.Products)
	newProducts = slices.DeleteFunc(newProducts, func(id string) bool {
		return id == productId
	})

	cart, err = s.cartDao.updateProducts(cart.Id, newProducts)
	if err != nil {
		return Cart{}, fmt.Errorf("unknown error")
	}

	return cart, nil
}

func (s *Service) clearCart(ctx echo.Context, customerId string) (Cart, error) {
	cart, err := s.getCurrentCart(ctx, customerId)
	if err != nil {
		return Cart{}, err
	}

	cart, err = s.cartDao.updateProducts(cart.Id, []string{})
	if err != nil {
		return Cart{}, fmt.Errorf("unknown error")
	}

	return cart, nil
}
