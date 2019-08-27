package api

import (
	"time"
)

type Domain struct {
	ID []uint8 `db:"id"`
	Name string `db:"name"`
	AppID []uint8 `db:"app_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	AppName string `db:"app_name"`
}

