package usecase

import (
	"errors"
	"fmt"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/repository"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"golang.org/x/crypto/bcrypt"
)

type UserUC interface {
	UpdateUser(dto.UserUpdateDto) (model.User, error)
	FindUserByEmail(string) (model.User, error)
	ChangePaswordUser(password, id string) (model.User, error)
	DeleteUserById(id string) error
	FindById(id string) (model.User, error)
}

type userUC struct {
	repo repository.UserRepository
}

func (u *userUC) FindById(id string) (model.User, error) {
	user, err := u.repo.Get(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil

}

// ChangePaswordUser implements UserUC.
func (u *userUC) ChangePaswordUser(password string, id string) (model.User, error) {
	_, err := u.FindById(id)

	if err != nil {
		return model.User{}, err
	}

	if len(password) < 8 {
		return model.User{}, common.InvalidError{Message: "password must be more than 8 characters"}
	}

	hassPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return model.User{}, common.InvalidError{Message: "failed generate password"}
	}

	return u.repo.UpdatePasword(string(hassPass), id)
}

func (u *userUC) UpdateUser(dtoPayload dto.UserUpdateDto) (model.User, error) {
	_, err := u.FindById(dtoPayload.Id)

	if err != nil {
		return model.User{}, err
	}

	payload := model.User{
		Id:          dtoPayload.Id,
		Name:        dtoPayload.Name,
		Email:       dtoPayload.Email,
		PhoneNumber: dtoPayload.PhoneNumber,
		Username:    dtoPayload.Username,
		Age:         dtoPayload.Age,
		Address:     dtoPayload.Address,
		Gander:      dtoPayload.Gender,
	}

	if payload.Name == "" && payload.Email == "" && payload.PhoneNumber == "" && payload.Username == "" && payload.Gander == "" && payload.Age <= 17 && payload.Address == "" {
		return model.User{}, errors.New("at least 1 value is updated")
	}

	if payload.Name != "" {
		if len(payload.Name) < 4 {
			return model.User{}, errors.New("name must be more than 4 characters")
		}
	}

	if payload.Email != "" {
		if len(payload.Email) < 8 {
			return model.User{}, errors.New("email invalid")
		}
	}

	if payload.Password != "" {
		if len(payload.Password) < 8 {
			return model.User{}, errors.New("password must be more than 8 characters")
		}
	}

	if payload.Address != "" {
		if len(payload.Address) < 10 {
			return model.User{}, errors.New("input full address")
		}
	}

	if payload.Age != 0 {
		if payload.Age < 18 {
			return model.User{}, errors.New("insufficient age")
		}
	}

	if payload.Gander != "" {
		if !payload.IsValidGender() {
			return model.User{}, errors.New("gender is not valid")
		}
	}

	if payload.PhoneNumber != "" {
		if len(payload.PhoneNumber) <= 9 {
			return model.User{}, errors.New("phone number is not valid")
		}
	}

	if payload.Username != "" {
		if len(payload.Username) <= 2 {
			return model.User{}, errors.New("username too short")
		}
	}

	if _, err := u.FindUserByEmail(payload.Email); err == nil {
		return model.User{}, errors.New("email already exist")
	}

	user, err := u.repo.UpdateUser(payload)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userUC) FindUserByEmail(email string) (model.User, error) {
	if email == "" {
		return model.User{}, errors.New("email required")
	}

	return u.repo.GetUserByEmail(email)
}

// DeleteUserById implements UserUC.
func (u *userUC) DeleteUserById(id string) error {
	return u.repo.DeleteUser(id)
}

func NewUserUC(repo repository.UserRepository) UserUC {
	return &userUC{repo: repo}
}
