package service

import (
	"github.com/angellllk/task-management/internal/models"
	"github.com/angellllk/task-management/internal/repository"
	"github.com/google/uuid"
	"time"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func (s *TaskService) Create(dto models.TaskCreateDTO) (*models.TaskResponseDTO, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	taskDB := models.TaskDB{
		ID:          uuid.NewString(),
		Title:       dto.Title,
		Description: dto.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	err := s.Repo.Create(taskDB)
	if err != nil {
		return nil, err
	}

	// Return the response DTO
	return taskDB.ToDTO(), nil
}

func (s *TaskService) Fetch(id string) (*models.TaskResponseDTO, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, err
	}

	taskDB, err := s.Repo.Fetch(id)
	if err != nil {
		return nil, err
	}

	return taskDB.ToDTO(), nil
}

func (s *TaskService) FetchAll() ([]models.TaskResponseDTO, error) {
	tasks, err := s.Repo.FetchAll()
	if err != nil {
		return nil, err
	}

	var tasksDTO []models.TaskResponseDTO
	for _, t := range tasks {
		tasksDTO = append(tasksDTO, *t.ToDTO())
	}
	return tasksDTO, nil
}

func (s *TaskService) Update(id string, dto models.TaskUpdateDTO) (*models.TaskResponseDTO, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, err
	}
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	taskDB := models.TaskDB{
		ID:          id,
		Title:       dto.Title,
		Description: dto.Description,
		Completed:   false,
		CreatedAt:   time.Time{},
	}
	if err := s.Repo.Update(taskDB); err != nil {
		return nil, err
	}

	return taskDB.ToDTO(), nil
}

func (s *TaskService) Delete(id string) error {
	if err := uuid.Validate(id); err != nil {
		return err
	}
	return s.Repo.Delete(id)
}
