package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if listItem, ok := c.items[key]; ok {
		listItem.Value.(*cacheItem).value = value

		c.queue.MoveToFront(listItem)
		return ok
	}
	cItem := &cacheItem{key: key, value: value}
	insertItem := NewListItem(cItem)
	c.items[key] = insertItem
	c.queue.PushFront(insertItem)
	if c.queue.Len() > c.capacity {
		removeItem := c.queue.Back()
		removeKey := removeItem.Value.(*cacheItem).key
		delete(c.items, removeKey)
		c.queue.Remove(removeItem)
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if listItem, ok := c.items[key]; ok {
		return listItem.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
