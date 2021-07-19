package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"wallet_ep/StructEvent"
)

// GetListOfEventsByTime function return all budget events
//method GET
func GetListOfEventsByTime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetListOfEventsByTime")

	vars := mux.Vars(r)
	id := vars["id"]
	beginTime := vars["beginTime"]
	endTime := vars["endTime"]

	var Amount float64
	var UserId, Currency, Description string
	var Events []StructEvent.Event
	var Levent StructEvent.Event

	session := GetCassandraSession(KeyspaceName)
	defer session.Close()

	queryCQL := "SELECT user_id, amount, currency, description FROM wallet WHERE user_id ='" + id +
		"' AND insertedtime >='" + beginTime + "' AND insertedtime <='" + endTime + "' ALLOW FILTERING"

	iter := session.Query(queryCQL).Iter()

	for iter.Scan(&UserId, &Amount, &Currency, &Description) {
		fmt.Println("Event:", UserId, Amount, Currency, Description)
		Levent.UserId = UserId
		Levent.Amount = fmt.Sprintf("%v", Amount)
		Levent.Currency = Currency
		Levent.Description = Description
		Events = append(Events, Levent)
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(Events)
}
