package queue

import (
	"container/list"
	"fmt"
)

type LQueue struct {
	queue *list.List
}

// 맨 뒤 요소에 큐를 추가
func (lq *LQueue) Enqueue(value any) {
	lq.queue.PushBack(value)
}

// 맨 앞 요소 삭제
func (lq *LQueue) Dequeue() error {
	if !lq.Empty() {
		element := lq.queue.Front()
		lq.queue.Remove(element)

		return nil
	}

	return fmt.Errorf("dequeue is empty we got an error")
}

// 맨 앞 요소 반환
func (lq *LQueue) Front() (any, error) {
	if !lq.Empty() {
		val := lq.queue.Front().Value
		return val, nil
	}

	return "", fmt.Errorf("error queue is empty")
}

// 맨 뒤 요소 반환
func (lq *LQueue) Back() (any, error) {
	if !lq.Empty() {
		val := lq.queue.Back().Value
		return val, nil
	}

	return "", fmt.Errorf("error queue is empty")
}

// 큐의 길이 반환
func (lq *LQueue) Len() int {
	return lq.queue.Len()
}

// 큐가 비어있을 경우 true, 그렇지 않을 경우 false 반환
func (lq *LQueue) Empty() bool {
	return lq.queue.Len() == 0
}
