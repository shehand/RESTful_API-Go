package controllers
import (
	"../models"
	u "../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateAddresses(){}

var GetAddresses = func (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetAddresses(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}