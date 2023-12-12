package jwtmock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

type JwtMock struct {
	mock.Mock
}

func (m *JwtMock) GenereteToken(userId, email, role string) (model.TokenModel, error) {
	args := m.Called(userId, email, role)

	return args.Get(0).(model.TokenModel), args.Error(1)
}

func (m *JwtMock) VerifyToken(token model.TokenModel) (jwt.MapClaims, error) {
	args := m.Called(token)

	return args.Get(0).(jwt.MapClaims), args.Error(1)
}
