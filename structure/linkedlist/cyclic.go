package linkedlist

import "fmt"

// 원형 연결리스트의 노드 사이즈를 가지는 Size 필드와 Head 노드를 가르키는 Head 필드를 가지는 Cyclic 구조체
type Cyclic struct {
	Size int
	Head *Node
}

// Cyclic 구조체를 생성하는 생성자 메서드
func NewCyclic() *Cyclic {
	return &Cyclic{0, nil}
}

// cyclic_add.png 참고
func (cl *Cyclic) Add(val any) {
	n := NewNode(val)
	cl.Size++
	if cl.Head == nil {
		n.Prev = n
		n.Next = n
		cl.Head = n
	} else {
		n.Prev = cl.Head.Prev
		n.Next = cl.Head
		cl.Head.Prev.Next = n
		cl.Head.Prev = n
	}
}

// 인수로 전달받은 places가 양수이면 오른쪽으로 places 값만큼 index 이동
// 반대로 음수이면 왼쪽으로 places 값만큼 index 이동 후 이동한 노드의 인덱스 반환
// cyclic_rotate.png 참고

/*
노드가 3개인 원형 연결리스트 cl의 cl.Size = 3, 파라미터로 전달받은 places 값은 1이라고 가정할 때,
else {
	for i:= 0; i < places; i++ {
		cl.Head = cl.Head.Next
	}
}
다음 반복문을 거쳐 cl.Head가 오른쪽으로 한 칸 이동하게 된다.

노드각 3개인 원형 연결리스트 cl의 cl.Size = 3, 파라미터로 전달받은 places 값은 -1이라고 가정할 때,
multiple := 2
places = 5
places = 2
places = 1

for i := 0; i < 1; i++ {
	cl.Head = cl.Head.Prev
}
*/
func (cl *Cyclic) Rotate(places int) {
	if cl.Size > 0 {
		if places < 0 {
			multiple := cl.Size - 1 - places/cl.Size
			places += multiple * cl.Size
		}
		places %= cl.Size

		if places > cl.Size/2 {
			places = cl.Size - places
			for i := 0; i < places; i++ {
				cl.Head = cl.Head.Prev
			}
		} else if places == 0 {
			return
		} else {
			for i := 0; i < places; i++ {
				cl.Head = cl.Head.Next
			}
		}
	}
}

// Head 노드 삭제
func (cl *Cyclic) Delete() bool {
	var deleted bool // bool 기본값 false
	fmt.Printf("deleted:%v", deleted)
	var prevItem, thisItem, nextItem *Node

	// 삭제할 노득가 없기 때문에 false 반환
	if cl.Size == 0 {
		return deleted
	}

	deleted = true
	thisItem = cl.Head       // 현재 노드
	nextItem = thisItem.Next // 다음 노드
	prevItem = thisItem.Prev // 이전 노드

	if cl.Size == 1 { // Size가 1이라면 cl.Head nil 처리
		cl.Head = nil
	} else {
		cl.Head = nextItem       // 헤드를 다음 노드로 옮기고
		nextItem.Prev = prevItem // 이전 노드를 다음 노드의 이전 노드로 연결
		prevItem.Next = nextItem // 다음 노드를 이전 노드의 다음 노드로 연결
	}
	// 중간에 껴있던 (cl.Head)현재 노드는 사라지게 된다.
	cl.Size--

	return deleted
}

// 노드 전체 삭제
func (cl *Cyclic) Destroy() {
	for cl.Delete() {
		continue
	}
}

// Show list body
func (cl *Cyclic) Walk() *Node {
	var start *Node
	start = cl.Head

	for i := 0; i < cl.Size; i++ {
		fmt.Printf("%v \n", start.Val)
		start = start.Next
	}
	return start
}

// Josephus_problem.png 참고
func JosephusProblem(cl *Cyclic, k int) int {
	for cl.Size > 1 {
		fmt.Printf("%v\n", cl.Size)
		cl.Rotate(k)
		cl.Delete()
		cl.Rotate(-1)
		fmt.Printf("value: %v\n", cl.Head.Val.(int))
	}
	retval := cl.Head.Val.(int)
	cl.Destroy()
	return retval
}
