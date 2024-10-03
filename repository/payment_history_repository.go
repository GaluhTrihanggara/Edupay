package repository

import (
	"Edupay/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// PaymentHistoryRepository adalah interface untuk operasi CRUD pada entitas PaymentHistory
type PaymentHistoryRepository interface {
	GetAllHistoriesRepository(page, limit int) ([]*model.PaymentHistory, error)
	GetHistoryByIDRepository(id string) (*model.PaymentHistory, error)
	GetHistoriesByParentIDRepository(parentID string) ([]*model.PaymentHistory, error)
	GetHistoriesByPaymentIDRepository(paymentID string) ([]*model.PaymentHistory, error)
	CreateHistoryRepository(history *model.PaymentHistory) (*model.PaymentHistory, error)
	UpdateHistoryByIDRepository(id string, history *model.PaymentHistory) (*model.PaymentHistory, error)
	DeleteHistoryByIDRepository(id string) error
}

// paymentHistoryRepository adalah struct yang mengimplementasikan PaymentHistoryRepository
type paymentHistoryRepository struct {
	db *gorm.DB
}

// NewPaymentHistoryRepository membuat instance baru dari paymentHistoryRepository
func NewPaymentHistoryRepository(db *gorm.DB) *paymentHistoryRepository {
	return &paymentHistoryRepository{db}
}

// GetAllHistoriesRepository mengambil semua riwayat pembayaran dengan pagination
func (r *paymentHistoryRepository) GetAllHistoriesRepository(page, limit int) ([]*model.PaymentHistory, error) {
	var histories []*model.PaymentHistory
	offset := (page - 1) * limit

	result := r.db.Offset(offset).Limit(limit).Order("payment_date DESC").Find(&histories)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting payment histories: %s", result.Error)
	}
	return histories, nil
}

// GetHistoryByIDRepository mengambil riwayat pembayaran berdasarkan ID
func (r *paymentHistoryRepository) GetHistoryByIDRepository(id string) (*model.PaymentHistory, error) {
	var history model.PaymentHistory
	result := r.db.First(&history, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("payment history with ID %s not found", id)
		}
		return nil, fmt.Errorf("error getting payment history with ID %s: %s", id, result.Error)
	}
	return &history, nil
}

// GetHistoriesByParentIDRepository mengambil semua riwayat pembayaran berdasarkan ParentID
func (r *paymentHistoryRepository) GetHistoriesByParentIDRepository(parentID string) ([]*model.PaymentHistory, error) {
	var histories []*model.PaymentHistory
	result := r.db.Where("parent_id = ?", parentID).Order("payment_date DESC").Find(&histories)
	if result.Error != nil {
		return nil, result.Error
	}
	return histories, nil
}

// GetHistoriesByPaymentIDRepository mengambil semua riwayat pembayaran berdasarkan PaymentID
func (r *paymentHistoryRepository) GetHistoriesByPaymentIDRepository(paymentID string) ([]*model.PaymentHistory, error) {
	var histories []*model.PaymentHistory
	result := r.db.Where("payment_id = ?", paymentID).Order("payment_date DESC").Find(&histories)
	if result.Error != nil {
		return nil, result.Error
	}
	return histories, nil
}

// CreateHistoryRepository membuat entri riwayat pembayaran baru
func (r *paymentHistoryRepository) CreateHistoryRepository(history *model.PaymentHistory) (*model.PaymentHistory, error) {
	result := r.db.Create(history)
	if result.Error != nil {
		return nil, result.Error
	}
	return history, nil
}

// UpdateHistoryByIDRepository memperbarui riwayat pembayaran berdasarkan ID
func (r *paymentHistoryRepository) UpdateHistoryByIDRepository(id string, history *model.PaymentHistory) (*model.PaymentHistory, error) {
	result := r.db.Model(&model.PaymentHistory{}).Where("id = ?", id).Updates(history)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("payment history not found")
	}
	return history, nil
}

// DeleteHistoryByIDRepository menghapus riwayat pembayaran berdasarkan ID
func (r *paymentHistoryRepository) DeleteHistoryByIDRepository(id string) error {
	result := r.db.Delete(&model.PaymentHistory{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("payment history not found")
	}
	return nil
}
