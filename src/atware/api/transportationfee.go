package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	rt "gopkg.in/dancannon/gorethink.v1"
)

var (
	Session *rt.Session
)

type Record struct {
	From    location `json:"from" gorethink:"from"`
	To      location `json:"to" gorethink:"to"`
	Balance int      `json:"balance" gorethink:"balance"`
	SegNum  int      `json:"seg_num" gorethink:"seg_num"`
	Type    string   `json:"type" gorethink:"type"`
	Time    string   `json:"time" gorethink: "time"`
}

type location struct {
	RWC string `json:"rail_way_code" gorethink:"rwc"`
	SC  string `json:"station_code" gorethink:"sc"`
}

type Filter struct {
	PageSize   int `json:"page_size"`
	PageNumber int `json:"page_number"`
}

type Card struct {
	CardId     int    `json:"card_id" gorethink:"card_id"`
	CardHolder string `json:"card_holder" gorethink:"card_holder"`

	Records []Record `json:"records" gorethink:"records"`
}

func getCard(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	card_id := vars["card_id"]

	var newCard Card

	res, _ := rt.Table("card").Get(card_id).Run(Session)

	res.One(&newCard)
	data, _ := json.Marshal(newCard)
	rw.Write([]byte(data))
}

func updateCard(rw http.ResponseWriter, req *http.Request) {

}

func postCard(rw http.ResponseWriter, req *http.Request) {

	// parse body
	jsonData, _ := ioutil.ReadAll(req.Body)
	var newCard Card
	err := json.Unmarshal(jsonData, &newCard)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	cursor, err := rt.Table("card").Insert(newCard).Run(Session)
	if err != nil {
		log.Println(err.Error())
		return
	}

	var res interface{}
	err = cursor.One(&res)
	if err != nil {
		fmt.Print(err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
