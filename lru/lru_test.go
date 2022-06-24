package lru

import (
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestCache_Get(t *testing.T) {
	key1 := "key1"
	value1 := "val1"

	key2 := "key2"
	value2 := "val2"

	totalLen := int64(len(key1) + len(key2) + len(value1) + len(value2))

	cache := New(totalLen, func(key string, val Value) {
		t.Logf("key %s has been removed", key)
	})
	cache.Add(key1, String(value1))
	cache.Add(key2, String(value2))

	t.Log(cache.Get(key1))

	cache.Add("key3", String("val3"))

	t.Log(cache.Get(key2))
}

func Test_noReturn(t *testing.T) {
	aa, b := testNoReturn(1)
	t.Log(aa)
	t.Log(b)

	cc, d := testNoReturn(-1)
	t.Log(cc)
	t.Log(d)
}

func testNoReturn(a int) (aa int, b string) {
	if a > 0 {
		return 10, "10"
	}
	return
}
