package utils

import (
	"encoding/json"
	"net/http"
)

func Messege(status bool, messege string, )(map[string]interface{}){
	return map[string]interface{}{"status" : status, "messege": messege}
}

func Respond(w http.ResponseWriter, data map[string]interface{}){
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(data)
}
