package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	NamaLengkap string
	KodePeserta string
}

var irfan = Person{
	NamaLengkap: "Muhammad Irfan",
	KodePeserta: "149368582101-639",
}

// - middleware 1 : halo middleware
// - middleware 2 : nampilin struct saja (pada soal no. 1)
// - endpoint : nampilin semua resp dari sesi 1.
func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(tampilNama))

	endpoint := http.HandlerFunc(greet)
	mux.Handle("/hello/middleware", middlewares(middlewares2(endpoint)))
	http.ListenAndServe(":3000", mux)
}

func middlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("halo middleware")
		next.ServeHTTP(w, r)
	})
}
func middlewares2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nNama Lengkap : %s\nKode Peserta : %s\n", irfan.NamaLengkap, irfan.KodePeserta)
		next.ServeHTTP(w, r)
	})
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")

	fmt.Println("Scalable", "web", "Service", "Golang")

	w.Write([]byte("Hello"))
}

func tampilNama(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(irfan)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
