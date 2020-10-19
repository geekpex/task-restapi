package main

import (
	"fmt"
	"net/http"
	"task-restapi/data"
	"task-restapi/storage"
)

const (
	pathOrders       = "/customers/{customerID}/orders"
	pathOrdersSingle = "/customers/{customerID}/orders/{orderID}"
)

func addOrderLocation(w http.ResponseWriter, customerID, orderID int64) {
	w.Header().Set("Location", fmt.Sprintf("/customers/%d/orders/%d", customerID, orderID))
}

func getIDs(r *http.Request) (customerID, orderID int64) {
	customerID, _ = getInt64(r, "customerID")
	orderID, _ = getInt64(r, "orderID")
	return
}

func ordersGetAll(w http.ResponseWriter, r *http.Request) {
	customerID, ok := getInt64(r, "customerID")
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, err := storage.LoadAllCustomerOrders(customerID)
	if err != nil {
		// Because only storage.ErrCustomerNotExist exists we can default to that errors response
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if resp == nil {
		resp = make([]data.Order, 0)
	}

	jsonResponse(w, resp, http.StatusOK)
}

func ordersPost(w http.ResponseWriter, r *http.Request) {
	customerID, ok := getInt64(r, "customerID")
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var order data.Order
	err := jsonRequest(r, &order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err = storage.CreateCustomerOrder(customerID, order)
	if err != nil {
		// Because only storage.ErrCustomerNotExist exists we can default to that errors response
		w.WriteHeader(http.StatusNotFound)
		return
	}

	addOrderLocation(w, customerID, order.ID)
	jsonResponse(w, order, http.StatusCreated)
}

func ordersGetSingle(w http.ResponseWriter, r *http.Request) {
	customerID, orderID := getIDs(r)
	if customerID == 0 || orderID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	order, ok := storage.LoadCustomerOrder(customerID, orderID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse(w, order, http.StatusOK)
}

func ordersCancel(w http.ResponseWriter, r *http.Request) {
	customerID, orderID := getIDs(r)
	if customerID == 0 || orderID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	order, ok, err := storage.CancelCustomerOrder(customerID, orderID)
	if err != nil || !ok {
		// Because only storage.ErrCustomerNotExist exists we can default to that errors response
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse(w, order, http.StatusOK)
}
