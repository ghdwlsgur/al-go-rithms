package queue

// 배열 ListQueue (FIFO: First In-First Out)
var ListQueue []any

// 큐 추가
func EnQueue(n any) {
	ListQueue = append(ListQueue, n)
}

// 큐 삭제
func DeQueue() any {
	// 삭제되는 가장 맨 앞에 오는 요소
	data := ListQueue[0]
	// 가장 맨 앞 요소를 제외한 나머지 요소를 ListQueue에 저장
	ListQueue = ListQueue[1:]
	return data
}

// 0번째 요소 반환
func FrontQueue() any {
	return ListQueue[0]
}

// 마지막 요소 반환
func BackQueue() any {
	return ListQueue[len(ListQueue)-1]
}

// 리스트큐의 길이 반환
func LenQueue() int {
	return len(ListQueue)
}

// 리스트큐의 길이가 0일 때 true 반환, 그렇지 않을 경우 false 반환
func IsEmptyQueue() bool {
	return len(ListQueue) == 0
}
