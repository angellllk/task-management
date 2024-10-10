package service

import (
	"github.com/angellllk/task-management/internal/models"
	"github.com/angellllk/task-management/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func (s *TaskService) Create(task models.Task) error {
	return s.Repo.Create(task)
}

func (s *TaskService) Fetch(id string) (models.Task, error) {
	return s.Repo.Fetch(id)
}

func (s *TaskService) FetchAll() ([]models.Task, error) {
	return s.Repo.FetchAll()
}

func (s *TaskService) Update(id string, update models.Task) error {
	return s.Repo.Update(id, update)
}

func (s *TaskService) Delete(id string) error {
	return s.Repo.Delete(id)
}
