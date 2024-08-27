// 数组
package array

// Filter 过滤数组
//
// 对 arr 中的每个元素执行 filterFunc, 结果为 true 的元素才会保留, 并返回过滤完的新数组
func Filter[T any](arr []T, filterFunc func(elm T) bool) []T {
	if arr == nil {
		return nil
	}
	res := make([]T, 0)
	for _, num := range arr {
		if filterFunc(num) {
			res = append(res, num)
		}
	}
	return res
}

// SafetyGet 安全地获取数组 arr 中指定索引下的元素
//
// typeElm 作用是限定元素类型, 参数值没有实际作用
func SafetyGet[T any](arr []any, idx int, typeElm T) (T, bool) {
	// 越界判断
	if idx < 0 || len(arr) < idx+1 {
		return typeElm, false
	}
	// 类型判断
	if result, ok := arr[idx].(T); ok {
		return result, ok
	}
	return typeElm, false
}
