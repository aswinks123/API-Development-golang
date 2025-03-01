/*
=============================================================================
DEVELOPER: Aswin KS
DATE: 01-03-2025
ABOUT: Create login form API using Go
Functions added: User Creation, User Authentication, Session Management
===========================================================================
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// structure for holding user information
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// inmemory MAP for storing the user credentials
var (
	users = make(map[string]string) //inmemory user store using Map
	mu    sync.Mutex                // This is a lock to make sure only one process use this at a time. This is same as  var mu sync.Mutex

)

// handler function definition for /register api
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return

	}
	var user User                                //a new user variable of type User(struct). User struct was created on top
	err := json.NewDecoder(r.Body).Decode(&user) // The JSON body of the request is decoded into a User struct using json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Username == "" || user.Password == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	mu.Lock()         // function is called to lock the mutex, ensuring that only one goroutine can access the users map at a time.
	defer mu.Unlock() //ensures that the mutex is unlocked when the function exits, even if an error occurs.
	if _, exists := users[user.Username]; exists {
		http.Error(w, "Username  already exist", http.StatusConflict)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost) //The password is hashed using bcrypt.GenerateFromPassword
	if err != nil {
		http.Error(w, "Failed to has password", http.StatusInternalServerError)
		return
	}

	//Store the user to struct
	users[user.Username] = string(hashedPassword) //add password to username field of the user
	w.WriteHeader((http.StatusCreated))
	fmt.Fprint(w, "User registered Successfully")
}

// handler function definition for /login api
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Username == "" || user.Password == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	savedpassword, exists := users[user.Username] //check whether the password present in the struct. exists variable return true if its same
	mu.Unlock()

	if !exists || bcrypt.CompareHashAndPassword([]byte(savedpassword), []byte(user.Password)) != nil { //compare the struct password to provided password
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}
	// Set a session cookie (for simplicity, we're just using the username as the session value)
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    user.Username,
		Path:     "/",
		HttpOnly: true,                           // Prevent client-side scripts from accessing the cookie
		Secure:   true,                           // Only send the cookie over HTTPS
		Expires:  time.Now().Add(24 * time.Hour), // Set an expiration time
	})
	fmt.Fprintln(w, "Login Successful")

}

func main() {

	//handler for "/register" api
	http.HandleFunc("/register", registerHandler)

	//handler for /login api
	http.HandleFunc("/login", loginHandler)

	// Start the server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
