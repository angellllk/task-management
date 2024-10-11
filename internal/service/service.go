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
	return &models.TaskResponseDTO{
		ID:          taskDB.ID,
		Title:       taskDB.Title,
		Description: taskDB.Description,
		Completed:   taskDB.Completed,
		CreatedAt:   taskDB.CreatedAt,
	}, nil
}

func (s *TaskService) Fetch(id string) (*models.TaskResponseDTO, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, err
	}

	taskDB, err := s.Repo.Fetch(id)
	if err != nil {
		return nil, err
	}

	return &models.TaskResponseDTO{
		ID:          taskDB.ID,
		Title:       taskDB.Title,
		Description: taskDB.Description,
		Completed:   taskDB.Completed,
		CreatedAt:   taskDB.CreatedAt,
	}, nil
}

func (s *TaskService) FetchAll() ([]models.TaskDB, error) {
	return s.Repo.FetchAll()
}

func (s *TaskService) Update(id string, dto models.TaskUpdateDTO) error {
	if err := uuid.Validate(id); err != nil {
		return err
	}
	if err := dto.Validate(); err != nil {
		return err
	}

	taskDB := models.TaskDB{
		ID:          id,
		Title:       dto.Title,
		Description: dto.Description,
		Completed:   false,
		CreatedAt:   time.Time{},
	}
	// Call repository to update the task
	return s.Repo.Update(taskDB)
}

func (s *TaskService) Delete(id string) error {
	if err := uuid.Validate(id); err != nil {
		return err
	}
	return s.Repo.Delete(id)
}
