package maps

import (
	"fmt"
	"sort"
	"testing"
)

func TestSetMaps(t *testing.T) {
	maps := NewMaps()
	a := maps.Set("key1", "123123")
	fmt.Println(a)
}

func TestExistMaps(t *testing.T) {
	maps := NewMaps()

	a := maps.Set("key2", "this is value 2")

	fmt.Println(a.Exist("key3"))
}

func TestGetMaps(t *testing.T) {
	maps := NewMaps()

	a := maps.Set("key1", "value1").Set("key2", "value 2").Set("key3", "value 3")

	fmt.Println(a.Get("key4"))
}

func TestDeleteMaps(t *testing.T) {
	maps := NewMaps()

	a := maps.Set("key1", "value1").Set("key2", "value 2").Set("key3", "value 3")

	fmt.Println(a.Delete("key1", "key2", "key3", "key4"))
}

func TestKeysValues(t *testing.T) {
	maps := NewMaps()

	a := maps.Set("key1", "value1").Set("key2", "value 2").Set("key3", "value 3")

	fmt.Println(a.Keys())
	fmt.Println(a.Values())
}

func TestOriginMap(t *testing.T) {
	maps := NewMaps()

	a := maps.Set("key1", "value1").Set("key2", "value 2").Set("key3", "value 3")

	fmt.Println(a.OriginMap())
}

func TestSortKey(t *testing.T) {
	maps := NewMaps()

	a := maps.Set("key3", "value1").Set("key4", "value 2").Set("key1", "value 3")

	a.SortKey(sort.Strings)

	fmt.Println(a)
}
