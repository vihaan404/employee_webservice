package database

import (
	"encoding/json"
	"log"
	"os"
)

type Database struct {
	Conn *os.File
}

func CreateDatabaseConnection() (*Database, error) {
	file, err := os.OpenFile("employee.json", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	dat, err := os.ReadFile(file.Name())
	if err != nil {
		return nil, err
	}
	if string(dat) == "" {
		os.WriteFile(file.Name(), []byte("[]"), 0644)
	}

	return &Database{
		file,
	}, nil
}

func (d Database) CloseConnection() {
	err := d.Conn.Close()
	if err != nil {
		log.Println("error closing the connection")
	}
}

func (d Database) CreateEmployee(e Employee) (string, error) {
	dat, err := os.ReadFile(d.Conn.Name())
	if err != nil {
		return "", err
	}
	var employees []Employee

	err = json.Unmarshal(dat, &employees)
	if err != nil {
		return "", err
	}

	employees = append(employees, e)
	inputDat, err := json.MarshalIndent(employees, "", " ")
	if err != nil {
		return "", err
	}
	os.WriteFile(d.Conn.Name(), inputDat, 0644)

	return e.EmployeeID, nil
}

func (d Database) GetAllEmployee() ([]Employee, error) {
	dat, err := os.ReadFile(d.Conn.Name())
	if err != nil {
		return nil, err
	}

	var employees []Employee

	err = json.Unmarshal(dat, &employees)
	if err != nil {
		return nil, err
	}

	return employees, nil
}
