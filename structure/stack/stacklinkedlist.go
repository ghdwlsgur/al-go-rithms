package stack

// 데이터와 다음 노드를 필드로 갖고 있는 Node 구조체
type Node struct {
	Val  any
	Next *Node
}

// 최상위 노드를 top 필드로, 길이를 length 필드로 가지는 Stack 구조체 (LIFO: Last In First Out)
type Stack struct {
	top    *Node
	length int
}

// 요소 추가
func (ll *Stack) push(n any) {
	newStack := &Node{}

	// 인수로 전달받은 n을 가지는 새로운 스택 추가
	newStack.Val = n
	// 현재 가장 위에 위치한 스택을 새로 생성한 스택 아래에 배치
	newStack.Next = ll.top

	// 새로운 스택을 가장 위에 배치
	ll.top = newStack
	ll.length++
}

// top 요소 삭제
func (ll *Stack) pop() any {
	result := ll.top.Val

	if ll.top.Next == nil { // 스택의 길이가 1일 경우
		ll.top = nil
	} else { // 스택의 길이가 1이 아닐 경우 2번째 위치한 스택을 제일 상단에 배치
		ll.top.Val, ll.top.Next = ll.top.Next.Val, ll.top.Next.Next
	}

	ll.length--
	return result
}

// stack의 길이가 0일 경우 true 반환, 그렇지 않을 경우 false 반환
func (ll *Stack) isEmpty() bool {
	return ll.length == 0
}

// stack의 길이 반환
func (ll *Stack) len() int {
	return ll.length
}

// 가장 위에 위치하는 스택의 요소를 반환
func (ll *Stack) peak() any {
	return ll.top.Val
}

// 스택의 각 요소 반환
func (ll *Stack) show() (in []any) {
	current := ll.top

	for current != nil {
		in = append(in, current.Val)
		current = current.Next
	}
	return
}
