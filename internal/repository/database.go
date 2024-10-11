package repository

import (
	"github.com/angellllk/task-management/internal/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type TaskRepository struct {
	DB *gorm.DB
}

func New() (*TaskRepository, error) {
	dsn := "host=localhost user=postgres password=password dbname=tasks port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	errAM := db.AutoMigrate(&models.TaskDB{})
	if errAM != nil {
		return nil, errAM
	}

	return &TaskRepository{DB: db}, nil
}

func (r *TaskRepository) Create(task models.TaskAPI) (string, error) {
	taskDB := &models.TaskDB{
		ID:          uuid.NewString(),
		Title:       task.Title,
		Description: task.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	return taskDB.ID, r.DB.Create(taskDB).Error
}

func (r *TaskRepository) Fetch(id string) (models.TaskDB, error) {
	var task models.TaskDB
	err := r.DB.Where("id = ?", id).First(&task).Error
	return task, err
}

func (r *TaskRepository) FetchAll() ([]models.TaskDB, error) {
	var tasks []models.TaskDB
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) Update(id string, update models.TaskAPI) error {
	var task models.TaskDB
	if err := r.DB.Where("id = ?", id).First(&task).Error; err != nil {
		return err
	}
	if err := r.DB.Model(&task).Select("Title", "Description").Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) Delete(id string) error {
	return r.DB.Where("id = ?", id).Delete(&models.TaskDB{}).Error
}
