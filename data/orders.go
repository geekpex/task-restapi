package data

import "time"

type OrderStatus string

const (
	StatusActive    = OrderStatus("active")
	StatusCancelled = OrderStatus("cancelled")
)

type Order struct {
	ID         int64 `json:"orderId"`
	CustomerID int64 `json:"customerId"`

	CreatedAt time.Time   `json:"createdAt"`
	Status    OrderStatus `json:"status"`

	Books []Book `json:"books"`
}

type Book struct {
	Name string `json:"name"`
}
