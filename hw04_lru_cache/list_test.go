package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("fill", func(t *testing.T) {
		l := NewList()

		i1 := l.PushFront(1)
		require.Equal(t, i1, l.Front())
		require.Equal(t, i1, l.Back())

		i2 := l.PushFront(2)
		require.Equal(t, i2, l.Front())
		require.Equal(t, i1, l.Back())
		require.Equal(t, i1, l.Front().Next)
		require.Equal(t, i2, l.Back().Prev)

		i3 := l.PushFront(3)

		require.Equal(t, i3, l.Front())
		require.Equal(t, i2, l.Front().Next)
		require.Equal(t, i1, l.Front().Next.Next)
		require.Equal(t, i1, l.Back())
		require.Equal(t, i2, l.Back().Prev)
		require.Equal(t, i3, l.Back().Prev.Prev)

		l = NewList()

		i1 = l.PushBack(1)
		require.Equal(t, i1, l.Front())
		require.Equal(t, i1, l.Back())
	})

	t.Run("remove", func(t *testing.T) {
		l := NewList()

		i1 := l.PushFront(1)
		i2 := l.PushFront(2)
		i3 := l.PushFront(3)
		i4 := l.PushFront(4)

		l.Remove(i3)
		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{4, 2, 1}, elems)

		l.Remove(i4)
		elems = make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{2, 1}, elems)

		l.Remove(i1)
		l.Remove(i2)
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("move to front", func(t *testing.T) {
		l := NewList()

		l.PushFront(1)
		i2 := l.PushFront(2)
		l.PushFront(3)

		l.MoveToFront(i2)
		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{2, 3, 1}, elems)
	})
}
