package turn

import (
	"github.com/desafioFinalBack/internal/domain"
	"github.com/desafioFinalBack/pkg/store"
	"errors"
)

type Repository interface {
	GetByID(id int) (domain.Turn, error)
	GetByDNI(dni int) (domain.Turn, error)
	Create(turn domain.Turn) (domain.Turn, error)
	Update(id int, turn domain.Turn) (domain.Turn, error)
	Delete(id int) error
}

type repository struct {
	Storage store.StoreInterfaceTurn
}

func NewRepository(storage store.StoreInterfaceTurn) Repository {
	return &repository{
		Storage: storage,
	}
}

// GetByID busca un turno por su id
func (r *repository) GetByID(id int) (domain.Turn, error) {
	turn, err := r.Storage.ReadTurn(id)
	if err != nil {
		return domain.Turn{}, errors.New("turn not found")
	}
	return turn, nil
}

// GetByDNI busca un turno por su dni
func (r *repository) GetByDNI(dni int) (domain.Turn, error) {
	turn, err := r.Storage.ReadTurnByDni(dni)
	if err != nil {
		return domain.Turn{}, errors.New("turn not found")
	}
	return turn, nil
}

// Create crea un nuevo turno
func (r *repository) Create(turn domain.Turn) (domain.Turn, error) {
	err := r.Storage.CreateTurn(turn)
	if err != nil {
		return domain.Turn{}, errors.New("error creating turn")
	}
	return turn, nil
}

// Update actualiza un turno
func (r *repository) Update(id int, turn domain.Turn) (domain.Turn, error) {
	turn.Id = id
	err := r.Storage.UpdateTurn(turn)
	if err != nil {
		return domain.Turn{}, errors.New("error updating turn")
	}
	return turn, nil
}


// Delete elimina un paciente
func (r *repository) Delete(id int) error {
	err := r.Storage.DeleteTurn(id)
	if err != nil {
		return err
	}
	return nil
}