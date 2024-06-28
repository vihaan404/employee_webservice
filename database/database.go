package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
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

func (d Database) GetEmployee(employeeId string) (*Employee, error) {
	dat, err := os.ReadFile(d.Conn.Name())
	if err != nil {
		return nil, err
	}

	var employees []Employee

	err = json.Unmarshal(dat, &employees)
	if err != nil {
		return nil, err
	}
	for _, e := range employees {
		if e.EmployeeID == employeeId {
			return &e, nil
		}
	}

	return nil, ErrEmployeeNotFound
}

var ErrEmployeeNotFound = errors.New("employee not found yep")

func (d Database) UpdateEmployee(e Employee, employeeId string) (*Employee, error) {
	dat, err := os.ReadFile(d.Conn.Name())
	if err != nil {
		return nil, err
	}

	var employees []Employee

	err = json.Unmarshal(dat, &employees)
	if err != nil {
		return nil, err
	}
	found := false
	for i := range employees {
		if employees[i].EmployeeID == employeeId {
			employees[i] = e
			found = true

		}
	}

	if !found {
		return nil, ErrEmployeeNotFound
	}

	return &e, nil
}

func (d Database) DeleteEmployee(employeeId string) (*Employee, error) {
	dat, err := os.ReadFile(d.Conn.Name())
	if err != nil {
		return nil, err
	}

	var employees []Employee

	err = json.Unmarshal(dat, &employees)
	if err != nil {
		return nil, err
	}
	found := false
	index := 0

	e := Employee{}
	for i := range employees {
		if employees[i].EmployeeID == employeeId {
			found = true
			index = i
			e = employees[i]

		}
	}
	if !found {
		return nil, ErrEmployeeNotFound
	}

	employees = append(employees[:index], employees[index+1:]...)
	updatedData, err := json.MarshalIndent(employees, "", " ")
	if err != nil {
		return nil, err
	}

	os.WriteFile(d.Conn.Name(), updatedData, 0644)

	return &e, nil
}

func (d Database) SearchEmployees(searchParams EmployeeSearch) ([]Employee, error) {
	dat, err := os.ReadFile(d.Conn.Name())
	if err != nil {
		return nil, err
	}

	var employees []Employee

	err = json.Unmarshal(dat, &employees)
	if err != nil {
		return nil, err
	}
	searchResult := []Employee{}

	for _, field := range searchParams.Fields {
		fieldName := field.FieldName

		ans := []Employee{}
		for _, e := range employees {
			v := reflect.ValueOf(e)

			fields := v.FieldByName(strings.Title(fieldName))
			if !fields.IsValid() {
				return nil, fmt.Errorf("no such field: %s in employee", fieldName)
			}
			if field.Eq != "" {
				if fields.String() == field.Eq {
					ans = append(ans, e)
				}
			}
			if field.Neq != "" {
				if fields.String() != field.Neq {
					ans = append(ans, e)
				}
			}

		}

		if searchParams.Condition == "OR" {
			searchResult = append(searchResult, ans...)
		}
		if searchParams.Condition == "AND" {
			searchResult = ans
			employees = searchResult

		}

	}

	return searchResult, nil
}
