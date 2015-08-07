package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	rt "gopkg.in/dancannon/gorethink.v1"
)

var (
	newCard = Card{
		CardId:     11,
		CardHolder: "UserA",
		Records: []Record{
			{
				From: location{
					RWC: "01",
					SC:  "11",
				},
				To: location{
					RWC: "4d",
					SC:  "ec",
				},
				Balance: 2688,
				SegNum:  12,
				Type:    "gate",
				Time:    "2015-08-06",
			},
			{
				From: location{
					RWC: "01",
					SC:  "11",
				},
				To: location{
					RWC: "4d",
					SC:  "ec",
				},
				Balance: 2288,
				SegNum:  13,
				Type:    "gate",
				Time:    "2015-08-06",
			},
		},
	}
	r *mux.Router
)

type response struct {
	Errors        int      `gorethink:"errors"`
	GeneratedKeys []string `gorethink:"generated_keys"`
}

func init() {
	Session, _ = rt.Connect(rt.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
	})

	// rt.DB("test").TableDrop("card").Run(Session)
	// rt.DB("test").TableCreate("card").Run(Session)

	r = mux.NewRouter()
	SetupRouter(r, "/api/v1")

}
func TestGetCard(t *testing.T) {
	//create new records
	res, err := rt.Table("card").Insert(newCard).Run(Session)
	if err != nil {
		panic(err.Error())
	}
	var result response
	res.One(&result)

	req, _ := http.NewRequest("GET",
		fmt.Sprintf("/api/v1/cards/%v", result.GeneratedKeys[0]), nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("expected %v but got %v", http.StatusOK, w.Code)
		return
	}

	var cardRes Card
	err = json.Unmarshal(w.Body.Bytes(), &cardRes)
	if err != nil {
		t.Error(err.Error())
	}

	if cardRes.CardId != newCard.CardId {
		t.Errorf("Expected CardId %v but got %v", newCard.CardId, cardRes.CardId)
		return
	}
}

func TestPostCard(t *testing.T) {

	newCardBytes, _ := json.Marshal(newCard)
	req, _ := http.NewRequest("POST", "api/v1/cards", bytes.NewReader(newCardBytes))

	w := httptest.NewRecorder()

	postCard(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected %v but got %v", http.StatusCreated, w.Code)
		return
	}
}
