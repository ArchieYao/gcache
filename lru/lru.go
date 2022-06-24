package lru

import "container/list"

type Cache struct {
	maxBytes  int64
	nbytex    int64
	ll        *list.List                  // 双链表
	cache     map[string]*list.Element    // map 元素时双链表中对应的节点指针
	OnEvicted func(key string, val Value) // 移除元素时的回调函数
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

// New get new cache instance
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		OnEvicted: onEvicted,
		cache:     make(map[string]*list.Element),
		ll:        list.New(),
	}
}

// Get get element by key
func (c *Cache) Get(key string) (val Value, ok bool) {
	if e, ok := c.cache[key]; ok {
		c.ll.MoveToFront(e)
		kv := e.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOld
func (c *Cache) RemoveOldest() {
	e := c.ll.Back() // 取队首元素
	if e != nil {
		c.ll.Remove(e)
		kv := e.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytex -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, val Value) {
	if e, ok := c.cache[key]; ok {
		// key 存在，则更新元素
		c.ll.MoveToFront(e)
		kv := e.Value.(*entry)
		c.nbytex += int64(val.Len()) - int64(kv.value.Len())
		kv.value = val
	} else {
		// key 不存在，新增元素
		e := c.ll.PushFront(&entry{key, val})
		c.cache[key] = e
		c.nbytex += int64(len(key)) + int64(val.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytex {
		c.RemoveOldest()
	}
}

// Len
func (c *Cache) Len() int {
	return c.ll.Len()
}
