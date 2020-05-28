package configuration

import (
	/*
		"encoding/json"
		"fmt"
		"os"

		"github.com/Cloud-Native-programming-with-Golang/chapter02/myevents/src/lib/persistence/dblayer"
	*/
	"encoding/json"
	"fmt"
	"os"

	"github.com/Cloud-Native-programming-with-Golang/chapter02/myevents/src/lib/persistence/dblayer"
)

var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault    = "localhost:8181"
)

type ServiceConfig struct {
	Databasetype     dblayer.DBTYPE `json:"databasetype"`
	DBConnection     string         `json:"dbconnection"`
	RestfulEPDefault string         `json:"restfulapi_endpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		fmt.Printf("*** Access endpoint: %v\n", conf.RestfulEPDefault)

		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	fmt.Printf("*** Access endpoint: %v\n", conf.RestfulEPDefault)

	return conf, err
}
