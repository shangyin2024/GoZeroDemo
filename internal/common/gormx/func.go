package gormx

import (
	"errors"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// CheckError 检查错误
func CheckError(e error) (exists bool, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if e != nil {
		return false, err
	}
	return true, nil
}

// Paginate 数据分页
func Paginate(page, pageSize int64) func(gen.Dao) gen.Dao {

	return func(dao gen.Dao) gen.Dao {
		if page == 0 {
			page = 1
		}
		if pageSize > 100 {
			pageSize = 100
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return dao.Limit(int(pageSize)).Offset(int(offset))
	}
}

func PaginateWithRaw(page, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
