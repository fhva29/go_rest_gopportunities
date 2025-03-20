package handler

import (
	"fmt"
	"reflect"
)

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

// CreateOpening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int    `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int")
	}

	return nil
}

// UpdateOpening

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int    `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// Validate field types
	if reflect.TypeOf(r.Role).Kind() != reflect.String {
		return fmt.Errorf("role must be a string")
	}
	if reflect.TypeOf(r.Company).Kind() != reflect.String {
		return fmt.Errorf("company must be a string")
	}
	if reflect.TypeOf(r.Location).Kind() != reflect.String {
		return fmt.Errorf("location must be a string")
	}
	if reflect.TypeOf(r.Link).Kind() != reflect.String {
		return fmt.Errorf("link must be a string")
	}
	if reflect.TypeOf(r.Remote).Kind() != reflect.Bool {
		return fmt.Errorf("remote must be a bool")
	}
	if reflect.TypeOf(r.Salary).Kind() != reflect.Int {
		return fmt.Errorf("salary must be an int")
	}

	// If any field is provided, validation is true
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one field must be provided")
}
