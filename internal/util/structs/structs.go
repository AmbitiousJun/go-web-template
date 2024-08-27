// struct 相关工具函数包
package structs

import "reflect"

// Merge 合并两个结构体, 属性互补, second 优先
func Merge[T interface{}](first, second T) T {
	firstVal := reflect.ValueOf(first).Elem()
	secondVal := reflect.ValueOf(second).Elem()

	for i := 0; i < firstVal.NumField(); i++ {
		firstField := firstVal.Field(i)
		secondField := secondVal.Field(i)

		if !secondField.IsZero() {
			firstField.Set(secondField)
		}
	}
	return first
}
