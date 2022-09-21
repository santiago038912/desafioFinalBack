package store

import (
	"database/sql"
	"github.com/desafioFinalBack/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &SqlStore{
		DB: db,
	}
}

// Read devuelve un dentista por su id
func (s *SqlStore) Read(id int) (domain.Dentist, error) {
	var dentist domain.Dentist
	row := s.DB.QueryRow("SELECT * FROM dentists WHERE id = ?", id)
	err := row.Scan(&dentist.Id, &dentist.Name, &dentist.LastName, &dentist.RegisterNumber)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *SqlStore) Create(dentist domain.Dentist) error {
	query := "INSERT INTO dentists (id, name, last_name, register_number) VALUES (?, ?, ?, ?)"
	st, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := st.Exec(dentist.Id, dentist.Name, dentist.LastName, dentist.RegisterNumber)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Update(dentist domain.Dentist) error {
	stmt, err := s.DB.Prepare("UPDATE dentists SET name = ?, last_name = ?, register_number = ? WHERE id = ?")
	if err != nil {
		return err
	}
	
	_, err = stmt.Exec(dentist.Name, dentist.LastName, dentist.RegisterNumber, dentist.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Delete(id int) error {
	stmt := "DELETE FROM dentists WHERE id = ?"
	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}


// Exists verifica si un dentista existe
func (s *SqlStore) Exists(codeValue string) bool {
	var id int
	row := s.DB.QueryRow("select id from products where code_value = ?", codeValue)
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	if id > 0 {
		return true
	}

	return false
}