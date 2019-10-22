package main

import (
    "fmt"
	"github.com/gorilla/mux"
	"./controllers"
	"net/http"
	"os"
	"log"
)
func main() {
	router := mux.NewRouter()
	//router.HandleFunc("/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts
	router.HandleFunc("/users/{id}", controllers.GetContactsFor).Methods("GET")
	log.Println("Listening on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}