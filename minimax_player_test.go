package sanmoku

import "testing"

func TestNegaMax1(t *testing.T) {
	b := NewBoard()
	b.Push(Move{1, 1, Black})
	b.Push(Move{1, 2, White})
	b.Push(Move{2, 1, Black})
	b.Push(Move{2, 2, White})
	p := MinimaxPlayer{}
	if val := p.negaMax(b); val != 10000 {
		b.Show()
		t.Errorf("expect 100 but got %v", val)
	}
}
func TestNegaMax2(t *testing.T) {
	b := NewBoard()
	b.Push(Move{1, 1, Black})
	b.Push(Move{1, 2, White})
	b.Push(Move{2, 1, Black})
	b.Push(Move{2, 2, White})
	b.Push(Move{3, 1, Black})
	p := MinimaxPlayer{}
	if val := p.negaMax(b); val != -10000 {
		b.Show()
		t.Errorf("expect -100 but got %v", val)
	}
}
