package models

import (
	"errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

type FlashSale struct {
	Model

	ProductID       uint      `gorm:"not null" json:"product_id"`
	Product         Product   `json:"product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	DiscountPercent int       `gorm:"not null;check:discount_percent >= 0 AND discount_percent <= 100" json:"discount_percent"`
	Stock           uint      `gorm:"not null;check:stock >= 0" json:"stock"`
	StartTime       time.Time `gorm:"not null" json:"start_time"`
	EndTime         time.Time `gorm:"not null" json:"end_time"`
}

// GetFlashSales gets a list of flash sales based on paging constraints
func GetFlashSales(queryMap map[string]interface{}) ([]*FlashSale, error) {
	var flashSale []*FlashSale
	var conditions []string
	var values []interface{}
	for condition, value := range queryMap {
		conditions = append(conditions, condition)
		values = append(values, value)
	}
	queryStr := strings.Join(conditions, " AND ")

	err := db.Preload("Product").Where(queryStr, values...).Find(&flashSale).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return flashSale, nil
}

// GetFlashSale Get a single flashSale based on ID
func GetFlashSale(id int) (*FlashSale, error) {
	var flashSale FlashSale
	err := db.Preload("Product").Where("id = ? AND deleted_on = ? ", id, 0).First(&flashSale).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(&flashSale).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &flashSale, nil
}

// EditFlashSale modify a single flashSale
func EditFlashSale(id uint, data interface{}) error {
	if err := db.Model(&FlashSale{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddFlashSale add a single flashSale
func AddFlashSale(data map[string]interface{}) (error error, flashSale *FlashSale) {
	flashSale = &FlashSale{
		ProductID:       data["product_id"].(uint),
		Stock:           data["stock"].(uint),
		DiscountPercent: data["discount_percent"].(int),
		StartTime:       data["start_time"].(time.Time),
		EndTime:         data["end_time"].(time.Time),
	}
	if err := db.Create(&flashSale).Error; err != nil {
		return err, nil
	}

	return nil, flashSale
}

// DeleteFlashSale delete a single flashSale
func DeleteFlashSale(id uint) error {
	var data = make(map[string]interface{})
	data["deleted_on"] = time.Now().Unix()

	if err := db.Model(&FlashSale{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
