package structs

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           uint      `json:"-" gorm:"primary_key;autoIncrement"`
	CustomerName string    `json:"customerName" gorm:"not null;"`
	OrderedAt    time.Time `json:"orderedAt" gorm:"autoCreateTime" `
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID;references:ID"`
}

type Item struct {
	gorm.Model
	ID          uint   `json:"LineItemId" gorm:"primary_key;autoIncrement"`
	ItemCode    string `json:"itemCode" gorm:"not null;"`
	Description string `json:"description" gorm:"not null;"`
	Quantity    int    `json:"quantity" gorm:"not null;"`
	OrderID     uint   `json:"-" gorm:"not null;"`
}
