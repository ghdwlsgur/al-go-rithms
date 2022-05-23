package linkedlist

import (
	"reflect"
	"testing"
)

func TestSingly(t *testing.T) {
	list := NewSingly()
	list.AddAtBeg(1) // 1
	list.AddAtBeg(2) // 2 -- 1
	list.AddAtBeg(3) // 3 -- 2 -- 1

	t.Run("Test AddAtBeg()", func(t *testing.T) {
		want := []any{3, 2, 1}
		got := []any{}
		current := list.Head
		got = append(got, current.Val)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	list.AddAtEnd(4) // 3 -- 2 -- 1 -- 4

	t.Run("Test AddAtEnd()", func(t *testing.T) {
		want := []any{3, 2, 1, 4}
		got := []any{}
		current := list.Head
		got = append(got, current.Val)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtBeg()", func(t *testing.T) {
		want := any(3)
		got := list.DelAtBeg() // 3
		// list.Display() => 2 -- 1 -- 4
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test DelAtEnd()", func(t *testing.T) {
		want := any(4)
		got := list.DelAtEnd() // 4
		// list.Display() => 2 -- 1
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Count()", func(t *testing.T) {
		want := 2
		got := list.Count() // 2
		// list.Display() => 2 -- 1
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	list2 := Singly{}
	list2.AddAtBeg(1) // 1
	list2.AddAtBeg(2) // 2 -- 1
	list2.AddAtBeg(3) // 3 -- 2 -- 1
	list2.AddAtBeg(4) // 4 -- 3 -- 2 -- 1
	list2.AddAtBeg(5) // 5 -- 4 -- 3 -- 2 -- 1
	list2.AddAtBeg(6) // 6 -- 5 -- 4 -- 3 -- 2 -- 1

	// []any{6, 5, 4, 3, 2, 1}
	// =>
	// []any{1, 2, 3, 4, 5, 6}

	t.Run("Test Reverse", func(t *testing.T) {
		want := []any{1, 2, 3, 4, 5, 6}
		got := []any{}
		list2.Reverse() // 1 -- 2 -- 3 -- 4 -- 5 -- 6
		current := list2.Head
		got = append(got, current.Val)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test ReversePartition()", func(t *testing.T) {
		want := []any{1, 5, 4, 3, 2, 6}
		got := []any{}
		err := list2.ReversePartition(2, 5)

		if err != nil {
			t.Errorf("Incorrect boundary conditions entered%v", err)
		}
		current := list2.Head
		got = append(got, current.Val)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
