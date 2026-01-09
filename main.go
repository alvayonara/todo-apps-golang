package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo-apps-golang/database"
	"todo-apps-golang/handlers"
)

func main() {
	database.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
