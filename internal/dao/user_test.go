package dao_test

import (
	"fmt"
	"go_web_template/internal/dao"
	"go_web_template/internal/model/entity"
	"testing"
)

func TestUserPage(t *testing.T) {
	res, err := dao.User().Page(&entity.User{Account: "ang"}, 1, 3)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res.TotalPages, res.TotalRecords)
	for _, rec := range res.Records {
		fmt.Println(rec)
	}
}
