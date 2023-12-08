package main

import (
	"encoding/json"
	"log"
	"net/http"
	// "io/ioutil"

)

type person struct {
	First string
}

func main() {

	p1 := person{
		First: "ram",
	}
	p2 := person{
		First: "Sita",
	}

	xs := []person{p1, p2}

	bs, err := json.Marshal(xs)

	if err != nil {
		log.Panic(err)
	}
	log.Println(string(bs))

	xs2 := []person{}

	err = json.Unmarshal(bs, &xs2)

	if err != nil {
		log.Panic(err)

	}

	log.Println("Back into a go data structure", xs2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", boo)
	http.HandleFunc("/encodeNew", encodeNewFun)
	http.HandleFunc("/decodeNew", decodeNewFun)

	http.ListenAndServe(":8080", nil)

}
func foo(w http.ResponseWriter, r *http.Request) {

	p1 := person{
		First: "Sarath",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)

	}

}
func boo(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded data", err)

	}
	log.Println("Person", p1)

}
func encodeNewFun(w http.ResponseWriter, r *http.Request) {
	p1 := []person{
		{First: "Lakshmi"},
		{First: "Lakshmi"},
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Error in encoding", err)
	}
}
func decodeNewFun(w http.ResponseWriter, r *http.Request) {
	var p []person

	err := json.NewDecoder(r.Body).Decode(&p)
	// body,_ := ioutil.ReadAll(r.Body)
	// err :=json.Unmarshal(body,&p)
	if err != nil {
		log.Println("Error while decoding", err)
	}
	log.Println("Person Data", p)

}
