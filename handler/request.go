package handler

import "fmt"

func paramsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is Required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return paramsRequired("role", "String")
	}
	if r.Company == "" {
		return paramsRequired("company", "String")
	}
	if r.Location == "" {
		return paramsRequired("location", "String")
	}
	if r.Remote == nil {
		return paramsRequired("remote", "bool")
	}
	if r.Link == "" {
		return paramsRequired("link", "string")
	}
	if r.Salary <= 0 {
		return paramsRequired("salary", "Int64")
	}
	return nil
}

type UpdateRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Link != "" || r.Location != "" || r.Remote != nil || r.Salary <= 0 {
		return nil
	}
	return fmt.Errorf("Error errotr")
}
