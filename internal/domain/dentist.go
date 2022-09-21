package domain

type Dentist struct {
	Id             int    `json:"id"`
	Name           string `json:"name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	RegisterNumber string `json:"register_number" binding:"required"`
}