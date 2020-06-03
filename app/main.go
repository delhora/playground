package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", helloWorld)
	http.ListenAndServe(":8080",router)
}