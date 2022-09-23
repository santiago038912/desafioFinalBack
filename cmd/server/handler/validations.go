package handler

import (
	"errors"
	"github.com/desafioFinalBack/internal/domain"
)

// validateEmptysDentist valida que los campos no esten vacios
func validateEmptysDentist(dentist *domain.Dentist) (bool, error) {
	if dentist.Name == "" || dentist.LastName == "" || dentist.RegisterNumber == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}


// validateEmptysPatient valida que los campos no esten vacios
func validateEmptysPatient(patient *domain.Patient) (bool, error) {
	if patient.Name == "" || patient.LastName == "" || patient.Address == "" || patient.DNI == 0 || patient.Date == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// validateEmptysTurn valida que los campos no esten vacios
func validateEmptysTurn(turn *domain.Turn) (bool, error) {
	if turn.Dentist.RegisterNumber == "" || turn.Patient.DNI == 0 || turn.Date == "" || turn.Time == "" || turn.Description == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}