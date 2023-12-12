package usecase

import (
	"database/sql"
	"fmt"
	"strings"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/repository"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
)

type StackUseCase interface {
	RegisterNewStack(payload dto.StackRequestDto) (model.Stack, error)
	FindAll() ([]model.Stack, error)
	FindByID(id string) (model.Stack, error)
	UpdateStack(id string, payload model.Stack) (model.Stack, error)
	DeleteStack(id string) error
}

type stackUseCase struct {
	repo repository.StackRepository
}

func (s *stackUseCase) RegisterNewStack(payload dto.StackRequestDto) (model.Stack, error) {
	var stack model.Stack

	stack.Name = payload.Name

	return s.repo.Create(stack)
}

func (s *stackUseCase) FindAll() ([]model.Stack, error) {
	stacks, err := s.repo.List()
	if err != nil {
		return []model.Stack{}, err
	}

	if len(stacks) < 1 {
		return []model.Stack{}, common.InvalidError{Message: "data not found"}
	}

	return stacks, nil
}

func NewStackUseCase(repo repository.StackRepository) StackUseCase {
	return &stackUseCase{repo: repo}
}

func (s *stackUseCase) FindByID(id string) (model.Stack, error) {
	stack, err := s.repo.FindByID(id)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "uuid") || err == sql.ErrNoRows:
			return model.Stack{}, common.InvalidError{Message: fmt.Sprintf("stack with ID %s not found", id)}
		default:
			return model.Stack{}, err
		}
	}

	return stack, nil
}

func (s *stackUseCase) UpdateStack(id string, payload model.Stack) (model.Stack, error) {
	if payload.Status != "" && !payload.IsValidStatus() {
		return model.Stack{}, common.InvalidError{Message: "invalid status, make sure status is one of 'Active' or 'Inactive'"}
	}

	if payload.Name == "" && payload.Status == "" {
		return model.Stack{}, common.InvalidError{Message: "at least 1 value is updated"}
	}

	stack, err := s.repo.Update(id, payload)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "uuid") || err == sql.ErrNoRows:
			return model.Stack{}, common.InvalidError{Message: fmt.Sprintf("stack with ID %s not found", id)}
		default:
			return model.Stack{}, err
		}
	}

	return stack, err
}

func (s *stackUseCase) DeleteStack(id string) error {
	return s.repo.Delete(id)
}
