// 字符串工具包
package strs

import "strings"

// Empty 判断一个字符串是否为空
//
//	strs.Empty("")       // true
//	strs.Empty("   ")    // true
//	strs.Empty(" foo  ") // false
//	strs.Empty("foo")    // false
func Empty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// AnyEmby 有任意一个字符串为空时返回 true
func AnyEmpty(strs ...string) bool {
	for _, str := range strs {
		if Empty(str) {
			return true
		}
	}
	return false
}
