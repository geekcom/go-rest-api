package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID        string   `json:"id"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

var people []Person

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getPeople")
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	fmt.Println("Endpoint Hit: getPerson")
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	fmt.Println("Endpoint Hit: createPerson")
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		fmt.Println("Endpoint Hit: deletePerson")
		json.NewEncoder(w).Encode(people)
	}
}

func handleRequests() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})

	router.HandleFunc("/", homePage)
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	handleRequests()
}
