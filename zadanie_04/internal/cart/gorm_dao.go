package cart

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type DBCart struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey;type:uuid"`
	CustomerId string
	Products   string
}

func (dbC *DBCart) cart() Cart {
	products := strings.Split(dbC.Products, ",")
	if products[0] == "" {
		products = []string{}
	}

	return Cart{
		Id:         dbC.ID.String(),
		CustomerId: dbC.CustomerId,
		Products:   products,
	}
}

type GormDao struct {
	db gorm.DB
}

func NewGormDao(db *gorm.DB) Dao {
	db.AutoMigrate(&DBCart{})

	return &GormDao{
		db: *db,
	}
}

func (dao *GormDao) getNewestByCustomerId(customerId string) (Cart, bool, error) {
	var dbCart DBCart
	res := dao.db.Where("customer_id = ?", customerId).Order("created_at DESC").First(&dbCart)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return Cart{}, false, nil
		}
		return Cart{}, false, res.Error
	}

	return dbCart.cart(), true, nil
}

func (dao *GormDao) create(customerId string) (Cart, error) {
	dbCart := DBCart{
		ID:         uuid.New(),
		CustomerId: customerId,
		Products:   "",
	}

	result := dao.db.Create(&dbCart)
	if result.Error != nil {
		return Cart{}, fmt.Errorf("failed to create cart")
	}

	return dbCart.cart(), nil
}

func (dao *GormDao) updateProducts(cartId string, productIds []string) (Cart, error) {
	dbId, err := uuid.Parse(cartId)
	if err != nil {
		return Cart{}, fmt.Errorf("invalid id")
	}

	var dbCart DBCart
	res := dao.db.First(&dbCart, dbId)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return Cart{}, fmt.Errorf("cart not found")
		}
		return Cart{}, fmt.Errorf("unknown error")
	}

	dbCart.Products = strings.Join(productIds, ",")
	dao.db.Save(dbCart)
	return dbCart.cart(), nil
}
