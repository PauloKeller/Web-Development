package items

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var err error

//DeleteItem handle to call a delete Item by ID
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["item_id"]

	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	err := EraseItem(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//CreateItem handle to call a create item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	i := Item{}
	err := json.NewDecoder(r.Body).Decode(&i)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	err = SaveItem(i.Name, i.Description)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// FetchItems tries to call AllItems to retrieve all itens from the
func FetchItems(w http.ResponseWriter, r *http.Request) {
	items, err := AllItems()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}
