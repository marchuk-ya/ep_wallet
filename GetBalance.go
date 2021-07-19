package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//GetBalance function return all budget events
//method GET
func GetBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetBalance")

	var Amount, CurrentBalance float64
	var Response []string

	vars := mux.Vars(r)
	id := vars["id"]
	currency := vars["currency"]

	session := GetCassandraSession(KeyspaceName)
	defer session.Close()

	iter := session.Query(
		`SELECT amount FROM wallet WHERE user_id = ? AND currency = ? ALLOW FILTERING`,
		id,
		currency).Iter()

	for iter.Scan(&Amount) {
		CurrentBalance += Amount
	}

	fmt.Println("ID=", id, " Balance=", CurrentBalance, "Currency=", currency)
	Response = append(Response, fmt.Sprintf("%v", CurrentBalance))

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(Response)
}
