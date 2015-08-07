package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	rt "gopkg.in/dancannon/gorethink.v1"

	"atWare/api"
)

var (
	configPath = flag.String("configPath", "./config.json", ``)
)

func connect(config *Config) *rt.Session {
	session, err := rt.Connect(rt.ConnectOpts{
		Address:  config.DB.Address,
		Database: config.DB.DBName,
	})

	if err != nil {
		panic(err.Error())
	}

	return session
}

func main() {
	flag.Parse()

	config := parseConfig(*configPath)

	dbSession := connect(config)

	go func() {
		fmt.Printf("Api is running at port %v \n", config.API.Port)

		api.Session = dbSession
		r := mux.NewRouter()
		api.SetupRouter(r, "/api/v1")

		http.ListenAndServe(fmt.Sprintf(":%v", config.API.Port), r)

	}()

	go func() {
		fmt.Printf("Frontend is running at port %v \n", config.FrontEnd.Port)
		frontEnd := http.NewServeMux()
		frontEnd.Handle("/", http.FileServer(http.Dir("./public")))
		http.ListenAndServe(fmt.Sprintf(":%v", config.FrontEnd.Port), frontEnd)

	}()
	go func() {
		fmt.Printf("Server is running at port %v \n", config.Server.Port)
		r := mux.NewRouter()

		// server := Server{}
		// server.Handle("/admin", r)
		http.ListenAndServe(fmt.Sprintf(":%v", config.Server.Port), r)
	}()

	ch := make(chan int)
	<-ch
}
