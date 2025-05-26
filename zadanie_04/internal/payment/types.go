package payment

type Status struct {
	Status string `json:"status"`
}

type CardPayment struct {
	ProviderId string      `json:"providerId"`
	Card       CardDetails `json:"card"`
	Amount     float64     `json:"amount"`
}

type CardDetails struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	CardNumber int    `json:"cardNumber"`
	ExpireTime string `json:"expireTime"`
	CVV        int    `json:"CVV"`
}

type Test struct {
}
