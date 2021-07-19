package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

var _ List = &list{}

type list struct {
	indexListItems map[*ListItem]bool
	firstListItem  *ListItem
	lastListItem   *ListItem
}

func (l *list) Len() int {
	return len(l.indexListItems)
}

func (l *list) Front() *ListItem {
	if l.Len() == 0 {
		return nil
	}

	return l.firstListItem
}

func (l *list) Back() *ListItem {
	if l.Len() == 0 {
		return nil
	}

	return l.lastListItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.Len() == 0 {
		return l.pushToEmptyList(v)
	}
	newListItem := NewListItem(v)
	secondListItem := l.firstListItem
	secondListItem.Prev = newListItem
	l.firstListItem = newListItem
	l.firstListItem.Next = secondListItem
	l.indexListItems[newListItem] = true
	return newListItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.Len() == 0 {
		return l.pushToEmptyList(v)
	}
	newListItem := NewListItem(v)
	preLastListItem := l.lastListItem
	preLastListItem.Next = newListItem
	l.lastListItem = newListItem
	l.lastListItem.Prev = preLastListItem
	l.indexListItems[newListItem] = true
	return newListItem
}

func (l *list) pushToEmptyList(v interface{}) *ListItem {
	newListItem := NewListItem(v)
	l.firstListItem = newListItem
	l.lastListItem = newListItem
	l.indexListItems[newListItem] = true
	return newListItem
}

func (l *list) Remove(i *ListItem) {
	l.panicIfNotExist(i)
	if l.lastListItem == i {
		l.lastListItem = i.Prev
	}
	if l.firstListItem == i {
		l.firstListItem = i.Next
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	delete(l.indexListItems, i)
}

func (l *list) MoveToFront(i *ListItem) {
	l.panicIfNotExist(i)
	if l.firstListItem != i {
		l.Remove(i)
		l.PushFront(i)
	}
}

func (l *list) panicIfNotExist(i *ListItem) {
	if !l.IsItemExist(i) {
		errorString := fmt.Sprintf("нет такого элемента в списке: %v", i)
		panic(errorString)
	}
}

func (l *list) IsItemExist(i *ListItem) bool {
	_, ok := l.indexListItems[i]
	return ok
}

func NewList() List {
	l := new(list)
	l.indexListItems = make(map[*ListItem]bool)
	return l
}

func NewListItem(v interface{}) *ListItem {
	newListItem := &ListItem{}
	switch t := v.(type) {
	case *ListItem:
		newListItem = v.(*ListItem)
	default:
		newListItem.Value = t
	}
	return newListItem
}
