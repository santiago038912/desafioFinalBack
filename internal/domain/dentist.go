package domain

type Dentist struct {
	Id             int    `json:"id"`
	Name           string `json:"name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	RegisterNumber string `json:"register_number" binding:"required"`
}

type DentistDTO struct {
	Name           string `json:"name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	RegisterNumber string `json:"register_number,omitempty"`
}