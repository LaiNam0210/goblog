package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	DB struct {
		Port    int    `json:"port"`
		Address string `json:"address"`
		DBName  string `json:"dbName"`
	} `json:"db"`

	Server struct {
		Port    int    `json:"port"`
		Address string `json:"address"`
	} `json:"server"`

	API struct {
		Port    int    `json:"port"`
		Address string `json:"address"`
	} `json:"api"`

	FrontEnd struct {
		Port    int    `json: "port"`
		Address string `json: "address"`
	} `json:"frontEnd"`
}

func parseConfig(filename string) *Config {
	// config application
	file, err := os.Open(filename)

	if err != nil {
		panic(fmt.Sprintf("can not read filename: %v", filename))
	}

	if bytes, err := ioutil.ReadAll(file); err == nil {

		// unmarshal
		var config Config
		if err = json.Unmarshal(bytes, &config); err == nil {
			return &config
		}
		panic("can not unmarshal")

	}

	panic(fmt.Sprintf("can not read filename: %v", filename))
}
