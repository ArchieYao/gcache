package gcache

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetterFunc_Get(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})
	// 这里的f就是一个Getter的实现，实际上是 GetterFunc 实现该接口
	b, _ := f.Get("key")
	if reflect.DeepEqual(b, []byte("key")) {
		t.Log("equal")
	}
}

var (
	db = map[string]string{
		"zhangsan": "23",
		"lisi":     "234",
		"wangwu":   "34",
	}
)

func TestGroup_Get(t *testing.T) {
	gcache := NewGroup("test1", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			if val, ok := db[key]; ok {
				t.Logf("%s get from db", key)
				return []byte(val), nil
			}
			return nil, fmt.Errorf("cannot get %s", key)
		}))

	for k, v := range db {
		bv, err := gcache.GetAndLoad(k)
		if err != nil {
			t.Logf("get key %s with error %v", k, err)
		}
		t.Logf("key %s val %s, real v %s", k, bv.String(), v)

	}

	for k, v := range db {
		bv, err := gcache.Get(k)
		if err != nil {
			t.Logf("get key %s with error %v", k, err)
		}
		t.Logf("key %s val %s, real v %s", k, bv.String(), v)

	}
}
