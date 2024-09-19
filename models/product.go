package models

import (
	"errors"
	"gorm.io/gorm"
)

type Product struct {
	Model

	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float32 `json:"price"`
}

// GetProduct Get a single product based on ID
func GetProduct(id int) (*Product, error) {
	var product Product
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&product).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = db.Model(&product).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &product, nil
}

// EditProduct modify a single product
func EditProduct(id uint, data interface{}) error {
	if err := db.Model(&Product{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddProduct add a single product
func AddProduct(data map[string]interface{}) error {
	product := Product{
		Name:  data["name"].(string),
		Stock: data["stock"].(int),
		Price: data["price"].(float32),
	}
	if err := db.Create(&product).Error; err != nil {
		return err
	}

	return nil
}

// UpdateStock modify a single product stock
func UpdateStock(db *gorm.DB, id uint, data interface{}) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := db.Model(&Product{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// DeleteProduct delete a single PRODUCT
func DeleteProduct(id int) error {
	if err := db.Where("id = ?", id).Delete(Product{}).Error; err != nil {
		return err
	}

	return nil
}
