package products

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const (
	TestId = "test-id"
)

type MockDao struct {
	products map[string]Product
}

func newMockDao() *MockDao {
	return &MockDao{
		products: make(map[string]Product),
	}
}

func (m *MockDao) getById(id string) (Product, bool, error) {
	product, exists := m.products[id]
	return product, exists, nil
}

func (m *MockDao) getAll() ([]Product, error) {
	products := make([]Product, 0, len(m.products))
	for _, p := range m.products {
		products = append(products, p)
	}
	return products, nil
}

func (m *MockDao) create(name string, price float64) (Product, error) {
	product := Product{
		Id:    TestId,
		Name:  name,
		Price: price,
	}
	m.products[product.Id] = product
	return product, nil
}

func (m *MockDao) update(id string, name string, price float64) (Product, error) {
	product := Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
	m.products[id] = product
	return product, nil
}

func (m *MockDao) delete(id string) (bool, error) {
	_, exists := m.products[id]
	if exists {
		delete(m.products, id)
	}
	return exists, nil
}

func TestGetById(t *testing.T) {
	mockDao := newMockDao()
	service := NewProductService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	tests := []struct {
		name          string
		id            string
		setupMock     func()
		expectedError bool
	}{
		{
			name: "Product exists",
			id:   TestId,
			setupMock: func() {
				mockDao.products[TestId] = Product{
					Id:    TestId,
					Name:  "Test Product",
					Price: 10.99,
				}
			},
			expectedError: false,
		},
		{
			name:          "Product does not exist",
			id:            "non-existent",
			setupMock:     func() { /*mock function*/ },
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			product, err := service.getById(ctx, tt.id)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Empty(t, product)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.id, product.Id)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	mockDao := newMockDao()
	service := NewProductService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	mockDao.products["1"] = Product{Id: "1", Name: "Product 1", Price: 10.99}
	mockDao.products["2"] = Product{Id: "2", Name: "Product 2", Price: 20.99}

	products, err := service.getAll(ctx)

	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 2", products[1].Name)
}

func TestCreate(t *testing.T) {
	mockDao := newMockDao()
	service := NewProductService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	product, err := service.create(ctx, "New Product", 15.99)

	assert.NoError(t, err)
	assert.Equal(t, "New Product", product.Name)
	assert.Equal(t, 15.99, product.Price)
	assert.NotEmpty(t, product.Id)
}

func TestUpdate(t *testing.T) {
	mockDao := newMockDao()
	service := NewProductService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	mockDao.products[TestId] = Product{
		Id:    TestId,
		Name:  "Original Name",
		Price: 10.99,
	}

	product, err := service.update(ctx, TestId, "Updated Name", 20.99)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", product.Name)
	assert.Equal(t, 20.99, product.Price)
	assert.Equal(t, TestId, product.Id)
}

func TestDelete(t *testing.T) {
	mockDao := newMockDao()
	service := NewProductService(mockDao)
	ctx := echo.New().NewContext(nil, nil)

	mockDao.products[TestId] = Product{
		Id:    TestId,
		Name:  "Test Product",
		Price: 10.99,
	}

	tests := []struct {
		name          string
		id            string
		expectedFound bool
	}{
		{
			name:          "Delete existing product",
			id:            TestId,
			expectedFound: true,
		},
		{
			name:          "Delete non-existent product",
			id:            "non-existent",
			expectedFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := service.delete(ctx, tt.id)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedFound, found)
		})
	}
}
