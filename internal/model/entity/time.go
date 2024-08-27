package entity

import (
	"database/sql/driver"
	"fmt"
	"go_web_template/internal/util/datetime"
	"time"
)

// Time 自定义的时间类型, 重写序列化和反序列化格式
type Time struct {
	time.Time
}

// NewTime 初始化一个自定义的时间实体
func NewTime(t time.Time) *Time {
	return &Time{t}
}

// Scan 外部设置值到当前结构体中
func (ct *Time) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*ct = Time{v}
		return nil
	default:
		return fmt.Errorf("不支持的值: %v", v)
	}
}

// Value 自定义 Time 转换为 time.Time, 使得第三方库能够转换
func (ct Time) Value() (driver.Value, error) {
	return ct.Time, nil
}

func (ct *Time) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", datetime.Format(ct.Time, datetime.TF_Default))
	return []byte(formatted), nil
}

func (ct *Time) UnmarshalJSON(b []byte) error {
	str := string(b)
	if str == "null" {
		return nil
	}
	str = str[1 : len(str)-1]

	parsedTime, err := datetime.Parse(str, datetime.TF_Default)
	if err != nil {
		return err
	}
	*ct = Time{parsedTime}
	return nil
}
