package models

type StockDecreaseLog struct {
	Id        int32 `json:"id" gorm:"primaryKey"`
	OrderId   int32 `json:"order_id" gorm:"not null" validate:"required"`
	ProductId int32 `json:"product_id" gorm:"not null" validate:"required"`
	Quantity  int32 `json:"quantity" gorm:"not null" validate:"required,gte=0"`
}
