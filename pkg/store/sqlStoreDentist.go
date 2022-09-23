package store

import (
	"database/sql"
	"github.com/desafioFinalBack/internal/domain"
)

type SqlStoreDentist struct {
	DB *sql.DB
}

func NewSqlStoreDentist(db *sql.DB) StoreInterfaceDentist {
	return &SqlStoreDentist{
		DB: db,
	}
}

// Read devuelve un dentista por su id
func (s *SqlStoreDentist) ReadDentist(id int) (domain.Dentist, error) {
	var dentist domain.Dentist
	row := s.DB.QueryRow("SELECT * FROM dentists WHERE id = ?", id)
	err := row.Scan(&dentist.Id, &dentist.Name, &dentist.LastName, &dentist.RegisterNumber)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *SqlStoreDentist) CreateDentist(dentist domain.Dentist) error {
	query := "INSERT INTO dentists (id, name, last_name, register_number) VALUES (?, ?, ?, ?)"
	st, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := st.Exec(&dentist.Id, &dentist.Name, &dentist.LastName, &dentist.RegisterNumber)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStoreDentist) UpdateDentist(dentist domain.Dentist) error {
	stmt, err := s.DB.Prepare("UPDATE dentists SET name = ?, last_name = ?, register_number = ? WHERE id = ?")
	if err != nil {
		return err
	}
	
	res, err := stmt.Exec(&dentist.Name, &dentist.LastName, &dentist.RegisterNumber, &dentist.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStoreDentist) DeleteDentist(id int) error {
	stmt := "DELETE FROM dentists WHERE id = ?"
	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}