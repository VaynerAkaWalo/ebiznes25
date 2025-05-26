package products

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBProduct struct {
	gorm.Model
	ID    uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}

func (dbP *DBProduct) product() Product {
	return Product{
		Id:    dbP.ID.String(),
		Name:  dbP.Name,
		Price: dbP.Price,
	}
}

type GormDao struct {
	db gorm.DB
}

func NewGormDao(db *gorm.DB) Dao {
	err := db.AutoMigrate(&DBProduct{})
	if err != nil {
		return nil
	}

	return &GormDao{
		db: *db,
	}
}

func (dao *GormDao) getById(id string) (Product, bool, error) {
	dbId, err := uuid.Parse(id)
	if err != nil {
		return Product{}, false, fmt.Errorf("invalid id")
	}

	var dbProduct DBProduct
	res := dao.db.First(&dbProduct, dbId)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return Product{}, false, nil
		}
		return Product{}, false, fmt.Errorf("unknown error")
	}

	return dbProduct.product(), true, nil
}

func (dao *GormDao) getAll() ([]Product, error) {
	var dbProducts []DBProduct
	res := dao.db.Find(&dbProducts)
	if res.Error != nil {
		return nil, fmt.Errorf("unknown error")
	}

	var products []Product
	for _, dbProduct := range dbProducts {
		products = append(products, dbProduct.product())
	}

	return products, nil
}

func (dao *GormDao) create(name string, price float64) (Product, error) {
	dbProduct := DBProduct{
		ID:    uuid.New(),
		Name:  name,
		Price: price,
	}

	result := dao.db.Create(&dbProduct)
	if result.Error != nil {
		return Product{}, fmt.Errorf("failed to create product")
	}

	return dbProduct.product(), nil
}

func (dao *GormDao) update(id string, name string, price float64) (Product, error) {
	dbId, err := uuid.Parse(id)
	if err != nil {
		return Product{}, fmt.Errorf("invalid id")
	}

	var dbProduct DBProduct
	res := dao.db.First(&dbProduct, dbId)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return Product{}, fmt.Errorf("product not found")
		}
		return Product{}, fmt.Errorf("unknown error")
	}

	dbProduct.Name = name
	dbProduct.Price = price

	dao.db.Save(&dbProduct)
	return dbProduct.product(), nil
}

func (dao *GormDao) delete(id string) (bool, error) {
	dbId, err := uuid.Parse(id)
	if err != nil {
		return false, fmt.Errorf("invalid id")
	}

	res := dao.db.Delete(&DBProduct{}, dbId)
	if res.Error != nil {
		return false, fmt.Errorf("unknown error")
	}

	if res.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
