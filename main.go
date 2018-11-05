package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var people []Person
var data Data

func main() {
	router := mux.NewRouter()

	//Initialize people with some persons
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	//Declare routes and handler functions
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	//Listen on port 80000
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	data.Data = people
	data.Status = true
	data.Message = "People fetched successfully"
	json.NewEncoder(w).Encode(data)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item:= range people{
		if item.ID == params["id"] {
			data.Status = true
			data.Message = "Person with ID of " + params["id"] + " fetched successfully"

			var p []Person
			p = append(p, item)
			data.Data = p
			json.NewEncoder(w).Encode(data)
			return
		}
	}

	//Failed to fetch person with the ID
	data.Status = false
	data.Message = "Error fetching person with ID == " + params["id"]
	json.NewEncoder(w).Encode(data)
	//json.NewEncoder(w).Encode(&Person{})
	//json.NewEncoder(w).Encode(Error{Message:"Error fetching person with ID == " + params["id"]})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)

	data.Data = people
	data.Status = true
	data.Message = "Person created successfully"

	json.NewEncoder(w).Encode(data)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...) //May need a fix
			break
		}
	}

	data.Data = people
	data.Status = true
	data.Message = "Person deleted successfully"
	json.NewEncoder(w).Encode(data)
}