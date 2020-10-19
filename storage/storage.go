package storage

import (
	"sync"
	"task-restapi/data"
)

var (
	dataStore = new(sync.Map)

	customerIDs      = new(data.IDs)
	customerIDSerial = new(data.Serial)

	orderIDSerial = new(data.Serial)
)

type (
	dataIDCustomer          int64
	dataIDCustomerOrders    int64
	dataIDCustomerOrderPair struct {
		CustomerID int64
		OrderID    int64
	}
)

// Reset resets data in datastore
func Reset() {
	dataStore = new(sync.Map)

	customerIDs = new(data.IDs)
	customerIDSerial = new(data.Serial)

	orderIDSerial = new(data.Serial)
}
