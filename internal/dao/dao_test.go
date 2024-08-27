package dao_test

import (
	"fmt"
	"go_web_template/internal/config"
	"go_web_template/internal/dao"
	"go_web_template/internal/model/entity"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local"
	config.C = &config.Config{DB: &config.DatabaseConfig{Dsn: dsn}}
	if err := dao.InitDB(); err != nil {
		panic(err)
	}
	db = dao.DB()
}

func TestPage(t *testing.T) {
	res, err := dao.Page[*entity.User](db.Model(entity.User{}), 2, 3)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("Current: ", res.Current)
	fmt.Println("Size: ", res.Size)
	fmt.Println("TotalPages: ", res.TotalPages)
	fmt.Println("TotalRecords: ", res.TotalRecords)
	fmt.Println("Records: ", res.Records)
}
