package gcache

import "sync"

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

// Get 实现Getter接口
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var (
	mutex  sync.RWMutex
	groups = make(map[string]*Group)
)

// NewGroup 这里的getter是真实获取val的方法，初始化时传入
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mutex.Lock()
	defer mutex.Unlock()
	group := &Group{
		name:      name,
		mainCache: cache{cacheBytes: cacheBytes},
		getter:    getter,
	}
	groups[name] = group
	return group
}

func GetGroup(name string) *Group {
	mutex.RLock()
	group := groups[name]
	mutex.RUnlock()
	return group
}

// Get 从cache中获取
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, nil
	}
	if bv, ok := g.mainCache.get(key); ok {
		return bv, nil
	}
	return ByteView{}, nil
}

// GetAndLoad 优先从缓存中获取，如果缓存不存在，从getter实现中加载
func (g *Group) GetAndLoad(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, nil
	}
	if bv, ok := g.mainCache.get(key); ok {
		return bv, nil
	}
	if g.getter == nil {
		return ByteView{}, nil
	}
	return g.load(key)
}

func (g *Group) load(key string) (ByteView, error) {
	// 调用getter，获取val
	bv, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	val := ByteView{b: cloneBytes(bv)}
	g.populateCache(key, val)
	return val, nil
}

func (g *Group) populateCache(key string, val ByteView) {
	g.mainCache.add(key, val)
}
