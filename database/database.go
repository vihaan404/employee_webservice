package database

import (
	"bufio"
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
	writer := bufio.NewWriter(d.Conn)
	dat, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	_, _ = writer.Write(dat)
	err = writer.Flush()
	if err != nil {
		return "", err
	}
	return e.EmployeeID, nil
}
