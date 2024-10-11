package repository

import (
	"github.com/angellllk/task-management/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (r *TaskRepository) Create(task models.TaskDB) error {
	return r.DB.Create(task).Error
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

func (r *TaskRepository) Update(update models.TaskDB) error {
	var task models.TaskDB
	if err := r.DB.Where("id = ?", update.ID).First(&task).Error; err != nil {
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
