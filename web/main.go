package main

import (
	"errors"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var errBadRoute = errors.New("bad route")

func main()  {
	var repo Repository
	repo, _ = NewRedis("redis://redis:6379")

	var svc RecommenderService
	svc = recommenderService{}.New(repo)

	r := mux.NewRouter()
	rateHandler := httpTransport.NewServer(makeRateEndpoint(svc), decodeRateRequest, encodeResponse)
	r.Handle("/rate", rateHandler).Methods("POST")
	suggestedItemsHandler := httpTransport.NewServer(makeSuggestedItemsEndpoint(svc), decodeSuggestedItemsRequest, encodeResponse)
	r.Handle("/users/{id}/suggestions", suggestedItemsHandler).Methods("GET")
	serverMux := http.NewServeMux()
	serverMux.Handle("/", r)
	http.Handle("/", serverMux)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

