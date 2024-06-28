package database

type Employee struct {
	EmployeeID string `json:"employeeId"`
	Name       string `json:"name"`
	City       string `json:"city"`
}

type EmployeeSearch struct {
	Fields []struct {
		FieldName string `json:"fieldName"`
		Eq        string `json:"eq,omitempty"`
		Neq       string `json:"neq,omitempty"`
	} `json:"fields"`
	Condition string `json:"condition"`
}
