package model

import (
	"time"

	"github.com/zhenhua32/xingkong/third_party/deleteat"
)

type BaseModel struct {
	ID        uint               `gorm:"primarykey" json:"id"`
	CreatedAt time.Time          `json:"create_at"`
	UpdatedAt time.Time          `json:"update_at"`
	DeletedAt deleteat.DeletedAt `gorm:"index" json:"delete_at"`
}
