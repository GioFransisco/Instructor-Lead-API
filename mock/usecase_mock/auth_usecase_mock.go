package usecasemock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/mock"
)

type AuthUserMock struct {
	mock.Mock
}

func (m *AuthUserMock) FindByUsername(user dto.UserLoginDto) (model.TokenModel, error) {
	args := m.Called(user)

	return args.Get(0).(model.TokenModel), args.Error(1)
}

func (m *AuthUserMock) CreateNewUser(user model.User) (model.User, error) {
	args := m.Called(user)

	return args.Get(0).(model.User), args.Error(1)
}
