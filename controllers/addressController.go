package controllers

import (
	"../models"
	u "../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateAddresses = func (w http.ResponseWriter, r *http.Request){
	user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	address := &models.Address{}

	err := json.NewDecoder(r.Body).Decode(address)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	address.UserId = user
	resp := address.CreateAddress()
	u.Respond(w, resp)
}

var GetAddresses = func (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetAddressByUserID(uint(userId))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetAddress = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetAddressByID(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}