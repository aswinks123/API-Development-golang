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

// markComplete handler function
func markCompleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid: Only PUT requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	//Extract the ID from URLK
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	//convert ID from string to interger
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "Invalid request. Task ID must be a number", http.StatusBadRequest)
		return
	}

	//find and update the list element as completed
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			//Respond with the updated list
			w.Header().Set("Content-Type", "applicaion/json")
			json.NewEncoder(w).Encode(todos[i])
			fmt.Printf("✅ Task Marked as Completed: %+v\n", todos[i])
			return
		}
	}
}

// delete handler function
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	//check whether the method choosen is valid
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid. Only DELETE requests are alloweded", http.StatusMethodNotAllowed)
		return
	}
	//extract the ID from the URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	//convert the ID from string to interger
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "Invalid request: Task ID must be a number", http.StatusBadRequest)
		return
	}
	//find and delete the todo using the ID
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...) // Code to remove the task from the slice
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Task with ID %d deleted successfully", id)
			fmt.Printf("Delete To-Do: %+v\n", todo)
			return

		}
	}
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

	//handler to mark a To-Do as completed
	http.HandleFunc("/complete/", markCompleteHandler)

	//handler to delete a To-Do
	http.HandleFunc("/delete/", deleteHandler)

	// Start the server
	fmt.Println("✅ Server is running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("❌ Error starting server:", err)
	}

}
