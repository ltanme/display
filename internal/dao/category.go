// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"display/internal/dao/internal"
)

// categoryDao is the data access object for table gf_category.
// You can define custom methods on it to extend its functionality as you wish.
type categoryDao struct {
	*internal.CategoryDao
}

var (
	// Category is globally public accessible object for table gf_category operations.
	Category = categoryDao{
		internal.NewCategoryDao(),
	}
)

// Fill with you ideas below.
