package product_service

import "flash_sale/models"

type Product struct {
	ID       uint
	Name     string
	Stock    int
	Price    float32
	IsActive bool
}

func (p *Product) Add() error {
	product := map[string]interface{}{
		"name":      p.Name,
		"stock":     p.Stock,
		"price":     p.Price,
		"is_active": p.IsActive,
	}

	if err := models.AddProduct(product); err != nil {
		return err
	}

	return nil
}

func (p *Product) Edit() error {
	return models.EditProduct(p.ID, map[string]interface{}{
		"name":      p.Name,
		"stock":     p.Stock,
		"price":     p.Price,
		"is_active": p.IsActive,
	})
}
