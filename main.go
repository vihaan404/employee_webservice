package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/vihaan404/employee_webservice/database"
)

type api struct {
	db *database.Database
}

func main() {
	conn, err := database.CreateDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	api := &api{
		db: conn,
	}
	defer conn.CloseConnection()
	port := "8080"
	router := chi.NewRouter()
	router.Get("/greeting", api.handlerGreeting)
	router.Get("/employee/{id}", api.handlerGetEmployee)
	router.Post("/employee", api.handlerCreateEmployee)
	router.Get("/employees/all ", handlerGetAllEmployee)
	router.Put("/employee/{id}", handlerUpdateEmployee)
	router.Delete("/employee/{id}", handlerDeleteEmployee)
	router.Post("/employees/search", handlerEmployeeSearch)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	fmt.Println("Listening on port" + port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func createDataBase() {
	file, err := os.OpenFile("employee.json", os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error opening employee.json")
	}
	defer file.Close()
	dat := database.Employee{
		EmployeeID: "SomeRandomNumber",
		Name:       "somename",
		City:       "somecity",
	}
	result, err := json.Marshal(dat)
	if err != nil {
		log.Println("error Marshaling the json")
	}

	writer := bufio.NewWriter(file)
	numberOfBytes, err := writer.Write(result)
	if err != nil {
		log.Println("Could not write in the buffer ")
	}
	log.Println(numberOfBytes)

	err = writer.Flush()
	if err != nil {
		log.Println("flush faled")
	}
}

func (a api) handlerGreeting(w http.ResponseWriter, r *http.Request) {
	type status struct {
		Status string `json:"status"`
	}
	respondWithJson(w, http.StatusOK, status{
		Status: "ok",
	})
}

func (a api) handlerGetEmployee(w http.ResponseWriter, r *http.Request) {
}

func (a api) handlerCreateEmployee(w http.ResponseWriter, r *http.Request) {
	type CreateEmployeeBody struct {
		Name string `json:"name"`
		City string `json:"city"`
	}
	createEmployeeBody := CreateEmployeeBody{}

	err := json.NewDecoder(r.Body).Decode(&createEmployeeBody)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	params := database.Employee{
		EmployeeID: uuid.NewString(),
		Name:       createEmployeeBody.Name,
		City:       createEmployeeBody.City,
	}
	id, err := a.db.CreateEmployee(params)
	if err != nil {
		log.Println(err)
	}
	type response struct {
		EmployeeID string `json:"employeeId"`
	}
	respondWithJson(w, http.StatusCreated, response{
		EmployeeID: id,
	})
}

func handlerGetAllEmployee(w http.ResponseWriter, r *http.Request) {
}

func handlerUpdateEmployee(w http.ResponseWriter, r *http.Request) {
}

func handlerDeleteEmployee(w http.ResponseWriter, r *http.Request) {
}

func handlerEmployeeSearch(w http.ResponseWriter, r *http.Request) {
}
