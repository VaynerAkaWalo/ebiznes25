package products

import (
	"fmt"
	"github.com/google/uuid"
	"maps"
	"slices"
	"sync"
)

type MemoryDao struct {
	products map[string]Product
	mutex    *sync.RWMutex
}

func NewMemoryDao() Dao {
	return &MemoryDao{
		products: make(map[string]Product),
		mutex:    &sync.RWMutex{},
	}
}

func (dao *MemoryDao) getById(id string) (Product, bool, error) {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()

	product, found := dao.products[id]
	return product, found, nil
}

func (dao *MemoryDao) getAll() ([]Product, error) {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()

	return slices.Collect(maps.Values(dao.products)), nil
}

func (dao *MemoryDao) create(name string, price float64) (Product, error) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	id := uuid.New()
	product := Product{
		Id:    id.String(),
		Name:  name,
		Price: price,
	}

	dao.products[product.Id] = product
	return product, nil
}

func (dao *MemoryDao) update(id string, name string, price float64) (Product, error) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	product, found := dao.products[id]
	if !found {
		return Product{}, fmt.Errorf("product not found")
	}

	product.Name = name
	product.Price = price

	dao.products[id] = product

	return product, nil
}

func (dao *MemoryDao) delete(id string) (bool, error) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	_, found := dao.products[id]
	if !found {
		return false, nil
	}

	delete(dao.products, id)

	return true, nil
}
