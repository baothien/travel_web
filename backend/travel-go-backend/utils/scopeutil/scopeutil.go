package scopeutil

import (
	"gorm.io/gorm"
)

func UserTypeFilter(userType string) func(db *gorm.DB) *gorm.DB {
	if userType != "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("user_type = ?", userType)
		}
	}
	return nil
}
