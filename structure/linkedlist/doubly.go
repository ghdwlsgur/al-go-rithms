package linkedlist

import "fmt"

// Head 노드를 가르키는 이중 연결리스트 Doubly 구조체
type Doubly struct {
	Head *Node
}

// Doubly 구조체를 생성하는 생성자 메서드
func NewDoubly() *Doubly {
	return &Doubly{nil}
}

// Head 노드 앞에 노드 추가
func (ll *Doubly) AddAtBeg(val any) {
	n := NewNode(val) // 새로운 노드 생성
	n.Next = ll.Head  // 새로운 노드의 Next에 ll.Head 연결

	if ll.Head != nil {
		ll.Head.Prev = n
	}

	// ll.Head가 존재하지 않는다면 새로운 노드가 ll.Head가 된다.
	ll.Head = n
}

// 마지막 노드로 새로운 노드 추가
func (ll *Doubly) AddAtEnd(val any) {
	n := NewNode(val) // 새로운 노드 생성

	// 노드가 존재하지 않을 경우
	if ll.Head == nil {
		ll.Head = n
		return
	}

	cur := ll.Head
	// Head 노드에서 마지막 노드까지 반복
	for ; cur.Next != nil; cur = cur.Next {
	}
	// 마지막 노드의 Next에 새로운 노드 연결
	cur.Next = n
	// 새로운 노드의 Prev에 마지막 노드 연결
	n.Prev = cur
}

// 이중 연결리스트 최상단 노드 삭제
func (ll *Doubly) DelAtBeg() any {
	// 노드가 존재하지 않는다면 -1 반환
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next

	// ll.Head 노드의 다음 노드를 Head로 옮김

	// cur.Next의 이전 노드 연결을 해제하므로 이전 헤드 노드였던 ll.Head 삭제
	if ll.Head != nil {
		ll.Head.Prev = nil
	}

	return cur.Val
}

// 이중 연결리스트 최하단 노드 삭제
func (ll *Doubly) DelAtEnd() any {
	// 노드가 존재하지 않는다면 -1 반환
	if ll.Head == nil {
		return -1
	}

	// Head.Next가 nil일 경우 노드의 Size는 1이므로 ll.DelAtBeg()로 남은 노드 삭제
	if ll.Head.Next == nil {
		return ll.DelAtBeg()
	}

	cur := ll.Head
	// 최하단 노드까지 반복문으로 이동
	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	retval := cur.Next.Val
	// 최하단 노드 삭제
	cur.Next = nil
	return retval
}

// 노드 개수 카운트
func (ll *Doubly) Count() any {
	var ctr int = 0

	// 반복문이 노드를 순회하면서 노드 한 개를 순회할 때마다 ctr 값 1증가
	for cur := ll.Head; cur != nil; cur = cur.Next {
		ctr += 1
	}
	return ctr
}

// 노드 역순 배치
func (ll *Doubly) Reverse() {
	var Prev, Next *Node
	cur := ll.Head

	for cur != nil {
		Next = cur.Next // 현재 노드의 다음 노드를 Next에 저장
		cur.Next = Prev // 다음 노드의 Next에 빈 Node Prev 저장
		cur.Prev = Next // 빈 Node Prev에 현재 노드의 다음 노드 저장
		Prev = cur      // 현재 노드를 Prev에 저장
		cur = Next      // 현재 노드의 다음 노드를 cur에 저장
	}

	ll.Head = Prev
}

// 반복문을 통해 노드를 순회하면서 노드값 Print
func (ll *Doubly) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}

// 노드를 역으로 순회하면서 노드값 Print
func (ll *Doubly) DisplayReverse() {
	if ll.Head == nil {
		return
	}

	var cur *Node
	// 마지막 노드로 이동 // health check request
	for cur = ll.Head; cur.Next != nil; cur = cur.Next {
	}

	// 마지막 노드부터 첫 인덱스의 노드까지 역으로 순회하면서 Print
	for ; cur != nil; cur = cur.Prev {
		fmt.Print(cur.Val, " ")
	}
	fmt.Println()
}
