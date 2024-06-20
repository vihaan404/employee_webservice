package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := "8080"
	router := chi.NewRouter()
	router.Get("/greeting ", handlerGreeting)
	router.Get("/employee/{id}", handlerGetEmployee)
	router.Post("/employee", handlerCreateEmployee)
	router.Get("/employees/all ", handlerGetAllEmployee)
	router.Put("/employee/{id}", handlerUpdateEmployee)
	router.Delete("/employee/{id}", handlerDeleteEmployee)
	router.Post("/employees/search", handlerEmployeeSearch)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	fmt.Println("Listening on port" + port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handlerGreeting(w http.ResponseWriter, r *http.Request) {
}

func handlerGetEmployee(w http.ResponseWriter, r *http.Request) {
}

func handlerCreateEmployee(w http.ResponseWriter, r *http.Request) {
}

func handlerGetAllEmployee(w http.ResponseWriter, r *http.Request) {
}

func handlerUpdateEmployee(w http.ResponseWriter, r *http.Request) {
}

func handlerDeleteEmployee(w http.ResponseWriter, r *http.Request) {
}

func handlerEmployeeSearch(w http.ResponseWriter, r *http.Request) {
}
