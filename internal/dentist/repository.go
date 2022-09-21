package dentist

import (
	"github.com/desafioFinalBack/internal/domain"
	store "github.com/desafioFinalBack/pkg/storeDentists"
	"errors"
)

type Repository interface {
	GetByID(id int) (domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type repository struct {
	Storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{
		Storage: storage,
	}
}

// GetByID busca un dentista por su id
func (r *repository) GetByID(id int) (domain.Dentist, error) {
	dentist, err := r.Storage.Read(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")
	}
	return dentist, nil

}

// Create crea un nuevo dentista
func (r *repository) Create(dentist domain.Dentist) (domain.Dentist, error) {
	err := r.Storage.Create(dentist)
	if err != nil {
		return domain.Dentist{}, errors.New("error creating dentist")
	}
	return dentist, nil
}

// Update actualiza un dentista
func (r *repository) Update(id int, dentist domain.Dentist) (domain.Dentist, error) {
	dentist.Id = id
	err := r.Storage.Update(dentist)
	if err != nil {
		return domain.Dentist{}, errors.New("error updating dentist")
	}
	return dentist, nil
}

// Delete elimina un dentista
func (r *repository) Delete(id int) error {
	err := r.Storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
