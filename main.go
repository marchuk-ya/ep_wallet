package main

import (
	"fmt"
)

const KeyspaceName = "wallet_ep"

var usersLimits = make(map[string]float64)

func main() {
	fmt.Println("Rest API - RUN.")
	CassandraCheckTableExists()
	fmt.Println(InputUsersMap(usersLimits))
	handleRequests()
}
