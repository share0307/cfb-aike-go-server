package base

import "github.com/jinzhu/gorm"

/**
	模型基类，给予GORM
 */
type BaseModel struct {
	// 使用gorm
	gorm.Model
}


