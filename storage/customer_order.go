package storage

import (
	"errors"
	"task-restapi/data"
	"time"
)

var (
	// ErrCustomerNotExist ...
	ErrCustomerNotExist = errors.New("customer does not exist")
)

// ExistsCustomerOrder return ok = True if customer's order already exists
func ExistsCustomerOrder(customerID, orderID int64) (ok bool) {
	_, ok = LoadCustomerOrder(customerID, orderID)
	return
}

// LoadCustomerOrder returns stored Customer's Order if found. Otherwise ok is False
func LoadCustomerOrder(customerID, orderID int64) (order data.Order, ok bool) {
	var v interface{}
	v, ok = dataStore.Load(dataIDCustomerOrderPair{
		CustomerID: customerID,
		OrderID:    orderID,
	})
	if !ok {
		return
	}
	order, ok = v.(data.Order)
	return
}

// CreateCustomerOrder stores new order for existing customer. Returns stored order.
func CreateCustomerOrder(customerID int64, order data.Order) (data.Order, error) {
	if !ExistsCustomer(customerID) {
		return data.Order{}, ErrCustomerNotExist
	}
	order.ID = orderIDSerial.Get()
	order.CustomerID = customerID
	order.CreatedAt = time.Now()
	order.Status = data.StatusActive

	dataStore.Store(dataIDCustomerOrderPair{
		CustomerID: customerID,
		OrderID:    order.ID,
	},
		order,
	)

	addOrderIDForCustomer(customerID, order.ID)

	return order, nil
}

func getOrderIDsForCustomer(customerID int64) *data.IDs {
	key := dataIDCustomerOrders(customerID)
	v, ok := dataStore.Load(key)
	var ids *data.IDs
	if ok {
		ids, ok = v.(*data.IDs)
	}
	if !ok {
		ids = new(data.IDs)
	}
	return ids
}

func addOrderIDForCustomer(customerID, orderID int64) {
	ids := getOrderIDsForCustomer(customerID)
	ids.Add(orderID)

	dataStore.Store(dataIDCustomerOrders(customerID), ids)
}

// CancelCustomerOrder tries to cancel customer's order.
// If canceled return ok = True
func CancelCustomerOrder(customerID, orderID int64) (order data.Order, ok bool, err error) {
	if !ExistsCustomer(customerID) {
		err = ErrCustomerNotExist
		return
	}

	order, ok = LoadCustomerOrder(customerID, orderID)
	if !ok {
		return
	}

	order.Status = data.StatusCancelled

	dataStore.Store(dataIDCustomerOrderPair{
		CustomerID: customerID,
		OrderID:    order.ID,
	},
		order,
	)

	ok = true
	return
}

// LoadAllCustomerOrders returns all customer's orders
func LoadAllCustomerOrders(customerID int64) ([]data.Order, error) {
	if !ExistsCustomer(customerID) {
		return nil, ErrCustomerNotExist
	}

	ids := getOrderIDsForCustomer(customerID).IDs()

	var orders = make([]data.Order, 0, len(ids))

	var order data.Order
	var ok bool
	for _, id := range ids {
		order, ok = LoadCustomerOrder(customerID, id)
		if !ok {
			continue
		}
		orders = append(orders, order)
	}

	return orders, nil
}
