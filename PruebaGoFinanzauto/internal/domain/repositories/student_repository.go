package repositories

import (
	"PruebaGoFinanzauto/internal/domain/models"
)

type StudentRepository interface {
	Create(student *models.Student) error
	FindAll(filters map[string]interface{}) ([]models.Student, error)
	FindByID(id uint) (*models.Student, error)
	Update(student *models.Student) error
	Delete(id uint) error
}
