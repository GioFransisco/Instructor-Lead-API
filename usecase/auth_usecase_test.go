package usecase

import (
	"errors"
	"testing"
	"time"

	bcryptmock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/bcrypt_mock"
	jwtmock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/jwt_mock"
	repositorymock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/repository_mock"
	usecasemock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/usecase_mock"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecaseTestSuite struct {
	suite.Suite
	urm      *repositorymock.AuthMock
	jwtToken *jwtmock.JwtMock
	userUc   *usecasemock.UserUCMock
	bm       *bcryptmock.BcryptMock
	auth     AuthUsecase
}

func (s *AuthUsecaseTestSuite) SetupTest() {
	s.urm = new(repositorymock.AuthMock)
	s.jwtToken = new(jwtmock.JwtMock)
	s.userUc = new(usecasemock.UserUCMock)
	s.bm = new(bcryptmock.BcryptMock)
	s.auth = NewAuthUseCase(s.urm, s.jwtToken, s.userUc)
}

var mockUser = model.User{
	Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
	Name:        "amar",
	Email:       "amar78@gmail.com",
	PhoneNumber: "0811775468198",
	Username:    "amarholo",
	Password:    "$2a$10$Db.RTaGfd87D3.AZmwsXvef3OiALHTFEIAmP0DBZ.RB8uvB3PviiO",
	Age:         18,
	Address:     "Indramayu jakartasdisug",
	Gander:      "L",
	Role:        "Participant",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

var mockTokenModel = model.TokenModel{
	Token: "ldsalldjpwp93802",
}

var mockLoginDto = dto.UserLoginDto{
	Username: "amarholo",
	Password: "password",
}

func TestBillUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUsecaseTestSuite))
}

func (m *AuthUsecaseTestSuite) TestFindByUsername_Success() {
	m.urm.On("Login", mockUser.Username).Return(mockUser, nil)

	m.bm.On("CompareHashAndPassword", []byte(mockUser.Password), []byte(mockLoginDto.Password)).Return(nil)

	m.jwtToken.On("GenereteToken", mockUser.Id, mockUser.Email, mockUser.Role).Return(mockTokenModel, nil)

	_, err := m.auth.FindByUsername(mockLoginDto)
	assert.Nil(m.T(), err)
	assert.NoError(m.T(), err)
}

func (m *AuthUsecaseTestSuite) TestFindByUsername_FailedLogin() {
	m.urm.On("Login", mockUser.Username).Return(model.User{}, errors.New("error"))

	_, err := m.auth.FindByUsername(mockLoginDto)
	assert.Error(m.T(), err)
}

func (m *AuthUsecaseTestSuite) TestFindByUsername_FailedGenerete() {
	m.urm.On("Login", mockUser.Username).Return(mockUser, nil)
	m.bm.On("CompareHashAndPassword", []byte(mockUser.Password), []byte(mockLoginDto.Password)).Return(errors.New("error"))
	m.jwtToken.On("GenereteToken", mockUser.Id, mockUser.Email, mockUser.Role).Return(mockTokenModel, errors.New("error"))

	_, err := m.auth.FindByUsername(mockLoginDto)
	assert.Error(m.T(), err)
}

func (m *AuthUsecaseTestSuite) TestCreateNewUser_Success() {
	var mockUserRegist = model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar78@gmail.com",
		PhoneNumber: "0811775468198",
		Username:    "amarholo",
		Password:    "password",
		Age:         18,
		Address:     "Indramayu jakartasdisug",
		Gander:      "L",
		Role:        "Participant",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.bm.On("GenerateFromPassword", mock.AnythingOfType("model.User.Password"), bcrypt.DefaultCost).Return([]byte("$2a$10$KcaqASjKTFXs8U2Wnpxa4.MfeiJqLWbN9UaflInO9YeLNZBk5u42m"), nil)

	m.userUc.On("FindUserByEmail", mockUserRegist.Email).Return(model.User{}, errors.New("users not found"))

	m.urm.On("Register", mock.AnythingOfType("model.User")).Return(mockUserRegist, nil)

	_, err := m.auth.CreateNewUser(mockUserRegist)
	assert.Nil(m.T(), err)
	assert.NoError(m.T(), err)
}

func (m *AuthUsecaseTestSuite) TestCreateNewUser_Fail() {
	var mockUserRegist = model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar78",
		PhoneNumber: "0811775468198",
		Username:    "amarholo",
		Password:    "password",
		Age:         18,
		Address:     "Indramayu jakartasdisug",
		Gander:      "L",
		Role:        "Participant",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.bm.On("GenerateFromPassword", mock.AnythingOfType("model.User.Password"), bcrypt.DefaultCost).Return([]byte("$2a$10$KcaqASjKTFXs8U2Wnpxa4.MfeiJqLWbN9UaflInO9YeLNZBk5u42m"), nil)

	m.userUc.On("FindUserByEmail", mockUserRegist.Email).Return(mockUserRegist, nil)

	m.urm.On("Register", mock.AnythingOfType("model.User")).Return(mockUserRegist, nil)

	_, err := m.auth.CreateNewUser(mockUserRegist)
	assert.Error(m.T(), err)
}

