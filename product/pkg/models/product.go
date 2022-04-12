package models

type Product struct {
	Id                int32            `json:"id" gorm:"primaryKey"`
	Name              string           `json:"name" gorm:"not null" validate:"required"`
	Stock             int32            `json:"stock" gorm:"not null" validate:"required,gte=0"`
	Price             float64          `json:"price" gorm:"not null" validate:"required,gte=0"`
	Sku               string           `json:"sku" gorm:"not null" validate:"required"`
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductId" validate:"-"`
}
