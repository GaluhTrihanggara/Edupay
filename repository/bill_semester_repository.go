package repository

import (
	"Edupay/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// BillSemesterRepository adalah interface untuk operasi CRUD pada entitas BillSemester
type BillSemesterRepository interface {
	GetAllBillsRepository(page, limit int, semester, year string) ([]*model.BillSemester, error)
	GetBillByIDRepository(id string) (*model.BillSemester, error)
	CreateBillRepository(bill *model.BillSemester) (*model.BillSemester, error)
	UpdateBillByIDRepository(id string, bill *model.BillSemester) (*model.BillSemester, error)
	DeleteBillByIDRepository(id string) error
	GetBillsByStudentIDRepository(studentID string) ([]*model.BillSemester, error)
}

// billSemesterRepository adalah struct yang mengimplementasikan BillSemesterRepository
type billSemesterRepository struct {
	db *gorm.DB
}

// NewBillSemesterRepository membuat instance baru dari billSemesterRepository
func NewBillSemesterRepository(db *gorm.DB) *billSemesterRepository {
	return &billSemesterRepository{db}
}

// GetAllBillsRepository mengambil semua tagihan semester dengan pagination dan pencarian berdasarkan semester dan tahun
func (r *billSemesterRepository) GetAllBillsRepository(page, limit int, semester, year string) ([]*model.BillSemester, error) {
	var bills []*model.BillSemester
	offset := (page - 1) * limit

	query := r.db.Offset(offset).Limit(limit)
	if semester != "" {
		query = query.Where("semester LIKE ?", "%"+semester+"%")
	}
	if year != "" {
		query = query.Where("year = ?", year)
	}

	result := query.Order("created_at DESC").Find(&bills)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting bills: %s", result.Error)
	}
	return bills, nil
}

// GetBillByIDRepository mengambil tagihan semester berdasarkan ID
func (r *billSemesterRepository) GetBillByIDRepository(id string) (*model.BillSemester, error) {
	var bill model.BillSemester
	result := r.db.First(&bill, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("bill with ID %s not found", id)
		}
		return nil, fmt.Errorf("error getting bill with ID %s: %s", id, result.Error)
	}
	return &bill, nil
}

// CreateBillRepository membuat tagihan semester baru
func (r *billSemesterRepository) CreateBillRepository(bill *model.BillSemester) (*model.BillSemester, error) {
	result := r.db.Create(bill)
	if result.Error != nil {
		return nil, result.Error
	}
	return bill, nil
}

// UpdateBillByIDRepository memperbarui tagihan semester berdasarkan ID
func (r *billSemesterRepository) UpdateBillByIDRepository(id string, bill *model.BillSemester) (*model.BillSemester, error) {
	result := r.db.Model(&model.BillSemester{}).Where("id = ?", id).Updates(bill)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("bill not found")
	}
	return bill, nil
}

// DeleteBillByIDRepository menghapus tagihan semester berdasarkan ID
func (r *billSemesterRepository) DeleteBillByIDRepository(id string) error {
	result := r.db.Delete(&model.BillSemester{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("bill not found")
	}
	return nil
}

// GetBillsByStudentIDRepository mengambil semua tagihan semester berdasarkan StudentID
func (r *billSemesterRepository) GetBillsByStudentIDRepository(studentID string) ([]*model.BillSemester, error) {
	var bills []*model.BillSemester
	result := r.db.Where("student_id = ?", studentID).Find(&bills)
	if result.Error != nil {
		return nil, result.Error
	}
	return bills, nil
}
