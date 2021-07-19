package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"os"
	"time"
)

const MaxConnectionAttempt = 60

func GetCassandraSession(spaceName string) *gocql.Session {
	ipAddress := os.Getenv("CASSANDRA_HOST")
	if ipAddress == "" {
		panic("IP address nil")
	}
	cluster := gocql.NewCluster(ipAddress)
	if spaceName == KeyspaceName {
		cluster.Keyspace = spaceName
	}
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4

	var err error
	var session *gocql.Session
	for connectionAttempt := 0; connectionAttempt < MaxConnectionAttempt; connectionAttempt++ {
		session, err = cluster.CreateSession()
		if err == nil {
			fmt.Println("Connected to Cassandra!")
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic(err)
	}
	return session
}
