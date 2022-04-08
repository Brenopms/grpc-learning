package models

type Product struct {
	Id                int32            `json:"id" gorm:"primaryKey"`
	Name              string           `json:"name"`
	Stock             int32            `json:"stock"`
	Price             float64          `json:"price"`
	Sku               string           `json:"sku"`
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductId"`
}
