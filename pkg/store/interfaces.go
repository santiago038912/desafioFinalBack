package store

import (
	"github.com/desafioFinalBack/internal/domain"
)

type StoreInterfaceDentist interface {
	// Read devuelve un dentista por su id
	ReadDentist(id int) (domain.Dentist, error)
	// Create agrega un nuevo dentista
	CreateDentist(dentist domain.Dentist) error
	// Update actualiza un dentista
	UpdateDentist(dentist domain.Dentist) error
	// Delete elimina un dentista
	DeleteDentist(id int) error

}

type StoreInterfacePatient interface {
	// Read devuelve un paciente por su id
	ReadPatient(id int) (domain.Patient, error)
	// Create agrega un nuevo paciente
	CreatePatient(patient domain.Patient) error
	// Update actualiza un paciente
	UpdatePatient(patient domain.Patient) error
	// Delete elimina un paciente
	DeletePatient(id int) error

}

type StoreInterfaceTurn interface {
	// ReadTurn devuelve un paciente por su id
	ReadTurn(id int) (domain.Turn, error)
	// CreateTurn agrega un nuevo paciente
	CreateTurn(patient domain.Turn) error
	// UpdateTurn actualiza un paciente
	UpdateTurn(patient domain.Turn) error
	// DeleteTurn elimina un paciente
	DeleteTurn(id int) error
	// ReadTurnByDni devuelve un paciente por su id
	ReadTurnByDni(dni int) (domain.Turn, error)
}