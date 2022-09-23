package domain

type Turn struct {
	Id          int     `json:"id"`
	Dentist     Dentist `json:"dentist" binding:"required"`
	Patient     Patient `json:"patient" binding:"required"`
	Date        string  `json:"date" binding:"required"`
	Time        string  `json:"time" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type TurnDTO struct {
	Date                  string `json:"date,omitempty"`
	Time                  string `json:"time,omitempty"`
	Description           string `json:"description,omitempty"`
}