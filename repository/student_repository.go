package repository

import (
	"Edupay/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// StudentRepository adalah interface untuk operasi CRUD pada entitas Student
type StudentRepository interface {
	GetAllStudentsRepository(page, limit int, name, class string) ([]*model.Student, error)
	GetStudentByIDRepository(id string) (*model.Student, error)
	CreateStudentRepository(student *model.Student) (*model.Student, error)
	UpdateStudentByIDRepository(id string, student *model.Student) (*model.Student, error)
	DeleteStudentByIDRepository(id string) error
}

// studentRepository adalah struct yang mengimplementasikan StudentRepository
type studentRepository struct {
	db *gorm.DB
}

// NewStudentRepository membuat instance baru dari studentRepository
func NewStudentRepository(db *gorm.DB) *studentRepository {
	return &studentRepository{db}
}

// GetAllStudentsRepository mengambil semua siswa dengan pagination dan pencarian berdasarkan nama dan class
func (r *studentRepository) GetAllStudentsRepository(page, limit int, name, class string) ([]*model.Student, error) {
	var students []*model.Student
	offset := (page - 1) * limit

	query := r.db.Offset(offset).Limit(limit).Preload("Parent")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if class != "" {
		query = query.Where("class = ?", class)
	}

	result := query.Order("created_at DESC").Find(&students)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting students: %s", result.Error)
	}
	return students, nil
}

// GetStudentByIDRepository mengambil siswa berdasarkan ID
func (r *studentRepository) GetStudentByIDRepository(id string) (*model.Student, error) {
	var student model.Student
	result := r.db.Preload("Parent").First(&student, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("student with ID %s not found", id)
		}
		return nil, fmt.Errorf("error getting student with ID %s: %s", id, result.Error)
	}
	return &student, nil
}

// CreateStudentRepository membuat entri siswa baru di database
func (r *studentRepository) CreateStudentRepository(student *model.Student) (*model.Student, error) {
	result := r.db.Create(student)
	if result.Error != nil {
		return nil, result.Error
	}
	return student, nil
}

// UpdateStudentByIDRepository memperbarui data siswa berdasarkan ID
func (r *studentRepository) UpdateStudentByIDRepository(id string, student *model.Student) (*model.Student, error) {
	result := r.db.Model(&model.Student{}).Where("id = ?", id).Updates(student)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("student not found")
	}
	return student, nil
}

// DeleteStudentByIDRepository menghapus siswa berdasarkan ID
func (r *studentRepository) DeleteStudentByIDRepository(id string) error {
	result := r.db.Delete(&model.Student{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("student not found")
	}
	return nil
}
