package repository

import (
	"Edupay/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// TransactionRepository adalah interface untuk operasi CRUD pada entitas Transaction
type TransactionRepository interface {
	GetAllTransactionsRepository(page, limit int, parentId, itemId string) ([]*model.Transaction, error)
	GetTransactionByIDRepository(id string) (*model.Transaction, error)
	CreateTransactionRepository(transaction *model.Transaction) (*model.Transaction, error)
	UpdateTransactionByIDRepository(id string, transaction *model.Transaction) (*model.Transaction, error)
	DeleteTransactionByIDRepository(id string) error
	GetTransactionsByParentIDRepository(parentID string) ([]*model.Transaction, error)
	GetTransactionsByItemIDRepository(itemID string) ([]*model.Transaction, error)
}

// transactionRepository adalah struct yang mengimplementasikan TransactionRepository
type transactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository membuat instance baru dari transactionRepository
func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

// GetAllTransactionsRepository mengambil semua transaksi dengan pagination dan pencarian berdasarkan ParentId dan ItemId
func (r *transactionRepository) GetAllTransactionsRepository(page, limit int, parentId, itemId string) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	offset := (page - 1) * limit

	query := r.db.Offset(offset).Limit(limit)
	if parentId != "" {
		query = query.Where("parent_id = ?", parentId)
	}
	if itemId != "" {
		query = query.Where("item_id = ?", itemId)
	}

	result := query.Order("transaction_date DESC").Find(&transactions)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting transactions: %s", result.Error)
	}
	return transactions, nil
}

// GetTransactionByIDRepository mengambil transaksi berdasarkan ID
func (r *transactionRepository) GetTransactionByIDRepository(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	result := r.db.First(&transaction, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("transaction with ID %s not found", id)
		}
		return nil, fmt.Errorf("error getting transaction with ID %s: %s", id, result.Error)
	}
	return &transaction, nil
}

// CreateTransactionRepository membuat entri transaksi baru di database
func (r *transactionRepository) CreateTransactionRepository(transaction *model.Transaction) (*model.Transaction, error) {
	result := r.db.Create(transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return transaction, nil
}

// UpdateTransactionByIDRepository memperbarui data transaksi berdasarkan ID
func (r *transactionRepository) UpdateTransactionByIDRepository(id string, transaction *model.Transaction) (*model.Transaction, error) {
	result := r.db.Model(&model.Transaction{}).Where("id = ?", id).Updates(transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("transaction not found")
	}
	return transaction, nil
}

// DeleteTransactionByIDRepository menghapus transaksi berdasarkan ID
func (r *transactionRepository) DeleteTransactionByIDRepository(id string) error {
	result := r.db.Delete(&model.Transaction{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("transaction not found")
	}
	return nil
}

// GetTransactionsByParentIDRepository mengambil semua transaksi berdasarkan ParentID
func (r *transactionRepository) GetTransactionsByParentIDRepository(parentID string) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	result := r.db.Where("parent_id = ?", parentID).Order("transaction_date DESC").Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

// GetTransactionsByItemIDRepository mengambil semua transaksi berdasarkan ItemID
func (r *transactionRepository) GetTransactionsByItemIDRepository(itemID string) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	result := r.db.Where("item_id = ?", itemID).Order("transaction_date DESC").Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}
