package main

import (
	"./app"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main(){
	router:= mux.NewRouter()
	router.Use(app.JwtAuthentication)

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
