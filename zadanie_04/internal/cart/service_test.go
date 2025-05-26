package cart

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const (
	CustomerOne = "customer-1"
	ProductOne  = "product-1"
	ProductTwo  = "product-2"
	CartOne     = "cart-1"
)

type MockDao struct {
	carts map[string]Cart
}

func newMockDao() *MockDao {
	return &MockDao{
		carts: make(map[string]Cart),
	}
}

func (m *MockDao) getNewestByCustomerId(customerId string) (Cart, bool, error) {
	for _, cart := range m.carts {
		if cart.CustomerId == customerId {
			return cart, true, nil
		}
	}
	return Cart{}, false, nil
}

func (m *MockDao) create(customerId string) (Cart, error) {
	cart := Cart{
		Id:         "test-cart-id",
		Products:   []string{},
		CustomerId: customerId,
	}
	m.carts[cart.Id] = cart
	return cart, nil
}

func (m *MockDao) updateProducts(cartId string, productsIds []string) (Cart, error) {
	cart, exists := m.carts[cartId]
	if !exists {
		return Cart{}, nil
	}
	cart.Products = productsIds
	m.carts[cartId] = cart
	return cart, nil
}

func TestGetCurrentCart(t *testing.T) {
	mockDao := newMockDao()
	service := NewCartService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	tests := []struct {
		name          string
		customerId    string
		setupMock     func()
		expectedError bool
	}{
		{
			name:       "Existing cart",
			customerId: CustomerOne,
			setupMock: func() {
				mockDao.carts[CartOne] = Cart{
					Id:         CartOne,
					Products:   []string{ProductOne},
					CustomerId: CustomerOne,
				}
			},
			expectedError: false,
		},
		{
			name:          "New cart",
			customerId:    "customer-2",
			setupMock:     func() { /*mock function*/ },
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			cart, err := service.getCurrentCart(ctx, tt.customerId)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Empty(t, cart)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.customerId, cart.CustomerId)
			}
		})
	}
}

func TestAddProduct(t *testing.T) {
	mockDao := newMockDao()
	service := NewCartService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	mockDao.carts[CartOne] = Cart{
		Id:         CartOne,
		Products:   []string{ProductOne},
		CustomerId: CustomerOne,
	}

	cart, err := service.addProduct(ctx, CustomerOne, ProductTwo)

	assert.NoError(t, err)
	assert.Len(t, cart.Products, 2)
	assert.Contains(t, cart.Products, ProductOne)
	assert.Contains(t, cart.Products, ProductTwo)
}

func TestRemoveProduct(t *testing.T) {
	mockDao := newMockDao()
	service := NewCartService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	mockDao.carts[CartOne] = Cart{
		Id:         CartOne,
		Products:   []string{ProductOne, ProductTwo},
		CustomerId: CustomerOne,
	}

	cart, err := service.removeProduct(ctx, CustomerOne, ProductOne)

	assert.NoError(t, err)
	assert.Len(t, cart.Products, 1)
	assert.Contains(t, cart.Products, ProductTwo)
	assert.NotContains(t, cart.Products, ProductOne)
}

func TestClearCart(t *testing.T) {
	mockDao := newMockDao()
	service := NewCartService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	mockDao.carts[CartOne] = Cart{
		Id:         CartOne,
		Products:   []string{ProductOne, ProductTwo},
		CustomerId: CustomerOne,
	}

	cart, err := service.clearCart(ctx, CustomerOne)

	assert.NoError(t, err)
	assert.Empty(t, cart.Products)
}
