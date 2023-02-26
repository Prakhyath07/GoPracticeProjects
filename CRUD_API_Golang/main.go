package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `jsono:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firsname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r:= mux.NewRouter()

	movies = append(movies, Movie{ID: "1",Isbn: "233433", Title: "Movie 1", Director: &Director{FirstName: "Prak", LastName: "Bhandary"}})
	movies = append(movies, Movie{ID: "2",Isbn: "123456", Title: "Movie 2", Director: &Director{FirstName: "Sun", LastName: "Bhandary"}})

	r.HandleFunc("/movies",getMovies).Methods(("GET"))
	r.HandleFunc("/movies/{id}", getMovie).Methods(("GET"))
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}