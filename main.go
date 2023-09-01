package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "3455", Title: "Game of thrones", Director: &Director{Firstname: "Scar", Lastname: "Face"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4647", Title: "Legend of the seeker", Director: &Director{Firstname: "Jays", Lastname: "Vintage"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/(id)", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/(id)", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/(id)", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(":8000", r)
}
