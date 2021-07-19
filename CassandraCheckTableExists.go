package main

import (
	"log"
)

func CassandraCheckTableExists() {
	sessionKeyspace := GetCassandraSession("")
	defer sessionKeyspace.Close()

	// create keyspace
	errKeyspace := sessionKeyspace.Query("CREATE KEYSPACE IF NOT EXISTS wallet_ep WITH REPLICATION = " +
		"{'class' : 'SimpleStrategy', 'replication_factor':1};").Exec()
	if errKeyspace != nil {
		log.Println(errKeyspace)
		return
	}

	session := GetCassandraSession(KeyspaceName)
	defer session.Close()

	// create table
	err := session.Query("CREATE TABLE IF NOT EXISTS wallet_ep.wallet (user_id text, amount double, " +
		"currency text, description text, insertedtime timestamp PRIMARY KEY);").Exec()
	if err != nil {
		log.Println(err)
		return
	}
}
