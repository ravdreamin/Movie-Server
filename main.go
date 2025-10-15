package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movies struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movie []Movies

func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for index, item := range movie {

		if item.ID == param["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for _, item := range movie {

		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cmovies Movies

	err := json.NewDecoder(r.Body).Decode(&cmovies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmovies.ID = strconv.Itoa(rand.Intn(1000000))

	movie = append(movie, cmovies)
	json.NewEncoder(w).Encode(cmovies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)

	for index, item := range movie {
		if item.ID == param["id"] {

			movie = append(movie[:index], movie[index+1:]...)

			var updatedMovie Movies
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.ID = param["id"]
			movie = append(movie, updatedMovie)
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
}

func main() {

	r := mux.NewRouter()

	movie = append(movie, Movies{ID: "1", Isbn: "6745", Title: "Air", Director: &Director{
		Firstname: "Satyajit", Lastname: "Ray"}})

	movie = append(movie, Movies{ID: "2", Isbn: "6795", Title: "Hoe", Director: &Director{
		Firstname: "John", Lastname: "Doe"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server Started at port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))

}
