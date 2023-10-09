package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/xid"
)

type User struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Customer_XID []byte    `json:"customer_xid" gorm:"column:customer_xid"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type WalletManager struct {
	Customer_XID []byte    `json:"customer_xid" gorm:"column:customer_xid"`
	IsEnabled    bool      `json:"is_enabled"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type TransactionHistory struct {
	TransactionId int64     `json:"transaction_id,omitempty"`
	Customer_XID  xid.ID    `json:"customer_xid" gorm:"column:customer_xid"`
	Balance       int32     `json:"balance"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}

type Claims struct {
	Customer_XID string `json:"customer_xid" gorm:"column:customer_xid"`
	IsEnabled    bool   `json:"is_enabled"`
	jwt.RegisteredClaims
}
