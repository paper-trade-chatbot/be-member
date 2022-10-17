package specification

import (
	"github.com/jinzhu/gorm"
)

// NewPaginationSpecification generate pagination query
func NewPaginationSpecification(pagination *PaginationStruct) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(pagination.Index).Limit(pagination.PageSize)
	}
}
