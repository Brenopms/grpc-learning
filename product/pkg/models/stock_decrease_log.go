package models

type StockDecreaseLog struct {
	Id        int64 `json:"id" gorm:"primaryKey"`
	OrderId   int64 `json:"order_id"`
	ProductId int64 `json:"product_id"`
}
