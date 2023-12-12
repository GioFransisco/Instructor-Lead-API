package repository

import (
	"database/sql"
	"strconv"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
)

type StackRepository interface {
	Create(payload model.Stack) (model.Stack, error)
	List() ([]model.Stack, error)
	FindByID(id string) (model.Stack, error)
	Update(id string, payload model.Stack) (model.Stack, error)
	Delete(id string) error
}

type stackRepository struct {
	db *sql.DB
}

func (s *stackRepository) Create(payload model.Stack) (model.Stack, error) {
	var stack model.Stack

	db, err := s.db.Prepare(utilsmodel.StackCreate)
	if err != nil {
		return model.Stack{}, err
	}

	defer db.Close()

	err = db.QueryRow(payload.Name, utilsmodel.DefaultStatus, time.Now()).Scan(&stack.Id, &stack.CreatedAt, &stack.UpdatedAt)
	if err != nil {
		return model.Stack{}, err
	}

	stack.Name = payload.Name
	stack.Status = utilsmodel.DefaultStatus

	return stack, nil
}

func (s *stackRepository) List() ([]model.Stack, error) {
	var stacks []model.Stack

	db, err := s.db.Prepare(utilsmodel.StackList)
	if err != nil {
		return []model.Stack{}, err
	}

	defer db.Close()

	rows, err := db.Query()
	if err != nil {
		return []model.Stack{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var stack model.Stack
		if err := rows.Scan(&stack.Id, &stack.Name, &stack.Status, &stack.CreatedAt, &stack.UpdatedAt); err != nil {
			return []model.Stack{}, err
		}

		stacks = append(stacks, stack)
	}

	return stacks, nil
}

func NewStackRepository(db *sql.DB) StackRepository {
	return &stackRepository{db: db}
}

func (s *stackRepository) FindByID(id string) (model.Stack, error) {
	var stack model.Stack

	row := s.db.QueryRow(utilsmodel.StackFindById, id)
	err := row.Scan(&stack.Id, &stack.Name, &stack.Status, &stack.CreatedAt, &stack.UpdatedAt)
	if err != nil {
		return model.Stack{}, err
	}

	return stack, nil
}

func (s *stackRepository) Update(id string, payload model.Stack) (model.Stack, error) {
	var stack model.Stack

	index := 1
	var value []any
	qry := ""

	qry += utilsmodel.StackUpdate

	if payload.Name != "" {
		qry += "name=$" + strconv.Itoa(index)
		value = append(value, payload.Name)
		index++
	}

	if payload.Status != "" {
		if index > 1 {
			qry += ",status=$" + strconv.Itoa(index)
			index++
		} else {
			qry += "status=$" + strconv.Itoa(index)
			index++
		}

		value = append(value, payload.Status)
	}

	qry += ",updated_at=$" + strconv.Itoa(index)
	value = append(value, time.Now())
	index++

	qry += " Where id=$" + strconv.Itoa(index) + utilsmodel.StackUpdateReturning
	value = append(value, id)

	err := s.db.QueryRow(qry, value...).Scan(&stack.Id, &stack.Name, &stack.Status, &stack.CreatedAt, &stack.UpdatedAt)

	if err != nil {
		return model.Stack{}, err
	}

	return stack, nil
}

func (s *stackRepository) Delete(id string) error {
	_, err := s.db.Exec(utilsmodel.DeleteStackById, id)
	return err
}
