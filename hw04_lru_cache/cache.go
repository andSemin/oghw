package hw04lrucache

import (
	"golang.org/x/exp/maps"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	pointers map[*ListItem]Key
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if val, ok := c.items[key]; ok {
		val.Value = value
		c.queue.MoveToFront(val)
		return true
	}

	if c.capacity == c.queue.Len() {
		i := c.queue.Back()
		delete(c.items, c.pointers[i])
		delete(c.pointers, i)
		c.queue.Remove(i)
	}

	c.items[key] = c.queue.PushFront(value)
	c.pointers[c.items[key]] = key

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if val, ok := c.items[key]; ok {
		return val.Value, ok
	}
	return nil, false
}

func (c *lruCache) Clear() {
	maps.Clear(c.items)
	maps.Clear(c.pointers)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		pointers: make(map[*ListItem]Key, capacity),
	}
}
