package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to api de pepe")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	log.Fatal(http.ListenAndServe(":2000", router))
}
