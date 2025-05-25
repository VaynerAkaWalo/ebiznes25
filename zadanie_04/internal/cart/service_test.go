package cart

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
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

func TestService_GetCurrentCart(t *testing.T) {
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
			customerId: "customer-1",
			setupMock: func() {
				mockDao.carts["cart-1"] = Cart{
					Id:         "cart-1",
					Products:   []string{"product-1"},
					CustomerId: "customer-1",
				}
			},
			expectedError: false,
		},
		{
			name:          "New cart",
			customerId:    "customer-2",
			setupMock:     func() {},
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

func TestService_AddProduct(t *testing.T) {
	mockDao := newMockDao()
	service := NewCartService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	mockDao.carts["cart-1"] = Cart{
		Id:         "cart-1",
		Products:   []string{"product-1"},
		CustomerId: "customer-1",
	}

	cart, err := service.addProduct(ctx, "customer-1", "product-2")

	assert.NoError(t, err)
	assert.Len(t, cart.Products, 2)
	assert.Contains(t, cart.Products, "product-1")
	assert.Contains(t, cart.Products, "product-2")
}

func TestService_RemoveProduct(t *testing.T) {
	mockDao := newMockDao()
	service := NewCartService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	mockDao.carts["cart-1"] = Cart{
		Id:         "cart-1",
		Products:   []string{"product-1", "product-2"},
		CustomerId: "customer-1",
	}

	cart, err := service.removeProduct(ctx, "customer-1", "product-1")

	assert.NoError(t, err)
	assert.Len(t, cart.Products, 1)
	assert.Contains(t, cart.Products, "product-2")
	assert.NotContains(t, cart.Products, "product-1")
}

func TestService_ClearCart(t *testing.T) {
	mockDao := newMockDao()
	service := NewCartService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	mockDao.carts["cart-1"] = Cart{
		Id:         "cart-1",
		Products:   []string{"product-1", "product-2"},
		CustomerId: "customer-1",
	}

	cart, err := service.clearCart(ctx, "customer-1")

	assert.NoError(t, err)
	assert.Empty(t, cart.Products)
}
