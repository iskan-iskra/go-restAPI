package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	Location string `json:"location"`
	Skype    string `json:"skype"`
	Email    string `json:"email"`
	Skills   string `json:"skills"`
	JSlevel  string `json:"jslevel"`
	Comment  string `json:"comment"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(1000000))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}
func main() {
	r := mux.NewRouter()
	users = append(users, User{
		ID:       "1",
		Name:     "Iskander Shushayev",
		Age:      "25",
		Location: "Kazakhstan, Almaty",
		Skype:    "iskan_iskra",
		Email:    "bf2iskan@gmail.com",
		Skills:   "ALL skills",
		JSlevel:  "over 100%",
		Comment:  "TEST",
	})

	headers := handlers.AllowedHeaders([]string{"X-Requested-Width", "Content-Type", "Authorization", "application/json", "Accept"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(r)))
}
