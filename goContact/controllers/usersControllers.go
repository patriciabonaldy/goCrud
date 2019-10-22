package controllers

import (
	"encoding/json"
	"../models"
	u "../utils"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	contact := &models.User{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.Id = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	v, err:= strconv.Atoi(mux.Vars(r)["id"])	
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	id:=uint(v)
	data := models.GetUser(id)
	json.NewEncoder(w).Encode(data)
}

