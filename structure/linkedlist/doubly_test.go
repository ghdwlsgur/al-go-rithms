package linkedlist

import (
	"reflect"
	"testing"
)

func TestDoubly(t *testing.T) {
	newList := NewDoubly()

	newList.AddAtBeg(1)
	newList.AddAtBeg(2)
	newList.AddAtBeg(3)

	//  ____									 ____ 							 	  ____
	// | 3 |(Next) <==> (Prev)| 2 |(Next) <==> (Prev)| 1 |

	t.Run("Test AddAtBeg", func(t *testing.T) {
		wantNext := []int{3, 2, 1}
		wantPrev := []int{1, 2, 3}
		got := []int{}

		current := newList.Head

		got = append(got, current.Val.(int))

		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val.(int))
		}

		if !reflect.DeepEqual(got, wantNext) {
			t.Errorf("got: %v, want: %v", got, wantNext)
		}

		got = []int{}
		got = append(got, current.Val.(int))

		for current.Prev != nil {
			current = current.Prev
			got = append(got, current.Val.(int))
		}

		if !reflect.DeepEqual(got, wantPrev) {
			t.Errorf("got: %v, want: %v", got, wantPrev)
		}
	})

	newList.AddAtEnd(4) // 3 2 1 4

	t.Run("Test AddAtEnd", func(t *testing.T) {
		want := []int{3, 2, 1, 4}
		got := []int{}
		current := newList.Head
		got = append(got, current.Val.(int))
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val.(int))
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtBeg", func(t *testing.T) {
		want := 3
		got := newList.DelAtBeg()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtEnd", func(t *testing.T) {
		want := 4
		got := newList.DelAtEnd()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Count", func(t *testing.T) {
		want := 2
		got := newList.Count()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	newList2 := NewDoubly()
	newList2.AddAtBeg(1)
	newList2.AddAtBeg(2)
	newList2.AddAtBeg(3)
	t.Run("Test Reverse", func(t *testing.T) {
		want := []int{1, 2, 3}
		got := []int{}
		newList2.Reverse()
		current := newList2.Head
		got = append(got, current.Val.(int))
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val.(int))
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
