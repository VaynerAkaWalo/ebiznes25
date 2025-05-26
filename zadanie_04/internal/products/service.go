package products

import (
	"github.com/labstack/echo/v4"
)

const UnknownError = "unknown error"

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

type Dao interface {
	getById(string) (Product, bool, error)
	getAll() ([]Product, error)
	create(name string, price float64) (Product, error)
	update(id string, name string, price float64) (Product, error)
	delete(string) (bool, error)
}

type Service struct {
	productDao Dao
}

func NewProductService(dao Dao) *Service {
	return &Service{
		productDao: dao,
	}
}

func (s *Service) getById(ctx echo.Context, id string) (Product, error) {
	product, found, err := s.productDao.getById(id)
	if err != nil {
		return Product{}, echo.NewHTTPError(500, UnknownError, err)
	}

	if !found {
		return Product{}, echo.NewHTTPError(404, "product not found")
	}

	return product, nil
}

func (s *Service) getAll(ctx echo.Context) ([]Product, error) {
	products, err := s.productDao.getAll()
	if err != nil {
		return nil, echo.NewHTTPError(500, UnknownError, err)
	}
	return products, nil
}

func (s *Service) create(ctx echo.Context, name string, price float64) (Product, error) {
	product, err := s.productDao.create(name, price)
	if err != nil {
		return Product{}, echo.NewHTTPError(500, UnknownError, err)
	}

	return product, nil
}

func (s *Service) update(ctx echo.Context, id string, name string, price float64) (Product, error) {
	product, err := s.productDao.update(id, name, price)
	if err != nil {
		return Product{}, echo.NewHTTPError(500, UnknownError, err)
	}
	return product, nil
}

func (s *Service) delete(ctx echo.Context, id string) (bool, error) {
	found, err := s.productDao.delete(id)
	if err != nil {
		return false, echo.NewHTTPError(500, UnknownError, err)
	}

	return found, nil
}
