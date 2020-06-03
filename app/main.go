package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Greeting struct {
	Hello string

}
func helloWorld ( w http.ResponseWriter, r *http.Request)  {
	greeting := Greeting{Hello: "world"}
	js, err := json.Marshal(greeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {

	http.HandleFunc("/hello", helloWorld)
	log.Println("Server up and running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}