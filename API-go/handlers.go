package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func getFuncionarios(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	w.Header().Set("Content-Type", "application/json")
	allFuncionarios := getAllRedis(ctx)
	json.NewEncoder(w).Encode(allFuncionarios)

}

func getFuncionarioById(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(findRedis(ctx, params["id"]))

}

func postFuncionario(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	w.Header().Set("Content-Type", "application/json")
	var newFuncionario Funcionario
	_ = json.NewDecoder(r.Body).Decode(&newFuncionario)
	newFuncionario.Id = strconv.Itoa(rand.Intn(10000000))
	postRedis(ctx, newFuncionario)
	json.NewEncoder(w).Encode(newFuncionario)

}

func putFuncionario(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	deleteRedis(ctx, params["id"])

	var newFuncionario Funcionario
	_ = json.NewDecoder(r.Body).Decode(&newFuncionario)
	newFuncionario.Id = params["id"]
	postRedis(ctx, newFuncionario)

	json.NewEncoder(w).Encode(findRedis(ctx, params["id"]))

}

func deleteFuncionario(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(deleteRedis(ctx, params["id"]))

}

func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloworld")
}
