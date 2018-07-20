package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	Street string `json:"street,omitempty"`
	City   string `json:"city,omitempty"`
	State  string `json:"state,omitempty"`
}

var persons []Person

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range persons {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(persons)
}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	persons = append(persons, person)
	json.NewEncoder(w).Encode(person)
}
func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:index], persons[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(persons)
}

func main() {
	router := mux.NewRouter()
	persons = append(persons, Person{ID: "1", Firstname: "Natraj", Lastname: "Bontha", Address: &Address{City: "Mckinney", State: "Tx", Street: "Planet Geo"}})
	persons = append(persons, Person{ID: "2", Firstname: "Migmar", Lastname: "Lhamo", Address: &Address{City: "Mckinney", State: "Tx", Street: "Planet Geo"}})
	persons = append(persons, Person{ID: "3", Firstname: "Dolma", Lastname: "Bontha", Address: &Address{City: "Mckinney", State: "Tx", Street: "Planet Geo"}})
	persons = append(persons, Person{ID: "4", Firstname: "Nikhil", Lastname: "Bontha", Address: &Address{City: "Mckinney", State: "Tx", Street: "Planet Geo"}})
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":1234", router))
}
