package models

type Order struct {
	Id        int32   `json:"id" gorm:"primaryKey"`
	Price     float64 `json:"price" gorm:"not null" validate:"required,gte=0"`
	ProductId int32   `json:"product_id" gorm:"not null" validate:"required"`
	UserId    string  `json:"user_id" gorm:"not null" validate:"required"`
}
