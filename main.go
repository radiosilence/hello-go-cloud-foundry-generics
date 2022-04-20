package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Number interface {
	int64 | float64
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {

	http.HandleFunc("/", HelloHandler)

	fmt.Printf("Server started at port %s\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("content-type", "application/json")

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Fprintf(w, "Generic Sums: %v and %v\n", SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}
