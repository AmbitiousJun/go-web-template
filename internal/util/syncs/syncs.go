// 同步工具包
package syncs

import (
	"go_web_template/internal/util/strs"
	"sync"
)

const (
	KeyUserRegister = "user:register" // 用户注册并发控制
)

// globalMutex 包内部进行并发控制
var globalMutex = sync.Mutex{}

// mutexMap 集中存放 mutex
var mutexMap = map[string]*sync.Mutex{}

// Mutex 根据指定 key 获取一个 mutex 对象
func Mutex(key string) *sync.Mutex {
	if strs.Empty(key) {
		return nil
	}

	if mu, ok := mutexMap[key]; ok {
		return mu
	}
	globalMutex.Lock()
	defer globalMutex.Unlock()
	if mu, ok := mutexMap[key]; ok {
		return mu
	}

	mu := new(sync.Mutex)
	mutexMap[key] = mu
	return mu
}
