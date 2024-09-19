package cache_service

import (
	"flash_sale/models"
	"flash_sale/pkg/constants"
	"strconv"
	"strings"
)

type FlashSale struct {
	ID              int  `json:"id"`
	ProductID       int  `json:"product_id"`
	DiscountPercent int  `json:"discount_percent"`
	Stock           uint `json:"stock"`
}

func (f *FlashSale) GetFlashSaleKey() string {
	return constants.CACHE_FLASH_SALE + "_" + strconv.Itoa(int(f.ID))
}

func (f *FlashSale) GetFlashSalesKey() string {
	keys := []string{
		constants.CACHE_FLASH_SALE,
	}

	if f.ID > 0 {
		keys = append(keys, strconv.Itoa(int(f.ID)))
	} else {
		keys = append(keys, "*")
	}

	return strings.Join(keys, "_")
}

func (f *FlashSale) SetCacheFlashSale(flashSale models.FlashSale) {
	f.ID = int(flashSale.ID)
	f.ProductID = int(flashSale.ProductID)
	f.DiscountPercent = flashSale.DiscountPercent
	f.Stock = flashSale.Stock
}

func (f *FlashSale) PrepareFlashSale() map[string]interface{} {
	return map[string]interface{}{
		"id":               f.ID,
		"product_id":       f.ProductID,
		"discount_percent": f.DiscountPercent,
		"stock":            f.Stock,
	}
}
