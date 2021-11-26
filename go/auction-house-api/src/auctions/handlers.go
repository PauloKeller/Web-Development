package auctions

import (
	"encoding/json"
	"net/http"
)

var err error

func FetchAuctions(w http.ResponseWriter, r *http.Request) {
	a, err := AllAuctions()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(a)
}

func CreateAuction(w http.ResponseWriter, r *http.Request) {
	a := Auction{}

	err = json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	err = SaveAuction(a.MinimumBid, a.BuyOut, a.Owner, a.Item)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
