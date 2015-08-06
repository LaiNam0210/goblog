package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/dancannon/gorethink.v1"
)

var (
	Session *gorethink.Session
)

type Record struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Balance int    `json:"balance"`
	SegNum  int    `json:"seg_num"`
	Type    string `json:"type"`
	JSTTime string `json:"jst_time"`
}

type Filter struct {
	PageSize   int `json:"page_size"`
	PageNumber int `json:"page_number"`
}

type Card struct {
	Id int `json:"id"`

	CardId     int    `json:"card_id"`
	CardHolder string `json:"card_holder"`

	Records []Record `json:"records"`
}

func getCard(rw http.ResponseWriter, req *http.Request) {
	// vars := mux.Vars(req)
	id := mux.Vars(req)["card_id"]
	fmt.Println(req.URL.String())
	fmt.Println(id)
	rw.Write([]byte("get cards"))
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

	fmt.Println(string(jsonData))
	rw.WriteHeader(http.StatusCreated)
}
