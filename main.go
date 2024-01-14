package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	for index, film := range movies {
		if film.Id == vars["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
		json.NewEncoder(w).Encode(movies)

	}

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	for _, movie := range movies {
		if movie.Id == vars["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	var film Movie
	w.Header().Set("content-type", "application/json")
	json.NewDecoder(r.Body).Decode(&film)
	film.Id = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, film)
	json.NewEncoder(w).Encode(movies)

}
func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	for index, film := range movies {
		if film.Id == vars["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var film Movie
			json.NewDecoder(r.Body).Decode(&film)
			movies = append(movies, film)
			json.NewEncoder(w).Encode(movies)

		}

	}

}

func main() {
	movies = append(movies, Movie{Id: "1", Isbn: "423ad15", Title: "Last battle", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{Id: "2", Isbn: "789sw34", Title: "Winner Takes All", Director: &Director{Firstname: "Amaechi", Lastname: "Ebite"}})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
