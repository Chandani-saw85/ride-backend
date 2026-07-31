package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Driver struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Vehicle string `json:"vehicle"`
	Online  bool   `json:"online"`
}

var drivers = []Driver{
	{ID: 1, Name: "Ramesh", Vehicle: "Swift Dzire", Online: true},
	{ID: 2, Name: "Suresh", Vehicle: "Honda City", Online: false},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server Running!")
}

func listDrivers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drivers)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"router": "CHI IS WORKING",
	})
}

func main() {
	connectDB()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/", home)
	r.Get("/health", health)
	r.Get("/api/drivers", listDrivers)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
