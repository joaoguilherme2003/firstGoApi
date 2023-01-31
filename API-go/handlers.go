package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getFuncionarios(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	allFuncionarios := getAllRedis()
	json.NewEncoder(w).Encode(allFuncionarios)

}

func getFuncionarioById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(findRedis(params["id"]))

}

func postFuncionario(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newFuncionario Funcionario
	_ = json.NewDecoder(r.Body).Decode(&newFuncionario)
	newFuncionario.Id = strconv.Itoa(rand.Intn(10000000))
	postRedis(newFuncionario)
	json.NewEncoder(w).Encode(newFuncionario)

}

func putFuncionario(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	deleteRedis(params["id"])

	var newFuncionario Funcionario
	_ = json.NewDecoder(r.Body).Decode(&newFuncionario)
	newFuncionario.Id = params["id"]
	postRedis(newFuncionario)

	json.NewEncoder(w).Encode(findRedis(params["id"]))

}

func deleteFuncionario(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(deleteRedis(params["id"]))

}

func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloworld")
}
