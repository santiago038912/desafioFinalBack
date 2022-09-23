package turn

import (
	"github.com/desafioFinalBack/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Turn, error)
	GetByDNI(dni int) (domain.Turn, error)
	Create(turn domain.Turn) (domain.Turn, error)
	Update(id int, turn domain.Turn) (domain.Turn, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

// GetByID busca un turno por su id
func (s *service) GetByID(id int) (domain.Turn, error) {
	t, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turn{}, err
	}
	return t, nil
}

// GetByDNI busca un turno por su dni
func (s *service) GetByDNI(dni int) (domain.Turn, error) {
	t, err := s.r.GetByDNI(dni)
	if err != nil {
		return domain.Turn{}, err
	}
	return t, nil
}

// Create crea un nuevo turno
func (s *service) Create(turn domain.Turn) (domain.Turn, error) {
	turn, err := s.r.Create(turn)
	if err != nil {
		return domain.Turn{}, err
	}
	return turn, nil
}

// Update actualiza un turno
func (s *service) Update(id int, turn domain.Turn) (domain.Turn, error) {
	turn, err := s.r.Update(id, turn)
	if err != nil {
		return domain.Turn{}, err
	}
	return turn, nil
}

// Delete elimina un turno
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}