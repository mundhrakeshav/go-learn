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


type Movie struct {
	ID string `json:"id"`;
	Isbn string `json:"isbn"`;
	Title string `json:"title"`;
	Director *Director `json:"director"`;
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}
var movies []Movie;

func getMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json");
	json.NewEncoder(w).Encode(movies);
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json");
params := mux.Vars(r);
for index, movie := range movies {
	if movie.ID == params["id"] {
		movies = append(movies[:index], movies[index+1:]...); 
		var movie Movie;
		_ = json.NewDecoder(r.Body).Decode(&movie);
		movie.ID = strconv.Itoa(rand.Intn(10000));
		movies = append(movies, movie)
		json.NewEncoder(w).Encode(movies);
		break;
	}
}
json.NewEncoder(w).Encode(movies);

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...); 
			break;
		}
	}
	json.NewEncoder(w).Encode(movies);
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	var movie Movie;
	_ = json.NewDecoder(r.Body).Decode(&movie);
	movie.ID = strconv.Itoa(rand.Intn(10000));
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies);
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	params := mux.Vars(r);
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return;
		}
	}
}

func main()  {
	r := mux.NewRouter();
	movies = append(movies, Movie {
		ID: "0",
		Isbn: "0",
		Title: "Movie 1",
		Director: &Director{
			FirstName: "John",
			LastName: "Doe",
		},
	}, Movie {
		ID: "1",
		Isbn: "1",
		Title: "Movie 2",
		Director: &Director{
			FirstName: "Keshav",
			LastName: "Mundhra",
		},
	},
);
	r.HandleFunc("/movies", getMovies).Methods("GET");
	r.HandleFunc("/movies/:id", getMovie).Methods("GET");
	r.HandleFunc("/movies", createMovie).Methods("POST");
	r.HandleFunc("/movies", updateMovie).Methods("PUT");
	r.HandleFunc("/movies", deleteMovie).Methods("DELETE");
	fmt.Println("Starting port at 8080");
	log.Fatal(http.ListenAndServe(":8080", r));
}