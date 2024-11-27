package usecases

import (
	"PruebaGoFinanzauto/internal/domain/models"
	"errors"
)

type StudentRepository interface {
	CreateStudent(student *models.Student) error
	FindAllStudents(filters map[string]interface{}) ([]models.Student, error)
	FindStudentByID(id uint) (*models.Student, error)
	UpdateStudent(student *models.Student) error
	DeleteStudent(id uint) error
}

type StudentUseCase struct {
	repo StudentRepository
}

func NewStudentUseCase(repo StudentRepository) *StudentUseCase {
	return &StudentUseCase{
		repo: repo,
	}
}

func (uc *StudentUseCase) CreateStudent(student *models.Student) error {
	if student.FirstName == "" || student.LastName == "" {
		return errors.New("first name and last name are required")
	}

	if student.Email == "" {
		return errors.New("email is required")
	}

	return uc.repo.CreateStudent(student)
}

func (uc *StudentUseCase) GetStudents(filters map[string]interface{}) ([]models.Student, error) {
	return uc.repo.FindAllStudents(filters)
}

func (uc *StudentUseCase) GetStudentByID(id uint) (*models.Student, error) {
	student, err := uc.repo.FindStudentByID(id)
	if err != nil {
		return nil, errors.New("student not found")
	}
	return student, nil
}

func (uc *StudentUseCase) UpdateStudent(id uint, student *models.Student) error {
	existingStudent, err := uc.repo.FindStudentByID(id)
	if err != nil {
		return errors.New("student not found")
	}

	if student.FirstName == "" || student.LastName == "" {
		return errors.New("first name and last name are required")
	}

	if student.Email == "" {
		return errors.New("email is required")
	}

	//existingStudent.ID = student.ID
	existingStudent.FirstName = student.FirstName
	existingStudent.LastName = student.LastName
	existingStudent.Email = student.Email
	existingStudent.DateOfBirth = student.DateOfBirth

	return uc.repo.UpdateStudent(existingStudent)
}

func (uc *StudentUseCase) DeleteStudent(id uint) error {
	_, err := uc.repo.FindStudentByID(id)
	if err != nil {
		return errors.New("student not found")
	}

	return uc.repo.DeleteStudent(id)
}
