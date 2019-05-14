package main

import (
	"./app"
	"./controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main(){
	router:= mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")

	port:= os.Getenv("PORT")
	if(port == ""){
		port = "8000"
	}

	fmt.Println(port)

	err:=http.ListenAndServe(":" + port, router)
	if err != nil{
		fmt.Print(err)
	}
}
