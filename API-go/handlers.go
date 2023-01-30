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
	json.NewEncoder(w).Encode(funcionarios)
}

func getFuncionarioById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range funcionarios {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Funcionario{})
}

func postFuncionario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newFuncionario Funcionario
	_ = json.NewDecoder(r.Body).Decode(&newFuncionario)
	newFuncionario.Id = strconv.Itoa(rand.Intn(10000000))
	funcionarios = append(funcionarios, newFuncionario)
	json.NewEncoder(w).Encode(newFuncionario)
}

func putFuncionario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range funcionarios {
		if item.Id == params["id"] {
			funcionarios = append(funcionarios[:index], funcionarios[index+1:]...)
			var newFuncionario Funcionario
			_ = json.NewDecoder(r.Body).Decode(&newFuncionario)
			newFuncionario.Id = params["id"]
			funcionarios = append(funcionarios, newFuncionario)
			json.NewEncoder(w).Encode(newFuncionario)
			return
		}
	}
	json.NewEncoder(w).Encode(funcionarios)

}

func deleteFuncionario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range funcionarios {
		if item.Id == params["id"] {
			funcionarios = append(funcionarios[:index], funcionarios[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(funcionarios)
}

func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloworld")
}
