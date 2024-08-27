// dao 层实现 负责与数据库交互
package dao

import (
	"errors"
	"fmt"
	"go_web_template/internal/config"
	"go_web_template/internal/model/entity"
	"go_web_template/internal/util/array"
	"go_web_template/internal/util/singleton"
	"go_web_template/internal/util/strs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// 用户 dao 层实现
	User = func() *UserDao { return singleton.Get(NewUserDao) }
)

// db 与数据库交互的核心对象
var db *gorm.DB

func DB() *gorm.DB {
	return db
}

// InitDB 初始化 db 对象
func InitDB() (err error) {
	dsn := config.C.DB.Dsn
	if strs.Empty(dsn) {
		return errors.New("未配置 dsn 参数")
	}
	// 打开数据库连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	// 自动初始化表
	entities := []interface{}{entity.User{}}
	for _, ett := range entities {
		if err = db.AutoMigrate(ett); err != nil {
			return
		}
	}
	return nil
}

// Page 通用的分页查询函数
//
//	查询参数需要提前设置并通过 query 对象传入
func Page[T any](query *gorm.DB, curPage, pageSize int64) (*entity.Page[T], error) {
	if query == nil {
		return nil, errors.New("query 参数必传")
	}
	res := new(entity.Page[T])

	// 计算总记录数
	if err := query.Count(&res.TotalRecords).Error; err != nil {
		return nil, err
	}

	// 设置每页大小
	if pageSize < 1 {
		res.Size = 1
	} else {
		res.Size = pageSize
	}

	// 计算总页数
	res.TotalPages = (res.TotalRecords + res.Size - 1) / res.Size

	// 设置当前页
	if curPage < 1 {
		res.Current = 1
	} else {
		res.Current = curPage
	}

	// 计算偏移量
	offset := (res.Current - 1) * res.Size

	if err := query.Offset(int(offset)).Limit(int(res.Size)).Find(&res.Records).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// InspectDbError 聚合操作数据库时的异常信息
//
//	args[0]: 可以指定期望的数据库影响行数, 如果和 result 不匹配, 返回错误
func InspectDbError(result *gorm.DB, args ...interface{}) error {
	if result == nil {
		return nil
	}
	if result.Error != nil {
		return result.Error
	}
	if ept, ok := array.SafetyGet(args, 0, int64(0)); ok && ept != result.RowsAffected {
		return fmt.Errorf("expect RowsAffected num: %d, get: %d", ept, result.RowsAffected)
	}
	return nil
}
