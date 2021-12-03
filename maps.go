package maps

import (
	"encoding/json"
	"sync"
)

const (
	// SortAsc sort by asc
	SortAsc = "asc"
	// SortDesc sort by desc
	SortDesc = "desc"
)

type maps struct {
	rw    sync.RWMutex
	Key   []string
	Value []interface{}
	Map   map[string]interface{}
}

func NewMaps() Maps {
	return &maps{
		Key:   []string{},
		Value: []interface{}{},
		Map:   map[string]interface{}{},
	}
}

type Maps interface {
	// Set 设置键值对
	Set(key string, value interface{}) Maps
	// Exist 指定的key是否存在
	Exist(Key string) bool
	// Get 获取指定Key的值
	Get(key string) interface{}
	// Delete 删除指定的某一些key 返回受影响的数目
	Delete(keys ...string) int
	// Keys 获取这个map的全部key
	Keys() []string
	// Values 获取这个map的全部值
	Values() []interface{}
	// OriginMap 获取这个map的原始map
	OriginMap() map[string]interface{}
	// SortKey 根据Key排序
	SortKey(sortFunc func(keys []string))
	// String 数据转为String
	String() string
}

//	@method Set
//	@description: set
//	@receiver m
//	@param key string
//	@param value interface{}
//	@return Maps
func (m *maps) Set(key string, value interface{}) Maps {
	defer m.rw.Unlock()
	m.rw.Lock()

	m.Key = append(m.Key, key)
	m.Value = append(m.Value, value)

	m.Map[key] = value

	return m
}

//	@method Exist
//	@description: 判断某个key是否存在
//	@receiver m
//	@param Key string
//	@return bool
func (m *maps) Exist(Key string) bool {
	defer m.rw.RUnlock()

	m.rw.RLock()

	for _, v := range m.Key {
		if Key == v {
			return true
		}
	}

	return false
}

//	@method Get
//	@description: 获取指定key的数据
//	@receiver m
//	@param key string
//	@return interface{}
func (m *maps) Get(key string) interface{} {
	defer m.rw.RUnlock()

	m.rw.RLock()

	if !m.Exist(key) {
		return nil
	}

	return m.Map[key]
}

//	@method Delete
//	@description: 删除指定key
//	@receiver m
//	@param keys ...string
//	@return int 删除成功的数量
func (m *maps) Delete(keys ...string) int {
	defer m.rw.Unlock()

	m.rw.Lock()

	if len(keys) <= 0 {
		return 0
	}

	mCount := len(m.Value)

	for _, v := range keys {
		delete(m.Map, v)
	}

	return mCount - len(m.Value)
}

//	@method Keys
//	@description: 获取全部的key
//	@receiver m
//	@return []string
func (m *maps) Keys() []string {
	return m.Key
}

//	@method Values
//	@description: 获取全部的value
//	@receiver m
//	@return []interface{}
func (m *maps) Values() []interface{} {
	return m.Value
}

//	@method OriginMap
//	@description: 获取原始map
//	@receiver m
//	@return map[string]interface{}
func (m *maps) OriginMap() map[string]interface{} {
	return m.Map
}

//	@method SortKey
//	@description: 根据Key排序
//	@receiver m
//	@param sortFunc func(keys []string)
func (m *maps) SortKey(sortFunc func(keys []string)) {
	sortFunc(m.Key)
}

//	@method String
//	@description: 数据string
//	@receiver m
//	@return string
func (m *maps) String() string {
	result, err := json.Marshal(m)
	if err != nil {
		return ""
	}

	return string(result)
}
