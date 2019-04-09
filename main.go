package main

import(
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Person struct {
  ID string `json:"id,omitempty"`
  Firstname string `json:"firstname,omitempty"`
  Lastname string `json:"lastname,omitempty"`
  Address *Address `json:"address,omitempty"`
}
type Address struct {
  City string `json:"city,omitempty"`
  State string `json:"state,omitempty"`
}
func GetPeople(w http.ResponseWriter, r *http.Request) {
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
  // returns an empty person struct if nothing is found
  json.NewEncoder(w).Encode(&Person{})
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var person Person
  _ = json.NewDecoder(r.Body).Decode($person)
  person.ID = params["id"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for index, item := range people {
    if Item.ID == params["id"] {
      people = append(people[:index], people[index+1:]...)
      break
    }
    json.NewEncoder(w).Encode(people)
  }
}

var people []Person

func main() {
  router := mux.NewRouter()
  people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
  people = append(people, Person{ID: "3", Firstname: "Zhan", Lastname: "Liang"})
  // add handlers, the handler is usually nil
  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
  // listens on TCP network address and then calls Serve with handler to handle requests
  log.Fatal(http.ListenAndServe(":8080", router))
}
