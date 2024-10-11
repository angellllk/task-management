package service

import (
	"github.com/angellllk/task-management/internal/models"
	"github.com/angellllk/task-management/internal/repository"
	"github.com/google/uuid"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func (s *TaskService) Create(task models.TaskAPI) (string, error) {
	if err := task.Validate(); err != nil {
		return "", err
	}

	return s.Repo.Create(task)
}

func (s *TaskService) Fetch(id string) (models.TaskDB, error) {
	if err := uuid.Validate(id); err != nil {
		return models.TaskDB{}, err
	}
	return s.Repo.Fetch(id)
}

func (s *TaskService) FetchAll() ([]models.TaskDB, error) {
	return s.Repo.FetchAll()
}

func (s *TaskService) Update(id string, update models.TaskAPI) error {
	if err := uuid.Validate(id); err != nil {
		return err
	}
	if err := update.Validate(); err != nil {
		return err
	}
	return s.Repo.Update(id, update)
}

func (s *TaskService) Delete(id string) error {
	if err := uuid.Validate(id); err != nil {
		return err
	}
	return s.Repo.Delete(id)
}
