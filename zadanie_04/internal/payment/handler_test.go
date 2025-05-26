package payment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const PaymentPath = "/payment/card"

func TestCardPayment(t *testing.T) {
	handler := NewPaymentHandler()
	e := echo.New()

	tests := []struct {
		name           string
		request        CardPayment
		expectedStatus int
		expectedBody   Status
	}{
		{
			name: "Valid card payment",
			request: CardPayment{
				ProviderId: "visa",
				Card: CardDetails{
					FirstName:  "John",
					LastName:   "Doe",
					CardNumber: 4111111111111111,
					ExpireTime: "12/25",
					CVV:        123,
				},
				Amount: 100.50,
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   Status{Status: "Completed"},
		},
		{
			name: "Invalid card number",
			request: CardPayment{
				ProviderId: "visa",
				Card: CardDetails{
					FirstName:  "John",
					LastName:   "Doe",
					CardNumber: 123, // Too short
					ExpireTime: "12/25",
					CVV:        123,
				},
				Amount: 100.50,
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   Status{Status: "Completed"},
		},
		{
			name: "Invalid CVV",
			request: CardPayment{
				ProviderId: "visa",
				Card: CardDetails{
					FirstName:  "John",
					LastName:   "Doe",
					CardNumber: 4111111111111111,
					ExpireTime: "12/25",
					CVV:        12, // Too short
				},
				Amount: 100.50,
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   Status{Status: "Completed"},
		},
		{
			name: "Invalid expire time",
			request: CardPayment{
				ProviderId: "visa",
				Card: CardDetails{
					FirstName:  "John",
					LastName:   "Doe",
					CardNumber: 4111111111111111,
					ExpireTime: "13/25", // Invalid month
					CVV:        123,
				},
				Amount: 100.50,
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   Status{Status: "Completed"},
		},
		{
			name: "Zero amount",
			request: CardPayment{
				ProviderId: "visa",
				Card: CardDetails{
					FirstName:  "John",
					LastName:   "Doe",
					CardNumber: 4111111111111111,
					ExpireTime: "12/25",
					CVV:        123,
				},
				Amount: 0,
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   Status{Status: "Completed"},
		},
		{
			name: "Negative amount",
			request: CardPayment{
				ProviderId: "visa",
				Card: CardDetails{
					FirstName:  "John",
					LastName:   "Doe",
					CardNumber: 4111111111111111,
					ExpireTime: "12/25",
					CVV:        123,
				},
				Amount: -100.50,
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   Status{Status: "Completed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.request)
			req := httptest.NewRequest(http.MethodPost, PaymentPath, bytes.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.cardPayment(c)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			var response Status
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, response)
		})
	}
}

func TestCardPaymentInvalidJson(t *testing.T) {
	handler := NewPaymentHandler()
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, PaymentPath, bytes.NewReader([]byte("invalid json")))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.cardPayment(c)

	assert.Error(t, err)
	he, ok := err.(*echo.HTTPError)
	assert.True(t, ok)
	assert.Equal(t, http.StatusBadRequest, he.Code)
}

func TestCardPaymentMissingContentType(t *testing.T) {
	handler := NewPaymentHandler()
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, PaymentPath, bytes.NewReader([]byte("{}")))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.cardPayment(c)

	assert.Error(t, err)
	he, ok := err.(*echo.HTTPError)
	assert.True(t, ok)
	assert.Equal(t, http.StatusBadRequest, he.Code)
}
