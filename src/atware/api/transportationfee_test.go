package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCard(t *testing.T) {

	// TODO: why mux cannot get param
	req, _ := http.NewRequest("GET", "api/v1/cards/123", nil)

	w := httptest.NewRecorder()

	getCard(w, req)

	if w.Code != 200 {
		t.Errorf("expected %v but got %v", http.StatusOK, w.Code)
		return
	}

	log.Println(string(w.Body.Bytes()))

}

func TestPostCard(t *testing.T) {
	newCard := Card{
		Id:         1,
		CardId:     11,
		CardHolder: "UserA",
		Records: []Record{
			{
				From:    "横浜",
				To:      "日の出町",
				Balance: 2688,
				SegNum:  12,
				Type:    "gate",
				JSTTime: "2015-08-06",
			},
			{
				From:    "横浜",
				To:      "日の出町",
				Balance: 2288,
				SegNum:  13,
				Type:    "gate",
				JSTTime: "2015-08-06",
			},
		},
	}

	newCardBytes, _ := json.Marshal(newCard)
	req, _ := http.NewRequest("POST", "api/v1/cards", bytes.NewReader(newCardBytes))

	w := httptest.NewRecorder()

	postCard(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected %v but got %v", http.StatusCreated, w.Code)
		return
	}
}
