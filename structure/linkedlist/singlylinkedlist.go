package linkedlist

import (
	"errors"
	"fmt"
)

// 노드의 길이값을 가지고 있는 length와 노드를 나타내는 Head 필드를 가지는 구조체
type Singly struct {
	length int
	Head   *Node
}

// 구조체를 생성하는 메서드 (생성자 함수)
func NewSingly() *Singly {
	return &Singly{}
}

// Head 노드에 새로운 노드 추가
func (ll *Singly) AddAtBeg(val any) {
	n := NewNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++
}

// Tail 노드에 새로운 노드 추가
func (ll *Singly) AddAtEnd(val any) {
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		ll.length++
		return
	}

	cur := ll.Head
	for ; cur.Next != nil; cur = cur.Next {
	}
	// 가장 마지막 노드에 새로운 노드를 추가
	cur.Next = n
	ll.length++
}

// Head 노드 삭제
func (ll *Singly) DelAtBeg() any {
	// 노드가 존재하지 않으면 -1 반환
	if ll.Head == nil {
		return -1
	}

	// 현재 위치 : Head 노드 (삭제할 노드)
	cur := ll.Head
	// 현재 위치(Head)에서 다음 노드를 Head로 지정
	ll.Head = cur.Next
	ll.length--

	// 삭제할 노드의 Val 반환
	return cur.Val
}

// Tail 노드 삭제
func (ll *Singly) DelAtEnd() any {
	// 노드가 존재하지 않으면 -1 반환
	if ll.Head == nil {
		return -1
	}

	// ll.Head는 존재하면서 ll.Head.Next 노드는 존재하지 않으므로
	// 길이가 1인 Head 뿐인 노드에서 DelAtBeg() 메서드 실행
	if ll.Head.Next == nil {
		return ll.DelAtBeg()
	}

	cur := ll.Head

	// 앞서 if의 조건문에서 ll.Head.Next == nil, 즉 cur.Next != nil이므로
	// cur.Next.Next != nil일 때까지 노드를 순회하면서 반복
	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	// 가장 마지막 노드인 Tail 노드를 retval에 저장 후 nil값을 대입
	retval := cur.Next.Val
	cur.Next = nil
	ll.length--
	return retval
}

// 현재 노드의 길이를 반환
func (ll *Singly) Count() int {
	return ll.length
}

func (ll *Singly) Reverse() {

	// Singly 구조체의 노드를 참조하는 prev, Next 변수 선언
	var prev, Next *Node

	cur := ll.Head

	// 헤드가 nil일 아닐 때까지 반복, 즉 마지막 노드까지 순회
	/*
		Before
		   ______       ___________
			| cur | ---- | cur.Next |

		After
			 ___________			 ______
			| cur.Next | ---- | cur |
	*/
	for cur != nil {
		Next = cur.Next // 변수에 cur.Next 담아두기 (temp)
		cur.Next = prev // cur.Next에 prev 대입 (temp)
		prev = cur      // 현재 노드를 prev에 저장
		cur = Next      // 현재 노드에 Next저장
	}

	// 반복문 순회 결과 가장 마지막 노드에 위치한 노드를 Head로 이동
	// cur은 반복문 순회 결고가 최종적으로 nil값이 들어가게 된다.
	ll.Head = prev
}

// left 인덱스부터 right 인덱스까지에 한하며 Reverse() 수행
// 즉, left <= index <= right
func (ll *Singly) ReversePartition(left, right int) error {
	// left, right값 검사
	err := ll.CheckRangeFromIndex(left, right)
	if err != nil {
		return err
	}
	// 임시노드 생성
	tmpNode := NewNode(-1)
	// 임시노드 다음 노드부터 Head 노드 대입
	tmpNode.Next = ll.Head
	pre := tmpNode
	// 임시노드부터 반복문 순회, left의 최소값인 1의 경우 pre = pre.Next
	// 즉 가장 처음 노드가 pre에 들어가게 되며 left 인덱스값에 맞춰서 최소 노드를 지정하게 된다.
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	// Reverse()와 동일
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	// tmpNode.Next = pre.Next
	ll.Head = tmpNode.Next
	return nil
}

// left <= index <= right
// left 값은 1보다 크고 right보다는 작아야하며 right 값은 노드의 길이보다 작아야한다.
// 위 조건을 만족하지 않을 경우 에러를 발생시키는 함수
func (ll *Singly) CheckRangeFromIndex(left, right int) error {
	if left > right {
		return errors.New("left boundary must smaller than right")
	} else if left < 1 {
		return errors.New("left boundary starts from the first node")
	} else if right > ll.length {
		return errors.New("right boundary cannot be greater than the length of the linked list")
	}
	return nil
}

// fmt.Print(cur.val, "")
func (ll *Singly) Display() {
	// 현재 노드 값 = 연결리스트 헤더(노드)
	// 노드가 nil이 아닐 때까지 반복
	// 현재 노드 순회가 끝나면 다음 노드를 순회하여 연결된 모든 노드 순회
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val, "")
	}
	fmt.Println()
}
