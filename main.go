package main

import (
	"container/list"
	"fmt"
)

var CacheConstructorMap = map[string]func(levelInfo *LevelInfo) ICache{
	"LRU": LRUConstructor,
}

type ICache interface {
	read(key string) string
	write(key string, value string)
	getSize() int
	getTotalCapacity() int
	getWriteTime() int
	getReadTime() int
}

type LRUCache struct {
	capacity  int
	readTime  int
	writeTime int
	nodeMap   map[string]*list.Element
	list      *list.List
}

type Pair struct {
	key   string
	value string
}

func (c *LRUCache) getSize() int {
	return c.list.Len()
}

func (c *LRUCache) getTotalCapacity() int {
	return c.capacity
}

func (c *LRUCache) getReadTime() int {
	return c.readTime
}

func (c *LRUCache) getWriteTime() int {
	return c.writeTime
}

func LRUConstructor(info *LevelInfo) ICache {
	return &LRUCache{
		capacity:  info.capacity,
		readTime:  info.readTime,
		writeTime: info.writeTime,
		list:      new(list.List),
		nodeMap:   make(map[string]*list.Element, info.capacity),
	}
}

func (c *LRUCache) read(key string) string {
	if node, ok := c.nodeMap[key]; ok {
		val := node.Value.(*list.Element).Value.(Pair).value
		c.list.MoveToFront(node)
		return val
	}
	return ""
}

func (c *LRUCache) write(key string, value string) {
	if node, ok := c.nodeMap[key]; ok {
		c.list.MoveToFront(node)
		node.Value.(*list.Element).Value = Pair{
			key:   key,
			value: value,
		}
	} else {
		if c.list.Len() == c.capacity {
			idx := c.list.Back().Value.(*list.Element).Value.(Pair).key
			delete(c.nodeMap, idx)
			c.list.Remove(c.list.Back())
		}
		node := &list.Element{
			Value: Pair{
				key:   key,
				value: value,
			},
		}
		ptr := c.list.PushFront(node)
		c.nodeMap[key] = ptr
	}
}

type LevelInfo struct {
	capacity  int
	readTime  int
	writeTime int
	cacheType string
}

type CacheManager struct {
	lastNCount     int
	previousReads  []int
	previousWrites []int
	caches         []ICache
}

func NewCacheManager(levelInfos []*LevelInfo, lastNCount int) *CacheManager {
	var caches []ICache
	for _, levelInfo := range levelInfos {
		cacheConstructor := CacheConstructorMap[levelInfo.cacheType]
		cache := cacheConstructor(levelInfo)
		caches = append(caches, cache)
	}
	return &CacheManager{caches: caches, lastNCount: lastNCount}
}

func (c *CacheManager) read(key string) string {
	total := 0
	val := ""
	for _, cache := range c.caches {
		val = cache.read(key)
		total = total + cache.getReadTime() //new
		if val != "" {
			break
		}
	}
	c.levelRead(total)
	return val
}

func (c *CacheManager) write(key string, value string) {
	total := 0 //new
	for _, cache := range c.caches {
		res := cache.read(key)
		total = total + cache.getReadTime() //new 		total = total + c.levelInfos[index].readTime
		if res == value {
			continue
		}
		cache.write(key, value)
		total = total + cache.getWriteTime()
	}
	c.levelWrite(total) //new
}

func (c *CacheManager) levelRead(level int) {
	if len(c.previousReads) >= c.lastNCount {
		c.previousReads = c.previousReads[1:]
	}
	c.previousReads = append(c.previousReads, level)
}

func (c *CacheManager) levelWrite(level int) {
	if len(c.previousWrites) >= c.lastNCount {
		c.previousWrites = c.previousWrites[1:]
	}
	c.previousWrites = append(c.previousWrites, level)
}

type Statistics struct {
	levelStats    map[int]float64
	lastNReadAvg  float64
	lastNWriteAvg float64
}

func (c *CacheManager) getStats() *Statistics {
	stats := &Statistics{
		levelStats:    map[int]float64{},
		lastNReadAvg:  0,
		lastNWriteAvg: 0,
	}
	for level, levelCache := range c.caches {
		stats.levelStats[level] = (float64(levelCache.getSize()) / float64(levelCache.getTotalCapacity())) * 100
	}
	totalReadTime := 0
	totalWriteTime := 0
	for _, readTime := range c.previousReads {
		totalReadTime = totalReadTime + readTime
	}
	for _, writeTime := range c.previousWrites {
		totalWriteTime = totalWriteTime + writeTime
	}
	if len(c.previousReads) > 0 {
		stats.lastNReadAvg = float64(totalReadTime) / float64(len(c.previousReads))
	}
	if len(c.previousWrites) > 0 {
		stats.lastNWriteAvg = float64(totalWriteTime) / float64(len(c.previousWrites))
	}
	return stats
}

func main() {
	levelInfos := []*LevelInfo{
		{
			capacity:  1,
			readTime:  1,
			writeTime: 2,
			cacheType: "LRU",
		}, {
			capacity:  20,
			readTime:  3,
			writeTime: 4,
			cacheType: "LRU",
		}, {
			capacity:  30,
			readTime:  5,
			writeTime: 6,
			cacheType: "LRU",
		},
	}
	manager := NewCacheManager(levelInfos, 10)
	fmt.Println(manager.read("a"))
	fmt.Println(manager.getStats())
	manager.write("a", "abc")
	fmt.Println(manager.getStats())
	manager.write("a", "abc")
	fmt.Println(manager.getStats())
	manager.write("x", "xyz")
	fmt.Println(manager.getStats())
	manager.write("a", "abx")
	fmt.Println(manager.getStats())
	manager.write("x", "xyz")
	fmt.Println(manager.read("a"))
	fmt.Println(manager.read("x"))
	fmt.Println(manager.getStats())
}
