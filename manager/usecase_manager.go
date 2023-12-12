package manager

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/usecase"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
)

type UseCaseManager interface {
	StackUseCase() usecase.StackUseCase
	AuthUsecase() usecase.AuthUsecase
	UserUsecase() usecase.UserUC
	NoteUseCase() usecase.NoteUseCase
	ScheduleUseCase() usecase.ScheduleUseCase
	QuestionUsecase() usecase.QuestionUsecase
	AbsenceUseCase() usecase.AbsencesUseCase
	ScheduleApproveUseCase() usecase.ScheduleApproveUseCase
}

type useCaseManager struct {
	repo RepoManager
	jwt  common.JwtToken
}

// QuestionUsecase implements UseCaseManager.
func (u *useCaseManager) QuestionUsecase() usecase.QuestionUsecase {
	return usecase.NewQusetionUsecase(u.repo.QuestionRepository(), u.ScheduleUseCase(), u.UserUsecase())
}

// AuthRepository implements UseCaseManager.
func (u *useCaseManager) AuthUsecase() usecase.AuthUsecase {
	return usecase.NewAuthUseCase(u.repo.AuthRepository(), u.jwt, u.UserUsecase())
}

// UserRepository implements UseCaseManager.
func (u *useCaseManager) UserUsecase() usecase.UserUC {
	return usecase.NewUserUC(u.repo.UserRepository())
}

func (u *useCaseManager) StackUseCase() usecase.StackUseCase {
	return usecase.NewStackUseCase(u.repo.StackRepository())
}

func (u *useCaseManager) NoteUseCase() usecase.NoteUseCase {
	return usecase.NewNoteUseCase(u.repo.NoteRepository())
}


func (u *useCaseManager) ScheduleUseCase() usecase.ScheduleUseCase {
	return usecase.NewScheduleUseCase(u.repo.ScheduleRepository(), u.UserUsecase(), u.StackUseCase())
}

// AbsenceUseCase implements UseCaseManager.
func (u *useCaseManager) AbsenceUseCase() usecase.AbsencesUseCase {
	return usecase.NewAbsencesUseCase(u.repo.AbsenceRepository(), u.ScheduleUseCase(), u.UserUsecase())
}

func (u *useCaseManager) ScheduleApproveUseCase() usecase.ScheduleApproveUseCase {
	return usecase.NewScheduleApproveUseCase(u.repo.ScheduleApproveRepository(), u.ScheduleUseCase())
}

func NewUseCaseManager(repo RepoManager, jwt common.JwtToken) UseCaseManager {
	return &useCaseManager{repo: repo, jwt: jwt}
}
