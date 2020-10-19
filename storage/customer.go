package storage

import "task-restapi/data"

// ExistsCustomer return ok = True if customer already exists
func ExistsCustomer(customerID int64) (ok bool) {
	_, ok = LoadCustomer(customerID)
	return
}

// LoadCustomer returns stored Customer if found. Otherwise ok is False
func LoadCustomer(customerID int64) (customer data.Customer, ok bool) {
	var v interface{}
	v, ok = dataStore.Load(dataIDCustomer(customerID))
	if !ok {
		return
	}
	customer, ok = v.(data.Customer)
	return
}

// StoreCustomer stores new or modifys existing customer. Returns stored customer.
func StoreCustomer(customer data.Customer) data.Customer {
	if customer.ID == 0 {
		customer.ID = customerIDSerial.Get()
		customerIDs.Add(customer.ID)
	}

	dataStore.Store(dataIDCustomer(customer.ID), customer)

	return customer
}

// LoadAllCustomers returns all existing customers
func LoadAllCustomers() []data.Customer {
	ids := customerIDs.IDs()
	var customers = make([]data.Customer, 0, len(ids))

	var customer data.Customer
	var ok bool
	for _, id := range ids {
		customer, ok = LoadCustomer(id)
		if !ok {
			continue
		}
		customers = append(customers, customer)
	}

	return customers
}
