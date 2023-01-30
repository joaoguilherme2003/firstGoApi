package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {

	r := mux.NewRouter()

	r.HandleFunc("/helloworld", getHelloWorld).Methods("GET")
	r.HandleFunc("/funcionarios", getFuncionarios).Methods("GET")
	r.HandleFunc("/funcionarios/{id}", getFuncionarioById).Methods("GET")
	r.HandleFunc("/funcionarios", postFuncionario).Methods("POST")
	r.HandleFunc("/funcionarios/{id}", putFuncionario).Methods("PUT")
	r.HandleFunc("/funcionarios/{id}", deleteFuncionario).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
