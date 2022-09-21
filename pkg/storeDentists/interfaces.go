package store

import (
	"github.com/desafioFinalBack/internal/domain"
)

type StoreInterface interface {
	// Read devuelve un dentista por su id
	Read(id int) (domain.Dentist, error)
	// Create agrega un nuevo dentista
	Create(dentist domain.Dentist) error
	// Update actualiza un dentista
	Update(dentist domain.Dentist) error
	// Delete elimina un dentista
	Delete(id int) error
	// Read devuelve todos los dentista

	// Exists verifica si un dentista existe
	Exists(codeValue string) bool
}