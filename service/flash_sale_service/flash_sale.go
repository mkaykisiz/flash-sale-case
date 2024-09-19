package flash_sale_service

import (
	"encoding/json"
	"flash_sale/models"
	"flash_sale/pkg/logging"
	"flash_sale/pkg/redis"
	"flash_sale/service/cache_service"
	"time"
)

type FlashSale struct {
	ID              uint      `json:"id"`
	ProductID       uint      `json:"product_id"`
	DiscountPercent int       `json:"discount_percent"`
	Stock           int       `json:"stock"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	Unit            uint      `json:"unit"`
	UserID          uint      `json:"user_id"`
}

func (f *FlashSale) Add() (error, uint) {
	flashSaleData := map[string]interface{}{
		"product_id":       f.ProductID,
		"discount_percent": f.DiscountPercent,
		"stock":            f.Stock,
		"start_time":       f.StartTime,
		"end_time":         f.EndTime,
	}
	err, flashSale := models.AddFlashSale(flashSaleData)
	if err != nil {
		return err, 0
	}

	// update cache
	cache := cache_service.FlashSale{ID: int(flashSale.ID)}
	key := cache.GetFlashSaleKey()
	cache.SetCacheFlashSale(*flashSale)
	data := cache.PrepareFlashSale()
	err = redis.HMSet(key, data)
	if err != nil {
		logging.Error(err)
		return nil, flashSale.ID
	}
	return nil, flashSale.ID
}

func (f *FlashSale) Edit() error {
	err := models.EditFlashSale(f.ID, map[string]interface{}{
		"stock":            f.Stock,
		"discount_percent": f.DiscountPercent,
		"start_time":       f.StartTime,
		"end_time":         f.EndTime,
	})
	if err != nil {
		return err
	}

	// update cache
	flashSale, err := models.GetFlashSale(int(f.ID))
	if err != nil {
		return err
	}
	cache := cache_service.FlashSale{}
	cache.SetCacheFlashSale(*flashSale)
	key := cache.GetFlashSalesKey()
	data := cache.PrepareFlashSale()
	err = redis.HMSet(key, data)
	if err != nil {
		logging.Error(err)
		return err
	}

	return nil
}

func (f *FlashSale) Get() (*models.FlashSale, error) {
	var cacheFlashSale *models.FlashSale

	// update cache
	cache := cache_service.FlashSale{ID: int(f.ID)}
	key := cache.GetFlashSaleKey()
	if redis.Exists(key) {
		data, err := redis.Get(key)
		if err != nil {
			logging.Info(err)
			return nil, err
		} else {
			jsonData, _ := json.Marshal(data)
			err = json.Unmarshal(jsonData, &cacheFlashSale)
			if err != nil {
				logging.Info(err)
			}
		}
	}

	return cacheFlashSale, nil
}

func (f *FlashSale) GetAll() ([]*models.FlashSale, error) {
	var flashSales []*models.FlashSale

	// update cache
	cache := cache_service.FlashSale{}
	key := cache.GetFlashSalesKey()
	data, err := redis.HGetAll(key)
	if err != nil {
		logging.Info(err)
	} else {
		jsonData, _ := json.Marshal(data)
		err = json.Unmarshal(jsonData, &flashSales)
		if err != nil {
			logging.Info(err)
		}
	}
	return flashSales, nil
}

func (f *FlashSale) Delete() error {
	// update cache
	cache := cache_service.FlashSale{ID: int(f.ID)}
	key := cache.GetFlashSaleKey()
	if redis.Exists(key) {
		_, err := redis.Delete(key)
		if err != nil {
			logging.Info(err)
		}
	}

	return models.DeleteFlashSale(f.ID)
}

func (f *FlashSale) Buy() error {
	flashSale, err := models.GetFlashSale(int(f.ID))
	if err != nil {
		logging.Info(err)
		return err
	}
	flashSale.Stock -= f.Unit
	discountAmount := flashSale.Product.Price * float32(flashSale.DiscountPercent) / 100
	discountedPrice := flashSale.Product.Price - discountAmount
	order := models.Order{
		UserID:      f.UserID,
		FlashSaleID: flashSale.ID,
		NetPrice:    discountedPrice * float32(f.Unit),
		TotalPrice:  flashSale.Product.Price * float32(f.Unit),
		Items: []models.OrderItem{
			{
				ProductID:       flashSale.ProductID,
				Quantity:        f.Unit,
				Price:           flashSale.Product.Price,
				DiscountedPrice: discountedPrice,
			},
		},
	}

	err = models.BuyFlashSale(order)
	if err != nil {
		logging.Info(err)
		return err
	}

	// update cache
	cache := cache_service.FlashSale{}
	cache.SetCacheFlashSale(*flashSale)
	key := cache.GetFlashSalesKey()
	data := cache.PrepareFlashSale()
	err = redis.HMSet(key, data)
	if err != nil {
		logging.Error(err)
		return err
	}
	return nil
}

func (f *FlashSale) ExistByID() (bool, error) {
	// update cache
	cache := cache_service.FlashSale{ID: int(f.ID)}
	key := cache.GetFlashSaleKey()
	if redis.Exists(key) {
		_, err := redis.Delete(key)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
