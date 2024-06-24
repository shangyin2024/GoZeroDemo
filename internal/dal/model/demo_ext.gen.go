// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"encoding/json"
	"time"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

const TableNameDemoExt = "demo_ext"

// DemoExt demo ext
type DemoExt struct {
	ID        string           `gorm:"column:id;type:char(20);primaryKey" json:"id"`
	CreatedAt time.Time        `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt time.Time        `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
	DemoID    string           `gorm:"column:demo_id;type:char(20);not null" json:"demo_id"`
	Extra     *json.RawMessage `gorm:"column:extra;type:json;comment:扩展字段" json:"extra"` // 扩展字段
}

func (m *DemoExt) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = xid.New().String()
	return nil
}

// TableName DemoExt's table name
func (*DemoExt) TableName() string {
	return TableNameDemoExt
}