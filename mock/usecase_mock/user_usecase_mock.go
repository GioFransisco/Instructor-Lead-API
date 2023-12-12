package usecasemock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/mock"
)

type UserUCMock struct {
	mock.Mock
}

func (m *UserUCMock) UpdateUser(user dto.UserUpdateDto) (model.User, error) {
	args := m.Called(user)

	return args.Get(0).(model.User), args.Error(1)
}

func (m *UserUCMock) FindUserByEmail(email string) (model.User, error) {
	args := m.Called(email)

	return args.Get(0).(model.User), args.Error(1)
}

func (m *UserUCMock) ChangePaswordUser(password, id string) (model.User, error) {
	args := m.Called(password, id)

	return args.Get(0).(model.User), args.Error(1)
}

func (m *UserUCMock) DeleteUserById(id string) error {
	args := m.Called(id)

	return args.Error(0)
}

func (m *UserUCMock) FindById(id string) (model.User, error) {
	args := m.Called(id)

	return args.Get(0).(model.User), args.Error(1)
}
