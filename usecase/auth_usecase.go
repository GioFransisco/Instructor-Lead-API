package usecase

import (
	"errors"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/repository"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	FindByUsername(dto.UserLoginDto) (model.TokenModel, error)
	CreateNewUser(model.User) (model.User, error)
}

type authUsecase struct {
	repo   repository.AuthRepository
	jwt    common.JwtToken
	userUC UserUC
}

// CreateNewUser implements AuthUsecase.
func (u *authUsecase) CreateNewUser(payload model.User) (model.User, error) {
	hassPass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)

	if err != nil {
		return model.User{}, errors.New("failed to generate password")
	}

	payload.Role = "Participant"

	if len(payload.Name) < 4 {
		return model.User{}, errors.New("name must be more than 4 characters")
	}

	if len(payload.Email) < 8 {
		return model.User{}, errors.New("email invalid")
	}

	if len(payload.Password) < 8 {
		return model.User{}, errors.New("password must be more than 8 characters")
	}

	if len(payload.Address) < 10 {
		return model.User{}, errors.New("input full address")
	}

	if payload.Age < 18 {
		return model.User{}, errors.New("insufficient age")
	}

	if !payload.IsValidGender() {
		return model.User{}, errors.New("gender is not valid")
	}

	if !payload.IsValidRole() {
		return model.User{}, errors.New("role is not valid")
	}

	if len(payload.PhoneNumber) <= 9 {
		return model.User{}, errors.New("phone number is not valid")
	}

	if len(payload.Username) <= 2 {
		return model.User{}, errors.New("username too short")
	}

	if _, err := u.userUC.FindUserByEmail(payload.Email); err == nil {
		return model.User{}, errors.New("email already exist")
	}

	payload.Password = string(hassPass)

	payload, err = u.repo.Register(payload)

	if err != nil {
		return model.User{}, err
	}

	payload.Password = ""

	return payload, nil
}

// FindByUsername implements AuthUsecase.
func (u *authUsecase) FindByUsername(userDto dto.UserLoginDto) (model.TokenModel, error) {
	user, err := u.repo.Login(userDto.Username)

	if err != nil {
		return model.TokenModel{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password)); err != nil {
		return model.TokenModel{}, errors.New("password not match")
	}

	token, err := u.jwt.GenereteToken(user.Id, user.Email, user.Role)

	if err != nil {
		return token, err
	}

	return token, nil
}

func NewAuthUseCase(repo repository.AuthRepository, jwt common.JwtToken, userUC UserUC) AuthUsecase {
	return &authUsecase{
		repo:   repo,
		jwt:    jwt,
		userUC: userUC,
	}
}
