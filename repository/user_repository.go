package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
)

type UserRepository interface {
	UpdateUser(model.User) (model.User, error)
	UpdatePasword(password, id string) (model.User, error)
	GetUserByEmail(string) (model.User, error)
	DeleteUser(id string) error
	Get(id string) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

// Get implements UserRepository.
func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE id = $1`
	err := u.db.QueryRow(query, id).Scan(
		&user.Id, &user.Name, &user.Email, &user.PhoneNumber,
		&user.Username, &user.Password, &user.Age, &user.Address, &user.Gander,
		&user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// UpdatePasword implements UserRepository.
func (r *userRepository) UpdatePasword(password string, id string) (payload model.User, err error) {
	tx, err := r.db.Begin()

	if err != nil {
		return model.User{}, err
	}

	err = tx.QueryRow(utilsmodel.UserChangePassword, password, time.Now(), id).Scan(&payload.Id, &payload.Name, &payload.Email, &payload.PhoneNumber, &payload.Username, &payload.Age, &payload.Address, &payload.Gander, &payload.Role, &payload.CreatedAt, &payload.UpdatedAt)

	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}

// UpdateUser implements UserRepository.
func (r *userRepository) UpdateUser(payload model.User) (model.User, error) {
	tx, _ := r.db.Begin()
	index := 1
	var value []any
	qry := ""

	if payload.Name != "" {
		qry += "name=$" + strconv.Itoa(index)
		value = append(value, payload.Name)
		index++
	}

	if payload.Email != "" {
		if index > 1 {
			qry += ",email=$" + strconv.Itoa(index)
			index++
		} else {
			qry += "email=$" + strconv.Itoa(index)
			index++
		}
		value = append(value, payload.Email)
	}

	if payload.PhoneNumber != "" {
		if index > 1 {
			qry += ",phone_number=$" + strconv.Itoa(index)
			index++
		} else {
			qry += "phone_number=$" + strconv.Itoa(index)
			index++
		}
		value = append(value, payload.PhoneNumber)
	}

	if payload.Username != "" {
		if index > 1 {
			qry += ",username=$" + strconv.Itoa(index)
			index++
		} else {
			qry += "username=$" + strconv.Itoa(index)
			index++
		}
		value = append(value, payload.Username)
	}

	if payload.Password != "" {
		if index > 1 {
			qry += ",password=$" + strconv.Itoa(index)
			index++
		} else {
			qry += "password=$" + strconv.Itoa(index)
			index++
		}
		value = append(value, payload.Password)
	}

	if payload.Age != 0 {
		if index > 1 {
			qry += ",age=$" + strconv.Itoa(index)
			index++
		} else {
			qry += "age=$" + strconv.Itoa(index)
			index++
		}
		value = append(value, payload.Age)
	}

	if payload.Gander != "" {
		if index > 1 {
			qry += ",gender=$" + strconv.Itoa(index)
			index++
		} else {
			qry += "gender=$" + strconv.Itoa(index)
			index++
		}
		value = append(value, payload.Gander)
	}

	if payload.Address != "" {
		if index > 1 {
			qry += ",address=$" + strconv.Itoa(index)
			index++
		} else {
			qry += "address=$" + strconv.Itoa(index)
			index++
		}
		value = append(value, payload.Address)
	}

	qry += ",updated_at=$" + strconv.Itoa(index)
	value = append(value, time.Now())
	index++
	qry += " Where id=$" + strconv.Itoa(index) + utilsmodel.UserUpdateReturning
	value = append(value, payload.Id)

	err := tx.QueryRow(utilsmodel.UserUpdate+qry, value...).Scan(&payload.Id, &payload.Name, &payload.Email, &payload.PhoneNumber, &payload.Username, &payload.Age, &payload.Gander, &payload.Address, &payload.CreatedAt, &payload.UpdatedAt)

	if err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	tx.Commit()

	return payload, nil
}

func (r *userRepository) GetUserByEmail(email string) (payload model.User, err error) {
	err = r.db.QueryRow(utilsmodel.UserGetByEmail, email).Scan(&payload.Id, &payload.Name, &payload.Email, &payload.PhoneNumber, &payload.Username, &payload.Age, &payload.Address, &payload.Gander, &payload.Role)

	if payload.Id == "" {
		return model.User{}, errors.New("users not found")
	}

	return
}

func (u *userRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := u.db.Exec(query, id)
	return err
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}
