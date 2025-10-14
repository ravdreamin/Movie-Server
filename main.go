package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"math/rand"
	"github.com/gorilla/mux"
)

type Movies struct {
	id       string    `json:"id"`
	isbn     string    `json:"isbn"`
	title    string    `json:"title"`
	director *Director `json:"director"`
}

type Director struct {
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
}

var movie []Movies

func main() {

	r := mux.NewRouter()

	movie = append(movie, Movies{id: "1", isbn: "6745", titile: "Air", director: &Director{
		firstname: "Satyajit" , lastname: "Ray"}
})

movie = append(movie, Movies{id: "2", isbn: "6795", titile: "Hoe", director: &Director{
		firstname: "John" , lastname: "Doe"}
})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server Started at port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))

}
