package controllers

import (
	"encoding/json"
	"net/http"
	u "../utils"
	models "../models"
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