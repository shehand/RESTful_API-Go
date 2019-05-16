package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	u "../utils"
	models "../models"
	"strconv"
)

var CreateUser = func(w http.ResponseWriter, r * http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.CreateUser() //Create account
	u.Respond(w, resp)
}

var GetUser = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetUser(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}