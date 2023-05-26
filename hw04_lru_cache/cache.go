package hw04lrucache

import (
	"sync"

	"golang.org/x/exp/maps"
)

type Key string

type KeyVal struct {
	key Key
	val interface{}
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    *sync.Mutex
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if val, ok := c.items[key]; ok {
		val.Value.(*KeyVal).val = value
		c.queue.MoveToFront(val)
		return true
	}

	if c.capacity == c.queue.Len() {
		i := c.queue.Back()
		delete(c.items, i.Value.(*KeyVal).key)
		c.queue.Remove(i)
	}

	c.items[key] = c.queue.PushFront(&KeyVal{key, value})

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if val, ok := c.items[key]; ok {
		return val.Value.(*KeyVal).val, ok
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	maps.Clear(c.items)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mutex:    &sync.Mutex{},
	}
}