func (m *AuthUsecaseTestSuite) TestCreateNewUser_FailPassword() {
	var mockUserRegist = model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar78&aijsasoai",
		PhoneNumber: "0811775468198",
		Username:    "amarholo",
		Password:    "passwor",
		Age:         18,
		Address:     "Indramayu jakartasdisug",
		Gander:      "L",
		Role:        "Participant",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.bm.On("GenerateFromPassword", mock.AnythingOfType("model.User.Password"), bcrypt.DefaultCost).Return([]byte("$2a$10$KcaqASjKTFXs8U2Wnpxa4.MfeiJqLWbN9UaflInO9YeLNZBk5u42m"), nil)

	m.userUc.On("FindUserByEmail", mockUserRegist.Email).Return(mockUserRegist, nil)

	m.urm.On("Register", mock.AnythingOfType("model.User")).Return(mockUserRegist, nil)

	_, err := m.auth.CreateNewUser(mockUserRegist)
	assert.Error(m.T(), err)
}

func (m *AuthUsecaseTestSuite) TestCreateNewUser_FailAddress() {
	var mockUserRegist = model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar7809w15hss",
		PhoneNumber: "0811775468198",
		Username:    "amarholo",
		Password:    "password",
		Age:         18,
		Address:     "Indrama",
		Gander:      "L",
		Role:        "Participant",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.bm.On("GenerateFromPassword", mock.AnythingOfType("model.User.Password"), bcrypt.DefaultCost).Return([]byte("$2a$10$KcaqASjKTFXs8U2Wnpxa4.MfeiJqLWbN9UaflInO9YeLNZBk5u42m"), nil)

	m.userUc.On("FindUserByEmail", mockUserRegist.Email).Return(mockUserRegist, nil)

	m.urm.On("Register", mock.AnythingOfType("model.User")).Return(mockUserRegist, nil)

	_, err := m.auth.CreateNewUser(mockUserRegist)
	assert.Error(m.T(), err)
}

func (m *AuthUsecaseTestSuite) TestCreateNewUser_FailGender() {
	var mockUserRegist = model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar7809w15hss",
		PhoneNumber: "0811775468198",
		Username:    "amarholo",
		Password:    "password",
		Age:         18,
		Address:     "Indramayusdosiopopos",
		Gander:      "X",
		Role:        "Participant",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.bm.On("GenerateFromPassword", mock.AnythingOfType("model.User.Password"), bcrypt.DefaultCost).Return([]byte("$2a$10$KcaqASjKTFXs8U2Wnpxa4.MfeiJqLWbN9UaflInO9YeLNZBk5u42m"), nil)

	m.userUc.On("FindUserByEmail", mockUserRegist.Email).Return(mockUserRegist, nil)

	m.urm.On("Register", mock.AnythingOfType("model.User")).Return(mockUserRegist, nil)

	_, err := m.auth.CreateNewUser(mockUserRegist)
	assert.Error(m.T(), err)
}

func (m *AuthUsecaseTestSuite) TestCreateNewUser_FailPhoneNumber() {
	var mockUserRegist = model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar7809w15hss",
		PhoneNumber: "081177",
		Username:    "amarholo",
		Password:    "password",
		Age:         18,
		Address:     "Indramasdskldsakdlsaka",
		Gander:      "L",
		Role:        "Participant",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.bm.On("GenerateFromPassword", mock.AnythingOfType("model.User.Password"), bcrypt.DefaultCost).Return([]byte("$2a$10$KcaqASjKTFXs8U2Wnpxa4.MfeiJqLWbN9UaflInO9YeLNZBk5u42m"), nil)

	m.userUc.On("FindUserByEmail", mockUserRegist.Email).Return(mockUserRegist, nil)

	m.urm.On("Register", mock.AnythingOfType("model.User")).Return(mockUserRegist, nil)

	_, err := m.auth.CreateNewUser(mockUserRegist)
	assert.Error(m.T(), err)
}

func (m *AuthUsecaseTestSuite) TestCreateNewUser_FailUsername() {
	var mockUserRegist = model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar7809w15hss",
		PhoneNumber: "0811775468198",
		Username:    "a",
		Password:    "password",
		Age:         18,
		Address:     "Indramasdskldsakdlsaka",
		Gander:      "L",
		Role:        "Participant",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.bm.On("GenerateFromPassword", mock.AnythingOfType("model.User.Password"), bcrypt.DefaultCost).Return([]byte("$2a$10$KcaqASjKTFXs8U2Wnpxa4.MfeiJqLWbN9UaflInO9YeLNZBk5u42m"), nil)

	m.userUc.On("FindUserByEmail", mockUserRegist.Email).Return(mockUserRegist, nil)

	m.urm.On("Register", mock.AnythingOfType("model.User")).Return(mockUserRegist, nil)

	_, err := m.auth.CreateNewUser(mockUserRegist)
	assert.Error(m.T(), err)
}

func (m *AuthUsecaseTestSuite) TestCreateNewUser_FailAge() {
	var mockUserRegist = model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar7809w15hss",
		PhoneNumber: "0811775468198",
		Username:    "amarholo",
		Password:    "password",
		Age:         10,
		Address:     "Indramasdskldsakdlsaka",
		Gander:      "L",
		Role:        "Participant",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.bm.On("GenerateFromPassword", mock.AnythingOfType("model.User.Password"), bcrypt.DefaultCost).Return([]byte("$2a$10$KcaqASjKTFXs8U2Wnpxa4.MfeiJqLWbN9UaflInO9YeLNZBk5u42m"), nil)

	m.userUc.On("FindUserByEmail", mockUserRegist.Email).Return(mockUserRegist, nil)

	m.urm.On("Register", mock.AnythingOfType("model.User")).Return(mockUserRegist, nil)

	_, err := m.auth.CreateNewUser(mockUserRegist)
	assert.Error(m.T(), err)
}
