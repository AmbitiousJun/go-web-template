// 时间日期工具类
package datetime

import "time"

// TimeFormat 预定义的时间格式
type TimeFormat string

const (
	TF_Default TimeFormat = "yyyy-MM-dd HH:mm:ss"
)

// Today 根据指定的格式返回当天日期
func Today(format TimeFormat) string {
	return time.Now().Format(transformFormat(format))
}

// Format 格式化时间
func Format(t time.Time, format TimeFormat) string {
	return t.Format(transformFormat(format))
}

// Parse 转换时间
func Parse(timeStr string, format TimeFormat) (time.Time, error) {
	return time.Parse(transformFormat(format), timeStr)
}
