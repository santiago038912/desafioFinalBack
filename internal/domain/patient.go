package domain

type Patient struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	DNI      int    `json:"dni" binding:"required"`
	Date     string `json:"date" binding:"required"`
}

type PatientDTO struct {
	Name     string `json:"name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Address  string `json:"address,omitempty"`
	DNI      int    `json:"dni,omitempty"`
	Date     string `json:"date,omitempty"`
}
