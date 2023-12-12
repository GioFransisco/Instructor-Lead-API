package repositorymock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"github.com/stretchr/testify/mock"
)

type AuthMock struct {
	mock.Mock
}

func (m *AuthMock) Login(username string) (model.User, error) {
	args := m.Called(username)

	return args.Get(0).(model.User), args.Error(1)
}

func (m *AuthMock) Register(payloadUser model.User) (model.User, error) {
	args := m.Called(payloadUser)

	return args.Get(0).(model.User), args.Error(1)
}
