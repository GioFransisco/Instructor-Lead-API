package usecase

import (
	"errors"
	"testing"
	"time"

	bcryptmock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/bcrypt_mock"
	repositorymock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/repository_mock"
	usecasemock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/usecase_mock"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	urm     *repositorymock.UserRepoMock
	userUcM *usecasemock.UserUCMock
	bm      *bcryptmock.BcryptMock
	userUC  UserUC
}

func (s *UserUsecaseTestSuite) SetupTest() {
	s.urm = new(repositorymock.UserRepoMock)
	s.userUcM = new(usecasemock.UserUCMock)
	s.bm = new(bcryptmock.BcryptMock)
	s.userUC = NewUserUC(s.urm)
}

func TestAuthRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

func (s *UserUsecaseTestSuite) TestChangePaswordUser_Success() {
	var mockUserUpdate = model.User{
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

	s.userUcM.On("FindById", mockUserUpdate.Id).Return(mockUserUpdate, nil)

	s.urm.On("Get", mockUserUpdate.Id).Return(mockUserUpdate, nil)

	s.bm.On("GenerateFromPassword", mock.AnythingOfType("model.User.Password"), bcrypt.DefaultCost).Return([]byte("$2a$10$l6k9FVAgy1u4z6UNtlqd.OSN.gF2VKsL9U4oGLPqUA/q9PIVPP18m"), nil)

	s.urm.On("UpdatePasword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockUserUpdate, nil)

	_, err := s.userUC.ChangePaswordUser(mockUserUpdate.Password, mockUserUpdate.Id)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
}

func (s *UserUsecaseTestSuite) TestChangePaswordUser_FailGet() {
	var mockUserUpdate = model.User{
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

	s.userUcM.On("FindById", mockUserUpdate.Id).Return(model.User{}, errors.New("error"))

	s.urm.On("Get", mockUserUpdate.Id).Return(model.User{}, errors.New("error"))

	_, err := s.userUC.ChangePaswordUser(mockUserUpdate.Password, mockUserUpdate.Id)
	assert.Error(s.T(), err)
}

func (s *UserUsecaseTestSuite) TestUpdateUser_Success() {
	var mockUserUpdate = model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar78@gmail.com",
		PhoneNumber: "0811775468198",
		Username:    "amarholo",
		Age:         18,
		Address:     "Indramayu jakartasdisug",
		Gander:      "L",
	}

	var mockDtoUser = dto.UserUpdateDto{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar78@gmail.com",
		PhoneNumber: "0811775468198",
		Username:    "amarholo",
		Age:         18,
		Address:     "Indramayu jakartasdisug",
		Gender:      "L",
	}

	s.userUcM.On("FindById", mockUserUpdate.Id).Return(mockUserUpdate, nil)

	s.urm.On("Get", mockUserUpdate.Id).Return(mockUserUpdate, nil)

	s.urm.On("GetUserByEmail", mockUserUpdate.Email).Return(mockUserUpdate, errors.New("error"))

	s.urm.On("UpdateUser", mockUserUpdate).Return(mockUserUpdate, nil)

	_, err := s.userUC.UpdateUser(mockDtoUser)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
}
