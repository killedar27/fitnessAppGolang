package model

import (
	"time"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// Order represents an order in the system
type Order struct {
	ID         int         `json:"id" db:"id"`
	UserID     int         `json:"user_id" db:"user_id"`
	ProductID  int         `json:"product_id" db:"product_id"`
	Quantity   int         `json:"quantity" db:"quantity"`
	TotalPrice float64     `json:"total_price" db:"total_price"`
	Status     OrderStatus `json:"status" db:"status"`
	CreatedAt  time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at" db:"updated_at"`
}

// CreateOrderRequest represents the request to create a new order
type CreateOrderRequest struct {
	UserID     int     `json:"user_id" binding:"required"`
	ProductID  int     `json:"product_id" binding:"required"`
	Quantity   int     `json:"quantity" binding:"required,min=1"`
	TotalPrice float64 `json:"total_price" binding:"required,min=0"`
}

// UpdateOrderStatusRequest represents the request to update order status
type UpdateOrderStatusRequest struct {
	Status OrderStatus `json:"status" binding:"required,oneof=pending confirmed shipped delivered cancelled"`
}
