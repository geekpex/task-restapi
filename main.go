package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
)

func newRouter() http.Handler {
	r := chi.NewRouter()

	r.Get(pathCustomers, customersGetAll)
	r.Post(pathCustomers, customersPost)
	r.Get(pathCustomerSingle, customersGetSingle)
	r.Put(pathCustomerSingle, customersPut)

	r.Get(pathOrders, ordersGetAll)
	r.Post(pathOrders, ordersPost)
	r.Get(pathOrdersSingle, ordersGetSingle)
	r.Delete(pathOrdersSingle, ordersCancel)

	return r
}

const (
	listeningPort = 8080
)

func main() {
	fmt.Println("Starting task-restapi...")

	r := newRouter()

	go func() {
		fmt.Printf("Listening on port %d for REST calls...", listeningPort)
		fmt.Println()
		err := http.ListenAndServe(fmt.Sprintf(":%d", listeningPort), r)
		if err != nil {
			fmt.Printf("Failed to start listening on %d port:", listeningPort)
			fmt.Println()
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	fmt.Println("exited")
}
