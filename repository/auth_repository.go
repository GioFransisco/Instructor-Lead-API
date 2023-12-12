package repository

import (
	"database/sql"
	"errors"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
)

type AuthRepository interface {
	Login(username string) (model.User, error)
	Register(model.User) (model.User, error)
}

type authRepository struct {
	db *sql.DB
}

// Login implements AuthRepository.
func (r *authRepository) Login(user string) (payloadUser model.User, err error) {
	err = r.db.QueryRow(utilsmodel.UserLogin, user).Scan(&payloadUser.Id, &payloadUser.Password, &payloadUser.Email, &payloadUser.Role)

	if err != nil {
		return payloadUser, errors.New("user not found")
	}

	return
}

// Register implements AuthRepository.
func (r *authRepository) Register(payload model.User) (model.User, error) {
	err := r.db.QueryRow(utilsmodel.UserRegist, payload.Name, payload.Email, payload.PhoneNumber, payload.Username, payload.Password, payload.Age, payload.
		Address, payload.Gander, payload.Role, time.Now()).Scan(&payload.Id, &payload.CreatedAt, &payload.UpdatedAt)

	if err != nil {
		return model.User{}, errors.New("make sure all data is filled in correctly")
	}

	return payload, err
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db}
}
