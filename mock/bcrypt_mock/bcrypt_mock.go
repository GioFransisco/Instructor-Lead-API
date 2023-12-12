package bcryptmock

import "github.com/stretchr/testify/mock"

type BcryptMock struct {
	mock.Mock
}

func (m *BcryptMock) CompareHashAndPassword(hasspassword, password []byte) error {
	args := m.Called(hasspassword, password)
	return args.Error(1)
}

func (m *BcryptMock) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	args := m.Called(password)
	return args.Get(0).([]byte), args.Error(1)
}
