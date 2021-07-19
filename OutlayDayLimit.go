package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Limit float64
}

// OutlayDayLimit function read day limit from configuration file
func OutlayDayLimit() float64 {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Limit= ", configuration.Limit)

	return configuration.Limit
}
