package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Person struct (Model)
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address struct (Model)
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// Init people var as a Person struct
var people []Person

//Get all the persons
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// Get a specific person
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) // Gets params

	// Loop through persons and find one with the id from the params
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

//Add new person
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

//Update Person
func UpdatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			var person Person
			_ = json.NewDecoder(req.Body).Decode(&person)
			person.ID = params["id"]
			if params["firstname"] == "" {
				person.FirstName = people[index].FirstName
			} else {
				person.FirstName = params["firstname"]
			}
			if params["lastname"] == "" {
				person.LastName = people[index].LastName

			}else {
				person.LastName = params["lastname"]

			}
			people = append(people, person)
			json.NewEncoder(w).Encode(people)
			return
		}
	}
}

// Delete a person
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

//Main fuction
func main() {
	//Init router
	router := mux.NewRouter()

	//Create example data to the people array
	people = append(people, Person{ID: "1", FirstName: "Felipe", LastName: "Jurado",
		Address: &Address{City: "Palmira", State: "Valle del cauca"}})

	people = append(people, Person{ID: "2", FirstName: "Luz Isleny", LastName: "Murillo",
		Address: &Address{City: "Cali", State: "Valle del cauca"}})

	// Route handles & endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")

	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")

	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")

	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	router.HandleFunc("/people/{id}", UpdatePersonEndpoint).Methods("PUT")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", router))

}
