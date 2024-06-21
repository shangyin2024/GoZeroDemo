package common

import (
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type Method struct {
	ID string
}

func (m *Method) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = xid.New().String()
	return nil
}
