package queue

// 노드 구조체 선언
type Node struct {
	Data any   // 데이터
	Next *Node // 디음 노드
}

// head 노드와 tail 노드와 길이값을 가지는 큐 구조체
type Queue struct {
	head   *Node
	tail   *Node
	length int
}

// 큐 추가
func (ll *Queue) enqueue(n any) {
	// 노드 구조체 선언
	var newNode Node
	newNode.Data = n

	// tail 노드가 존재한다면 tail 노드 다음 노드에 새로운 노드 연결
	if ll.tail != nil {
		ll.tail.Next = &newNode
	}

	// 새로운 노드를 tail에 저장
	ll.tail = &newNode

	// head가 nil이라면 가장 첫 노드에 새로운 노드 추가
	if ll.head == nil {
		ll.head = &newNode
	}
	ll.length++
}

// 큐 삭제
func (ll *Queue) dequeue() any {
	// 큐가 비어있다면 삭제할 노드가 없으므로 -1 반환
	if ll.isEmpty() {
		return -1
	}
	// 큐가 비어있지 않다면 data 변수에 ll.head.Data 초기화
	data := ll.head.Data

	// 현재 가장 첫 노드인 head 노드는 삭제되므로 ll.head에 다음 노드 저장
	ll.head = ll.head.Next

	// head 노드가 비어있다면 tail 노드 nil 처리
	if ll.head == nil {
		ll.tail = nil
	}

	ll.length--
	// 삭제된 노드 데이터 반환
	return data
}

// 큐가 비어있다면 true 반환, 그렇지 않을 경우 false 반환
func (ll *Queue) isEmpty() bool {
	return ll.length == 0
}

// 큐의 길이 반환
func (ll *Queue) len() int {
	return ll.length
}

// 큐의 가장 첫 번째 head 반환
func (ll *Queue) frontQueue() any {
	return ll.head.Data
}

// 큐의 가장 마지막 tail 반환
func (ll *Queue) backQueue() any {
	return ll.tail.Data
}
