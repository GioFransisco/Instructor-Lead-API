package manager

import "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/repository"

type RepoManager interface {
	StackRepository() repository.StackRepository
	AuthRepository() repository.AuthRepository
	UserRepository() repository.UserRepository
	NoteRepository() repository.NoteRepository


	QuestionRepository() repository.QuestionRepository
	ScheduleRepository() repository.ScheduleRepository
	AbsenceRepository() repository.AbsencesRepository
	ScheduleApproveRepository() repository.ScheduleApproveRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) NoteRepository() repository.NoteRepository {
	return repository.NewNoteRepository(r.infra.Connection())
}


// QuestionRepository implements RepoManager.
func (r *repoManager) QuestionRepository() repository.QuestionRepository {
	return repository.NewQusetionRepository(r.infra.Connection())
}

func (r *repoManager) StackRepository() repository.StackRepository {
	return repository.NewStackRepository(r.infra.Connection())
}

func (r *repoManager) AuthRepository() repository.AuthRepository {
	return repository.NewAuthRepository(r.infra.Connection())
}

func (r *repoManager) UserRepository() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Connection())
}

func (r *repoManager) ScheduleRepository() repository.ScheduleRepository {
	return repository.NewScheduleRepository(r.infra.Connection())
}

func (r *repoManager) AbsenceRepository() repository.AbsencesRepository {
	return repository.NewAbsencesRepository(r.infra.Connection())
}

func (r *repoManager) ScheduleApproveRepository() repository.ScheduleApproveRepository {
	return repository.NewScheduleApproveRepository(r.infra.Connection())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
