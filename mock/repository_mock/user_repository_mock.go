package repositorymock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) UpdateUser(payloadUser model.User) (model.User, error) {
	args := m.Called(payloadUser)

	return args.Get(0).(model.User), args.Error(1)
}

func (m *UserRepoMock) UpdatePasword(password, id string) (model.User, error) {
	args := m.Called(password, id)

	return args.Get(0).(model.User), args.Error(1)
}

func (m *UserRepoMock) GetUserByEmail(email string) (model.User, error) {
	args := m.Called(email)

	return args.Get(0).(model.User), args.Error(1)
}

func (m *UserRepoMock) DeleteUser(id string) error {
	args := m.Called(id)

	return args.Error(0)
}

func (m *UserRepoMock) Get(id string) (model.User, error) {
	args := m.Called(id)

	return args.Get(0).(model.User), args.Error(1)
}
