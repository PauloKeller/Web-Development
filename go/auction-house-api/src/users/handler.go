package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var err error

//CreateUser handle to call a create item
func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := User{}
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	err = SaveUser(u.FullName, u.NickName, u.Email, u.Balance)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// FetchUsers tries to call AllItems to retrieve all itens from the
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	items, err := AllUsers()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

// FetchUserByID tries to call UserByID to retrieve the unique user
func FetchUserByID(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id := v["user_id"]
	u, err := UserByID(id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(u)
}
