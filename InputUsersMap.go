package main

import (
	"log"
)

func InputUsersMap(uL map[string]float64) map[string]float64 {
	var UserId string

	session := GetCassandraSession(KeyspaceName)
	defer session.Close()

	iter := session.Query(`SELECT user_id FROM wallet`).Iter()
	for iter.Scan(&UserId) {
		uL[UserId] = OutlayDayLimit()
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return uL
}
