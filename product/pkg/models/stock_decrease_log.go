package models

type StockDecreaseLog struct {
	Id        int32 `json:"id" gorm:"primaryKey"`
	OrderId   int32 `json:"order_id" gorm:"not null"`
	ProductId int32 `json:"product_id" gorm:"not null"`
}
