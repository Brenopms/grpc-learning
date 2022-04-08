package models

type Product struct {
	Id                int32            `json:"id" gorm:"primaryKey" validate:"required"`
	Name              string           `json:"name" validate:"required"`
	Stock             int32            `json:"stock" validate:"required,gte=0"`
	Price             float64          `json:"price" validate:"required,gte=0"`
	Sku               string           `json:"sku" validate:"required"`
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductId"`
}
