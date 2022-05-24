package stack

import (
	"container/list"
	"fmt"
)

// 컨테이너 리스트를 참조하는 stack 필드를 가지는 SList 구조체
type SList struct {
	stack *list.List
}

// 가장 맨 앞 요소에 전달받은 파라미터 요소 추가
func (sl *SList) Push(val any) {
	sl.stack.PushFront(val)
}

// 가장 맨 앞 요소 반환
func (sl *SList) Peak() (any, error) {
	if !sl.Empty() {
		element := sl.stack.Front()
		return element.Value, nil
	}
	return "", fmt.Errorf("stack list is empty")
}

// 가장 맨 앞 요소 삭제
func (sl *SList) Pop() (any, error) {
	if !sl.Empty() {
		element := sl.stack.Front()
		sl.stack.Remove(element)
		return element.Value, nil
	}
	return "", fmt.Errorf("stack list is empty")
}

// 스택의 길이 반환
func (sl *SList) Length() int {
	return sl.stack.Len()
}

// 스택의 길이가 0일 경우 true 반환, 그렇지 않을 경우 false 반환
func (sl *SList) Empty() bool {
	return sl.stack.Len() == 0
}
