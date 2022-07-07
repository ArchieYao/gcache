package http

import (
	"fmt"
	gcache "gcache/cache"
	"net/http"

	"github.com/archieyao/groute"
)

var GcacheInst *gcache.Group
var DefaultGroupName = "default"

func GenGcacheOrLoad(name string) *gcache.Group {
	if name == "" {
		name = DefaultGroupName
	}
	if GcacheInst == nil {
		GcacheInst = gcache.NewGroup(name, 2<<10, nil)
	}
	return GcacheInst
}

func StartServer() {
	engine := groute.New()
	engine.GET("/", func(c *groute.Context) {
		c.HTML(http.StatusOK, "<h1>hello world</h1>")
	})

	gcacheGroup := engine.Group("/gcache")
	gcacheGroup.GET("/get/:key", func(c *groute.Context) {
		key := c.Param("key")
		bv, err := GcacheInst.Get(key)
		if err == nil {
			c.JSON(http.StatusOK, bv.String())
		} else {
			c.JSON(http.StatusOK, fmt.Sprintf("%s not exists", key))
		}
	})

	gcacheGroup.POST("/set", func(c *groute.Context) {
		key := c.PostForm("key")
		val := c.PostForm("value")
		bv := gcache.ByteView{B: gcache.CloneBytes([]byte(val))}
		GcacheInst.Set(key, bv)
		c.JSON(http.StatusOK, "SUCCESS")
	})

	engine.Run(":8090")

}
