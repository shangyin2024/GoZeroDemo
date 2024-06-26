// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

const TableNameDemo = "demo"

// Demo demo
type Demo struct {
	ID          string         `gorm:"column:id;type:char(20);primaryKey" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
	Name        string         `gorm:"column:name;type:varchar(255);not null;comment:name" json:"name"`                     // name
	Sex         uint8          `gorm:"column:sex;type:tinyint unsigned;not null;default:3;comment:性别 1男 2女 3保密" json:"sex"` // 性别 1男 2女 3保密
	DateOfBirth string         `gorm:"column:date_of_birth;type:varchar(32);not null;comment:生日" json:"date_of_birth"`      // 生日
	Avatar      string         `gorm:"column:avatar;type:varchar(255);not null;comment:头像" json:"avatar"`                   // 头像
	Email       string         `gorm:"column:email;type:varchar(255);not null;comment:email" json:"email"`                  // email
	Phone       string         `gorm:"column:phone;type:varchar(32);not null;comment:手机号" json:"phone"`                     // 手机号
}

func (m *Demo) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = xid.New().String()
	return nil
}

// TableName Demo's table name
func (*Demo) TableName() string {
	return TableNameDemo
}
