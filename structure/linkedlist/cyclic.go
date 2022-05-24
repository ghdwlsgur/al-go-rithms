package linkedlist

import "fmt"

type Cyclic struct {
	Size int
	Head *Node
}

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

func (cl *Cyclic) Delete() bool {
	var deleted bool
	var prevItem, thisItem, nextItem *Node

	if cl.Size == 0 {
		return deleted
	}

	deleted = true
	thisItem = cl.Head
	nextItem = thisItem.Next
	prevItem = thisItem.Prev

	if cl.Size == 1 {
		cl.Head = nil
	} else {
		cl.Head = nextItem
		nextItem.Prev = prevItem
		prevItem.Next = nextItem
	}
	cl.Size--

	return deleted
}

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
