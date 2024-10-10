package service

import (
	"github.com/angellllk/task-management/internal/models"
	"github.com/angellllk/task-management/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func (s *TaskService) Create(task models.CreateTaskAPI) (string, error) {
	if err := task.Validate(); err != nil {
		return "", err
	}

	return s.Repo.Create(task)
}

func (s *TaskService) Fetch(id string) (models.TaskDB, error) {
	return s.Repo.Fetch(id)
}

func (s *TaskService) FetchAll() ([]models.TaskDB, error) {
	return s.Repo.FetchAll()
}

func (s *TaskService) Update(id string, update models.TaskDB) error {
	return s.Repo.Update(id, update)
}

func (s *TaskService) Delete(id string) error {
	return s.Repo.Delete(id)
}
