package main

import (
	"encoding/json"
	"fmt"
	"log"

	// "math/rand"
	"net/http"
	// "strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func main() {
	r := mux.NewRouter()
	director1 := Director{Firstname: "David", Lastname: "Fincher"}
	director2 := Director{Firstname: "Quentin", Lastname: "Tarantino"}
	director3 := Director{Firstname: "Christopher", Lastname: "Nolan"}
	movie1 := Movie{ID: "1", Isbn: "123", Title: "Se7en", Director: &director1}
	movie2 := Movie{ID: "2", Isbn: "842", Title: "Kill Bill", Director: &director2}
	movie3 := Movie{ID: "3", Isbn: "754", Title: "Tenet", Director: &director3}

	movies = append(movies, movie1)
	movies = append(movies, movie2)
	movies = append(movies, movie3)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	// r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at 8080")

	log.Fatal(http.ListenAndServe(":8080", r))

}
