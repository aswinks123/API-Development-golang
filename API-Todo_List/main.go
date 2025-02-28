/*
=============================================================================
DEVELOPER: Aswin KS
DATE: 26-02-2025
ABOUT: Create a TO-Do list using GO, HTML and Javascript
Functions added: Display homepage, Create new todos, List the existing todos
===========================================================================
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// create a structure that holds the To-do content
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

//create a list that stores the To-Dos

var todos []Todo
var idCounter = 1 //ID of todo stars from 1 instead of 0

// createHandler function
func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid: Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil || newTodo.Task == "" {
		http.Error(w, "Invalid request body or Missing Task", http.StatusBadRequest)
		return
	}

	//Assign an ID and store the task
	newTodo.ID = idCounter
	idCounter++
	todos = append(todos, newTodo)

	//Respond with the created task
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
	fmt.Printf("New To-Do Created: %+v\n", newTodo)
}

// homeHandler function
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the To-Do API")

}

// listHandler function
func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid: Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	//Display the content as response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// searchHandler function
func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid: Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	//extract the ID from the URL path
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	//convert the string ID to integer
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "Invalid request: Task ID must be a number", http.StatusBadRequest)
		return
	}
	//search the task
	for _, todo := range todos {
		if todo.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	// If no task is found, return 404
	http.Error(w, "Task not found", http.StatusNotFound)

}

func main() {
	// Serve static files (HTML, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handler to serve index.html on root "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	//handler to create a new To-Do
	http.HandleFunc("/create", createHandler)

	//handler to list the To-Dos
	http.HandleFunc("/list", listHandler)

	//handler to search the To-Dos
	http.HandleFunc("/search/", searchHandler)

	// Start the server
	fmt.Println("✅ Server is running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("❌ Error starting server:", err)
	}

}
