package sanmoku

import (
	"fmt"
	"testing"
)

func TestIsWin1(t *testing.T) {
	b := NewBoard()
	b.Push(Move{1, 1, Black})
	b.Push(Move{1, 2, White})
	b.Push(Move{2, 1, Black})
	b.Push(Move{2, 2, White})
	b.Push(Move{3, 1, Black})
	if !b.IsWin(Black) {
		b.Show()
		fmt.Println(b.IsWin(Black))
		t.Error()
	}
}
func TestIsDraw(t *testing.T) {
	b := NewBoard()
	b.Push(Move{1, 1, Black})
	b.Push(Move{1, 2, White})
	b.Push(Move{2, 1, Black})
	b.Push(Move{2, 2, White})
	b.Push(Move{3, 1, Black})
	if b.IsDraw() {
		t.Error()
	}
}

func TestPush(t *testing.T) {
	b := NewBoard()
	b.Push(Move{1, 1, Black})
	if b.Turn != White {
		b.Show()
		t.Error()
	}
}
