package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct{
	OrderID uint64 			`json:"order_id"`
	CustomerID uuid.UUID 	`json:"customer_id"`
	LineItem []LineItem 	`json:"line_item"`
	CreateAt *time.Time		`json:"create_at"`
	ShippedAt *time.Time 	`json:"shipped_at"`
	CompletedAt *time.Time 	`json:"completed_at"`
}

type LineItem struct{
	ItemID uuid.UUID	`json:"item_id"`
	Quantity uint		`json:"quantity"`
	Price uint			`json:"price"`
}