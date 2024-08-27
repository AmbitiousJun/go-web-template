// 通用实体类
package entity

import "gorm.io/gorm"

// E 封装数据库表通用字段
type E struct {
	Id         uint           `json:"id" gorm:"type:BIGINT;primaryKey;not null;autoIncrement;comment:主键"`
	CreateTime *Time          `json:"createTime" gorm:"type:DATETIME;autoCreateTime;comment:创建时间"`
	UpdateTime *Time          `json:"updateTime" gorm:"type:DATETIME;autoUpdateTime;comment:更新时间"`
	DeleteTime gorm.DeletedAt `json:"deleteTime" gorm:"type:DATETIME;comment:删除时间"`
}

// Page 数据库分页数据封装
type Page[T any] struct {
	Current      int64 `json:"current"`
	Size         int64 `json:"size"`
	TotalPages   int64 `json:"totalPages"`
	TotalRecords int64 `json:"totalRecords"`
	Records      []T   `json:"records"`
}
