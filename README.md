# maps

###maps
可根据map的key按照有序排序的map包

###使用:
- go get -u github.com/cowardmrx/maps
```go
// Set
maps := NewMaps()
maps.Set("key1", "123123")

// Exist
maps.Exist("key1")

// Get
maps.Get("key")

// Delete
maps.Delete("key1","key2")

// GetMapsKeys return a []string
maps.Keys() 

// GetMapsValues return a []interface
maps.Values()

// GetMapsOriginMap get this maps origin data return a map[string]interface
maps.OriginMap()

// SortKey sort this maps by keys  asc
maps.SortKey()

// other sort
// sort by desc
sort.Slice(maps.Keys,func (i,j int) bool {
    return maps.Keys()[i] > maps.Keys()[j]
})

// sort by asc
sort.Slice(maps.Keys,func (i,j int) bool {
    return maps.Keys()[i] < maps.Keys()[j]
})


// other example ./maps_test.go
```