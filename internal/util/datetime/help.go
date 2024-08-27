// 时间工具类辅助函数包, 不对外暴露
package datetime

import "strings"

// fmtStrMap 将常规的日期格式化字符串映射成 Go 语言的日期格式化字符串
var fmtStrMap = map[string]string{
	"yyyy": "2006",
	"MM":   "01",
	"dd":   "02",
	"HH":   "15",
	"mm":   "04",
	"ss":   "05",
}

// transformFormat 将常规的日期格式化字符串转换成 Go 内部能够解析的格式
func transformFormat(format TimeFormat) string {
	var str string = string(format)
	for k, v := range fmtStrMap {
		str = strings.ReplaceAll(str, k, v)
	}
	return str
}
