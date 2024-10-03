package repository

import (
	"Edupay/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// ItemRepository adalah interface untuk operasi CRUD pada entitas Item
type ItemRepository interface {
	GetAllItemsRepository(page, limit int, name string) ([]*model.Item, error)
	GetItemByIDRepository(id string) (*model.Item, error)
	CreateItemRepository(item *model.Item) (*model.Item, error)
	UpdateItemByIDRepository(id string, item *model.Item) (*model.Item, error)
	DeleteItemByIDRepository(id string) error
}

// itemRepository adalah struct yang mengimplementasikan ItemRepository
type itemRepository struct {
	db *gorm.DB
}

// NewItemRepository membuat instance baru dari itemRepository
func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db}
}

// GetAllItemsRepository mengambil semua item dengan pagination dan pencarian berdasarkan nama
func (r *itemRepository) GetAllItemsRepository(page, limit int, name string) ([]*model.Item, error) {
	var items []*model.Item
	offset := (page - 1) * limit

	query := r.db.Offset(offset).Limit(limit)
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	result := query.Order("created_at DESC").Find(&items)
	if result.Error != nil {
		return nil, fmt.Errorf("error getting items: %s", result.Error)
	}
	return items, nil
}

// GetItemByIDRepository mengambil item berdasarkan ID
func (r *itemRepository) GetItemByIDRepository(id string) (*model.Item, error) {
	var item model.Item
	result := r.db.First(&item, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("item with ID %s not found", id)
		}
		return nil, fmt.Errorf("error getting item with ID %s: %s", id, result.Error)
	}
	return &item, nil
}

// CreateItemRepository membuat item baru
func (r *itemRepository) CreateItemRepository(item *model.Item) (*model.Item, error) {
	result := r.db.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

// UpdateItemByIDRepository memperbarui data item berdasarkan ID
func (r *itemRepository) UpdateItemByIDRepository(id string, item *model.Item) (*model.Item, error) {
	result := r.db.Model(&model.Item{}).Where("id = ?", id).Updates(item)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("item not found")
	}
	return item, nil
}

// DeleteItemByIDRepository menghapus item berdasarkan ID
func (r *itemRepository) DeleteItemByIDRepository(id string) error {
	result := r.db.Delete(&model.Item{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("item not found")
	}
	return nil
}
