package singleton

import (
	"reflect"
	"sync"
)

// instanceHolder 存放已创建的实例
type instanceHolder struct {
	instance interface{}
}

// holderMap 存放每个 id 值的实例对象
var holderMap = map[string]*instanceHolder{}

// mutex 进行并发控制
var mutex = sync.Mutex{}

// Get 获取单例实例
func Get[T any](createFunc func() T) T {
	key := reflect.TypeOf((*T)(nil)).Elem().String()
	if holder, ok := holderMap[key]; ok {
		return holder.instance.(T)
	}
	mutex.Lock()
	defer mutex.Unlock()
	holder := new(instanceHolder)
	holder.instance = createFunc()
	holderMap[key] = holder
	return holder.instance.(T)
}
