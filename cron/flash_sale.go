package cron

import (
	"flash_sale/models"
	"flash_sale/pkg/logging"
	"flash_sale/pkg/redis"
	"flash_sale/service/cache_service"
	"time"
)

// SyncFlashSale update Flash Sale data
func SyncFlashSale() {
	now := time.Now()
	var query = make(map[string]interface{})
	query["start_time <= ?"] = now
	query["end_time >= ?"] = now
	query["stock >= ?"] = 0
	query["deleted_on = ?"] = 0

	flashSales, err := models.GetFlashSales(query)
	if err != nil {
		logging.Error(err)
		return
	}
	for _, flashSale := range flashSales {
		cache := cache_service.FlashSale{}
		cache.SetCacheFlashSale(*flashSale)
		key := cache.GetFlashSaleKey()
		_, err = redis.Delete(key)
		if err != nil {
			logging.Error(err)
		}
		data := cache.PrepareFlashSale()
		err = redis.HMSet(key, data)
		if err != nil {
			logging.Error(err)
			return
		}
	}
	logging.Info("[SyncFlashSale] updated Flash Sales count", len(flashSales))
}
