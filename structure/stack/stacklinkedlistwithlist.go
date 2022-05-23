package stack

import (
	"container/list"
	"fmt"
)

type SList struct {
	stack *list.List
}

func (sl *SList) Push(val any) {
	sl.stack.PushFront(val)
}

func (sl *SList) Peak() (any, error) {
	if !sl.Empty() {
		element := sl.stack.Front()
		return element.Value, nil
	}
	return "", fmt.Errorf("stack list is empty")
}

func (sl *SList) Pop() (any, error) {
	if !sl.Empty() {
		element := sl.stack.Front()
		sl.stack.Remove(element)
		return element.Value, nil
	}
	return "", fmt.Errorf("stack list is empty")
}

func (sl *SList) Length() int {
	return sl.stack.Len()
}

func (sl *SList) Empty() bool {
	return sl.stack.Len() == 0
}
