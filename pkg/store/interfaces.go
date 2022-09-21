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
	// Read devuelve un dentista por su id
	ReadPatient(id int) (domain.Patient, error)
	// Create agrega un nuevo dentista
	CreatePatient(dentist domain.Patient) error
	// Update actualiza un dentista
	UpdatePatient(dentist domain.Patient) error
	// Delete elimina un dentista
	DeletePatient(id int) error

}