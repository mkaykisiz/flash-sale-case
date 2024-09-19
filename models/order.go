package models

import (
	"errors"
	"flash_sale/pkg/logging"
	"gorm.io/gorm"
)

type Order struct {
	Model

	UserID      uint        `json:"user_id"`
	User        User        `json:"user" gorm:"foreignKey:UserID"`
	FlashSaleID uint        `json:"flash_sale_id"`
	FlashSale   FlashSale   `json:"flash_sale" gorm:"foreignKey:FlashSaleID"`
	Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	NetPrice    float32     `json:"net_price"`
	TotalPrice  float32     `json:"total_price"`
}

type OrderItem struct {
	Model

	OrderID         uint    `json:"order_id"`
	Order           Order   `json:"order" gorm:"foreignKey:OrderID"`
	ProductID       uint    `json:"product_id"`
	Product         Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity        uint    `json:"quantity"`
	Price           float32 `json:"price"`
	DiscountedPrice float32 `json:"discounted_price"`
}

// CreateOrder create order and order items
func CreateOrder(db *gorm.DB, order Order) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// BuyFlashSale buy new flash sale
func BuyFlashSale(order Order) error {
	// Transaction
	tx := db.Begin()
	if tx.Error != nil {
		logging.Error("Transaction Error: ", tx.Error)
	}

	var flashSale FlashSale
	err := tx.Preload("Product").Where("id = ? AND deleted_on = ? ", order.FlashSaleID, 0).First(&flashSale).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	var totalQuantity uint = 0
	for _, orderItem := range order.Items {
		totalQuantity += orderItem.Quantity
	}
	flashSale.Stock -= totalQuantity
	data := map[string]interface{}{
		"stock": flashSale.Stock,
	}
	if err := tx.Model(&FlashSale{}).Where("id = ?", flashSale.ID).Updates(data).Error; err != nil {
		return err
	}
	err = CreateOrder(tx, order)
	if err != nil {
		return err
	}
	var product Product
	err = tx.Where("id = ? AND deleted_on = ? ", flashSale.ProductID, 0).First(&product).Error
	if err != nil {
		return err
	}
	data = map[string]interface{}{
		"stock": product.Stock - int(totalQuantity),
	}
	err = UpdateStock(tx, flashSale.ProductID, data)
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		logging.Error("Transaction commit error: ", err)
		return err
	}
	return nil
}
