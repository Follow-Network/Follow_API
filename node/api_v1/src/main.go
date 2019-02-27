package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./bundles/state"
	"./core"
)

func initBundles() []core.Bundle {
	return []core.Bundle{state.NewStateBundle()}
}

func main() {

	router := mux.NewRouter()
	apiV1 := router.PathPrefix("/api/v1").Subrouter()

	bundles := initBundles()

	for _, bundle := range bundles {
		for _, route := range bundle.GetRoutes() {
			apiV1.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	log.Fatal(http.ListenAndServe(":5000", router))
}
