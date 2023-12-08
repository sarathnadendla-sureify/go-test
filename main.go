package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", boo)

	http.ListenAndServe(":8080", nil)

}
func foo(w http.ResponseWriter, r *http.Request) {

	p1 := person{
		First: "Jenny",
	}
	p2 := person{
		First: "James",
	}
	people := []person{p1, p2}

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println("Encoding issue")
	}

}
func boo(w http.ResponseWriter, r *http.Request) {
	p := []person{}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println("Decoding issue", err)
	}
	log.Println("People Data", p)

}
