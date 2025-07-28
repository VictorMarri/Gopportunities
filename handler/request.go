package handler

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func errRequestIsEmpty() error {
	return fmt.Errorf("request is empty")
}

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {

	fmt.Println("CreateOpeningRequest Validate started")

	if r.Role == "" &&
		r.Company == "" &&
		r.Location == "" &&
		r.Remote == nil &&
		r.Link == "" &&
		r.Salary == 0 {
		return errRequestIsEmpty()
	}

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
		return errParamIsRequired("salary", "int64")
	}

	fmt.Println("CreateOpeningRequest Validation returned Success")

	return nil
}

// Maybe use and interface here with CreateOpeningRequest?
type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	//If any field is provided, validation is true
	if r.Role != "" ||
		r.Company != "" ||
		r.Location != "" ||
		r.Remote != nil ||
		r.Link != "" ||
		r.Salary != 0 {
		return nil
	}

	//If none of the fields were provided, return false
	return fmt.Errorf("At least one valid field must be provided")
}
