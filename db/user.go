package db

import "time"

type UserModel struct {
	ID           int
	FirstName    string
	LastName     string
	Phone        string
	Email        string
	Prescription string
	Active       string
	CreatedAt    time.Time
	ModifiedAt   time.Time
}

func CreateUser(m UserModel) (UserModel, error) {
	nm := UserModel{}
	q := "INSERT INTO users (first_name, last_name, phone, email, prescription, active) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, first_name, last_name, phone, email, prescription, active, created_at, modified_at"
	row := DB.QueryRow(q, m.FirstName, m.LastName, m.Phone, m.Email, m.Prescription, m.Active)
	err := row.Scan(&nm.ID, &nm.FirstName, &nm.LastName, &nm.Phone, &nm.Email, &nm.Prescription, &nm.Active, &nm.CreatedAt, &nm.ModifiedAt)
	if err != nil {
		return nm, err
	}
	return nm, nil
}

func ShowUser(m UserModel) (UserModel, error) {
	nm := UserModel{}
	stmt := "SELECT * FROM users WHERE id=$1"
	err := DB.QueryRow(stmt, m.ID).Scan(&nm.ID, &nm.FirstName, &nm.LastName, &nm.Phone, &nm.Email, &nm.Prescription, &nm.Active, &nm.CreatedAt, &nm.ModifiedAt)
	if err != nil {
		return nm, err
	}
	return nm, nil
}

func ListUsers() ([]UserModel, error) {
	msl := []UserModel{}
	q := "SELECT * FROM users"
	rows, err := DB.Query(q)
	if err != nil {
		return msl, err
	}
	for rows.Next() {
		nm := UserModel{}
		err = rows.Scan(&nm.ID, &nm.FirstName, &nm.LastName, &nm.Phone, &nm.Email, &nm.Prescription, &nm.Active, &nm.CreatedAt, &nm.ModifiedAt)
		if err != nil {
			return msl, err
		}
		msl = append(msl, nm)
	}
	return msl, nil
}

func DeleteUser(m UserModel) (UserModel, error) {
	nm := UserModel{}
	stmt := "DELETE FROM users WHERE id=$1 RETURNING id, first_name, last_name, phone, email, prescription, active, created_at, modified_at"
	err := DB.QueryRow(stmt, m.ID).Scan(&nm.ID, &nm.FirstName, &nm.LastName, &nm.Phone, &nm.Email, &nm.Prescription, &nm.Active, &nm.CreatedAt, &nm.ModifiedAt)
	if err != nil {
		return nm, err
	}
	return nm, nil
}
