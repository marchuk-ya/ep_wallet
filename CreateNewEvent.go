package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"wallet_ep/StructEvent"
)

// CreateNewEvent function add new event in budget
// method POST
func CreateNewEvent(_ http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: CreateNewEvent")

	var Levent StructEvent.Event

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &Levent)

	session := GetCassandraSession(KeyspaceName)
	defer session.Close()

	AmountToDouble, _ := strconv.ParseFloat(Levent.Amount, 64)

	switch {
	case (AmountToDouble < 0 && CheckLimit(Levent.UserId, AmountToDouble, usersLimits) == true) || AmountToDouble > 0:
		if err := session.Query(`INSERT INTO wallet (user_id, amount, currency, description, insertedtime)
							VALUES (?, ?, ?, ?, toTimeStamp(now()))`,
			Levent.UserId, AmountToDouble, Levent.Currency, Levent.Description).Exec(); err != nil {
			log.Fatal(err)
		}
	case AmountToDouble < 0 && CheckLimit(Levent.UserId, AmountToDouble, usersLimits) == false:
		fmt.Println("Day limit is over! Your day limit at least " + fmt.Sprintf("%v", usersLimits[Levent.UserId]))
	}
}
