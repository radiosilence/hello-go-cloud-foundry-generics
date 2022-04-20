package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Number interface {
	int64 | float64
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	http.HandleFunc("/", SumHandler)
	fmt.Printf("Server started at port %s\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

type Response struct {
	Ints   int64
	Floats float64
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("content-type", "application/json")

	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	sums := Response{
		Ints:   SumIntsOrFloats(ints),
		Floats: SumIntsOrFloats(floats),
	}

	if res, err := json.Marshal(&sums); err != nil {
		panic(err)
	} else {
		w.Write(res)
		w.Write([]byte("\n"))
	}
}
