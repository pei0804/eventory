package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// ページネーションクエリ作成。
func CreatePagingQuery(page int) func(db *gorm.DB) *gorm.DB {
	if page > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Offset((page - 1) * 10)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// ソートクエリ作成
func CreateSortQuery(sort string) func(db *gorm.DB) *gorm.DB {

	if sort == "created_asc" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at asc")
		}
	}
	if sort == "created_desc" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at desc")
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// 文字列検索クエリ作成。検索したいカラムは適宜設定する。
func CreateLikeQuery(q string, columns ...string) func(db *gorm.DB) *gorm.DB {
	if q != "" && len(columns) > 0 {
		return func(db *gorm.DB) *gorm.DB {
			for _, v := range columns {
				lc := fmt.Sprintf("%s LIKE ?", v)
				lq := fmt.Sprintf("%%%s%%", q)
				db = db.Where(lc, lq)
			}
			return db
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// 削除済みレコード表示クエリ作成
func CreateUnscopedQuery(isExecute bool) func(db *gorm.DB) *gorm.DB {
	if isExecute {
		return func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}
