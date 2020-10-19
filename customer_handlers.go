package main

import (
	"fmt"
	"net/http"
	"task-restapi/data"
	"task-restapi/storage"
)

const (
	pathCustomers      = "/customers"
	pathCustomerSingle = "/customers/{customerID}"
)

func addCustomerLocation(w http.ResponseWriter, customerID int64) {
	w.Header().Set("Location", fmt.Sprintf("/customers/%d", customerID))
}

func customersGetAll(w http.ResponseWriter, r *http.Request) {
	var resp = storage.LoadAllCustomers()
	if resp == nil {
		resp = make([]data.Customer, 0)
	}

	jsonResponse(w, resp, http.StatusOK)
}

func customersPost(w http.ResponseWriter, r *http.Request) {
	var customer data.Customer
	err := jsonRequest(r, &customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	customer = storage.StoreCustomer(customer)

	addCustomerLocation(w, customer.ID)

	jsonResponse(w, customer, http.StatusCreated)
}

func customersGetSingle(w http.ResponseWriter, r *http.Request) {
	customerID, ok := getInt64(r, "customerID")
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	customer, ok := storage.LoadCustomer(customerID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse(w, customer, http.StatusOK)
}

func customersPut(w http.ResponseWriter, r *http.Request) {
	customerID, ok := getInt64(r, "customerID")
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var customer data.Customer
	err := jsonRequest(r, &customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	customer.ID = customerID

	var statusCode int
	if storage.ExistsCustomer(customerID) {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusCreated
		addCustomerLocation(w, customer.ID)
	}
	customer = storage.StoreCustomer(customer)

	jsonResponse(w, customer, statusCode)
}
