package database

import (
	"PruebaGoFinanzauto/internal/domain/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if host == "" {
		host = "localhost"
	}
	if user == "" {
		user = "postgres"
	}
	if password == "" {
		password = "1234"
	}
	if dbname == "" {
		dbname = "school_db"
	}
	if port == "" {
		port = "5432"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbname, port)

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&models.Student{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) CreateStudent(student *models.Student) error {
	return r.db.Create(student).Error
}

func (r *PostgresRepository) FindAllStudents(filters map[string]interface{}) ([]models.Student, error) {
	var students []models.Student
	query := r.db

	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s LIKE ?", key), fmt.Sprintf("%%%v%%", value))
	}

	err := query.Find(&students).Error
	return students, err
}

func (r *PostgresRepository) FindStudentByID(id uint) (*models.Student, error) {
	var student models.Student
	err := r.db.First(&student, id).Error
	return &student, err
}

func (r *PostgresRepository) UpdateStudent(student *models.Student) error {
	return r.db.Save(student).Error
}

func (r *PostgresRepository) DeleteStudent(id uint) error {
	return r.db.Delete(&models.Student{}, id).Error
}
