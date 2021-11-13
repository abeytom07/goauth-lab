package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	/*
		"github.com/Traceableai/goagent"
		"github.com/Traceableai/goagent/config"
		"github.com/Traceableai/goagent/instrumentation/net/traceablehttp"
	*/)

func main() {

	/*
		cfg := config.Load()

		shutdown := goagent.Init(cfg)
		defer shutdown()
	*/

	router := mux.NewRouter()

	/*
		router.Handle("/login", traceablehttp.NewHandler(http.HandlerFunc(Login), "/login"))
		router.Handle("/refresh", traceablehttp.NewHandler(http.HandlerFunc(Refresh), "/refresh"))
		router.Handle("/test/{id}", traceablehttp.NewHandler(isAuthorized(test), "/test/{id}")).Methods("GET")
		router.Handle("/customer/all", traceablehttp.NewHandler(isAuthorized(customercount), "/customer/all")).Methods("GET")
		router.Handle("/customer/byid/{id}", traceablehttp.NewHandler(isAuthorized(customerbyid), "/customer/byid/{id}")).Methods("GET")
		router.Handle("/crypto/home", traceablehttp.NewHandler(isAuthorized(cryptohome), "/crypto/home")).Methods("GET")
		router.Handle("/crypto/price", traceablehttp.NewHandler(isAuthorized(cryptoprice), "/crypto/price")).Methods("GET")
	*/

	router.HandleFunc("/login", Login).Methods("GET")
	router.HandleFunc("/refresh", Refresh).Methods("GET")
	router.Handle("/test/{id}", isAuthorized(test)).Methods("GET")
	router.Handle("/customer/all", isAuthorized(customercount)).Methods("GET")
	router.Handle("/customer/byid/{id}", isAuthorized(customerbyid)).Methods("GET")
	router.Handle("/crypto/home", isAuthorized(cryptohome)).Methods("GET")
	router.Handle("/crypto/price", isAuthorized(cryptoprice)).Methods("GET")

	//Start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", router))
}
