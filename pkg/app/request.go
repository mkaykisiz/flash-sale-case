package app

import (
	"flash_sale/pkg/logging"
	"github.com/astaxie/beego/validation"
	"time"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}

type Auth struct {
	Username string `json:"username" valid:"Required; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MaxSize(50)"`
}
type AddProductData struct {
	Name     string  `json:"name" valid:"Required;MaxSize(100)"`
	Stock    int     `json:"stock" valid:"Required;Min(1)"`
	Price    float32 `json:"price" valid:"Required"`
	IsActive bool    `json:"is_active" valid:"Required"`
}

type EditProductData struct {
	Name     string `json:"name" valid:"Required;MaxSize(100)"`
	Stock    int    `json:"stock" valid:"Min(0)"`
	IsActive bool   `json:"is_active" valid:"Required"`
}

type AddFlashSaleRequest struct {
	ProductID       uint      `json:"product_id"`
	DiscountPercent int       `json:"discount_percent"`
	Stock           int       `json:"stock"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
}

type AddFlashSaleResponse struct {
	ID uint `json:"id"`
}
type FlashSale struct {
	ID              uint `json:"id"`
	ProductID       uint `json:"product_id"`
	DiscountPercent int  `json:"discount_percent"`
	Stock           int  `json:"stock"`
}

type GetFlashSalesResponse struct {
	FlashSales []FlashSale `json:"flash_sales"`
	Count      int         `json:"count"`
}

type EditFlashSaleRequest struct {
	DiscountPercent int       `json:"discount_percent"`
	Stock           int       `json:"stock"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
}

type BuyFlashSaleRequest struct {
	Unit uint `json:"unit"`
}

type GetFlashSaleResponse struct {
	FlashSale
}
