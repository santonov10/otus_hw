package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity       int
	queue          List
	items          map[Key]*ListItem
	itemsKeysIndex map[*ListItem]Key
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if listItem, ok := c.items[key]; ok {
		listItem.Value = value

		c.queue.MoveToFront(listItem)
		return ok
	}

	insertItem := NewListItem(value)
	c.items[key] = insertItem
	c.itemsKeysIndex[insertItem] = key
	c.queue.PushFront(insertItem)
	if c.queue.Len() > c.capacity {
		removeItem := c.queue.Back()
		removeKey := c.itemsKeysIndex[removeItem]

		delete(c.items, removeKey)
		delete(c.itemsKeysIndex, removeItem)
		c.queue.Remove(removeItem)
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	resOk := false
	var res interface{}
	if listItem, ok := c.items[key]; ok {
		resOk = true
		res = listItem.Value
	}
	return res, resOk
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.itemsKeysIndex = make(map[*ListItem]Key, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity:       capacity,
		queue:          NewList(),
		items:          make(map[Key]*ListItem, capacity),
		itemsKeysIndex: make(map[*ListItem]Key, capacity),
	}
}
