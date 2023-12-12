package repositorymock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"github.com/stretchr/testify/mock"
)

type AbsencesRepoMock struct {
	mock.Mock
}

func (a *AbsencesRepoMock) Create(payload model.Absences) (model.Absences, error) {
	args := a.Called(payload)
	return args.Get(0).(model.Absences), args.Error(1)
}

func (a *AbsencesRepoMock) GetScheduleDetailId(scheduleDetailId string) (model.GetAbsences, error) {
	args := a.Called(scheduleDetailId)
	return args.Get(0).(model.GetAbsences), args.Error(1)
}
