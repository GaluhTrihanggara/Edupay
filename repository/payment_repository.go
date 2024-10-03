package repository

import (
	"Edupay/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// PaymentRepository adalah interface untuk operasi CRUD pada entitas Payment
type PaymentRepository interface {
	GetAllPaymentsRepository(page, limit int, status string) ([]*model.Payment, error)
	GetPaymentByIDRepository(id string) (*model.Payment, error)
	GetPaymentsByBillIDRepository(billID string) ([]*model.Payment, error)
	CreatePaymentRepository(payment *model.Payment) (*model.Payment, error)
	UpdatePaymentByIDRepository(id string, payment *model.Payment) (*model.Payment, error)
	DeletePaymentByIDRepository(id string) error
}

// paymentRepository adalah struct yang mengimplementasikan PaymentRepository
type paymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository membuat instance baru dari paymentRepository
func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{db}
}

// GetAllPaymentsRepository mengambil semua pembayaran dengan pagination dan pencarian berdasarkan status
func (r *paymentRepository) GetAllPaymentsRepository(page, limit int, status string) ([]*model.Payment, error) {
	var payments []*model.Payment
	offset := (page - 1) * limit

	query := r.db.Offset(offset).Limit(limit)
	if status != "" {
		query = query.Where("status LIKE ?", "%"+status+"%")
	}

	result := query.Order("payment_date DESC").Find(&payments)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting payments: %s", result.Error)
	}
	return payments, nil
}

// GetPaymentByIDRepository mengambil pembayaran berdasarkan ID
func (r *paymentRepository) GetPaymentByIDRepository(id string) (*model.Payment, error) {
	var payment model.Payment
	result := r.db.First(&payment, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("payment with ID %s not found", id)
		}
		return nil, fmt.Errorf("error getting payment with ID %s: %s", id, result.Error)
	}
	return &payment, nil
}

// GetPaymentsByBillIDRepository mengambil semua pembayaran berdasarkan BillID
func (r *paymentRepository) GetPaymentsByBillIDRepository(billID string) ([]*model.Payment, error) {
	var payments []*model.Payment
	result := r.db.Where("bill_id = ?", billID).Order("payment_date DESC").Find(&payments)
	if result.Error != nil {
		return nil, result.Error
	}
	return payments, nil
}

// CreatePaymentRepository membuat pembayaran baru
func (r *paymentRepository) CreatePaymentRepository(payment *model.Payment) (*model.Payment, error) {
	result := r.db.Create(payment)
	if result.Error != nil {
		return nil, result.Error
	}
	return payment, nil
}

// UpdatePaymentByIDRepository memperbarui data pembayaran berdasarkan ID
func (r *paymentRepository) UpdatePaymentByIDRepository(id string, payment *model.Payment) (*model.Payment, error) {
	result := r.db.Model(&model.Payment{}).Where("id = ?", id).Updates(payment)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("payment not found")
	}
	return payment, nil
}

// DeletePaymentByIDRepository menghapus pembayaran berdasarkan ID
func (r *paymentRepository) DeletePaymentByIDRepository(id string) error {
	result := r.db.Delete(&model.Payment{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("payment not found")
	}
	return nil
}
