package usecasemock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/mock"
)

type AbsencesUseCaseMock struct {
	mock.Mock
}

func (a *AbsencesUseCaseMock) CreateNewAbsence(payload model.Absences) (dto.AbsencesResponseDto, error) {
	args := a.Called(payload)
	return args.Get(0).(dto.AbsencesResponseDto), args.Error(1)
}

func (a *AbsencesUseCaseMock) FindAbsenceById(scheduleDetailId string) (model.GetAbsences, error) {
	args := a.Called(scheduleDetailId)
	return args.Get(0).(model.GetAbsences), args.Error(1)
}
