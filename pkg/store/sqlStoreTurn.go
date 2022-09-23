package store

import (
	"database/sql"

	"github.com/desafioFinalBack/internal/domain"
)

type SqlStoreTurn struct {
	DB *sql.DB
}

func NewSqlStoreTurn(db *sql.DB) StoreInterfaceTurn {
	return &SqlStoreTurn{
		DB: db,
	}
}

// ReadTurn devuelve un turno por su id
func (s *SqlStoreTurn) ReadTurn(id int) (domain.Turn, error) {
	var turn domain.Turn
	row := s.DB.QueryRow("SELECT turns.id, turns.date, turns.time, turns.description, dentists.id, dentists.name, dentists.last_name, dentists.register_number, patients.id, patients.name, patients.last_name, patients.address, patients.dni, patients.date FROM turns JOIN dentists ON dentists.register_number = turns.dentist_register_number JOIN patients ON patients.dni = turns.patient_dni WHERE turns.id = ?;", id)

	err := row.Scan(&turn.Id, &turn.Date, &turn.Time, &turn.Description,
		&turn.Dentist.Id, &turn.Dentist.Name, &turn.Dentist.LastName, &turn.Dentist.RegisterNumber,
		&turn.Patient.Id, &turn.Patient.Name, &turn.Patient.LastName, &turn.Patient.Address, &turn.Patient.DNI, &turn.Patient.Date)
	if err != nil {
		return domain.Turn{}, err
	}
	return turn, nil
}

// ReadTurnByDni devuelve un turno por su dni
func (s *SqlStoreTurn) ReadTurnByDni(dni int) (domain.Turn, error) {
	var turn domain.Turn
	row := s.DB.QueryRow("SELECT turns.id, turns.date, turns.time, turns.description, dentists.id, dentists.name, dentists.last_name, dentists.register_number, patients.id, patients.name, patients.last_name, patients.address, patients.dni, patients.date FROM turns JOIN dentists ON dentists.register_number = turns.dentist_register_number JOIN patients ON patients.dni = turns.patient_dni WHERE turns.patient_dni = ?;", dni)

	err := row.Scan(&turn.Id, &turn.Date, &turn.Time, &turn.Description,
		&turn.Dentist.Id, &turn.Dentist.Name, &turn.Dentist.LastName, &turn.Dentist.RegisterNumber,
		&turn.Patient.Id, &turn.Patient.Name, &turn.Patient.LastName, &turn.Patient.Address, &turn.Patient.DNI, &turn.Patient.Date)
	if err != nil {
		return domain.Turn{}, err
	}

	return turn, nil
}

// CreateTurn crea un turno
func (s *SqlStoreTurn) CreateTurn(turn domain.Turn) error {
	query := "INSERT INTO turns (id, dentist_register_number, patient_dni, date, time, description) VALUES (?, ?, ?, ?, ?, ?)"
	st, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := st.Exec(&turn.Id ,&turn.Dentist.RegisterNumber, &turn.Patient.DNI, &turn.Date, &turn.Time, &turn.Description)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// UpdateTurn devuelve un turno por su id
func (s *SqlStoreTurn) UpdateTurn(turn domain.Turn) error {
	stmt, err := s.DB.Prepare("UPDATE turns SET date = ?, time = ?, description = ? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&turn.Date, &turn.Time, &turn.Description, &turn.Id)
	if err != nil {
		return err
	}
	return nil
}

// DeletePatient turno un paciente
func (s *SqlStoreTurn) DeleteTurn(id int) error {
	stmt := "DELETE FROM turns WHERE id = ?"
	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
