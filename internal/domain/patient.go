package domain

type Patient struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	DNI      int    `json:"dni" binding:"required"`
	Date     string `json:"date" binding:"required"`
}
