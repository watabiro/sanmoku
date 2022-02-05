package sanmoku

import (
	"fmt"
)

type Color int

const (
	Black Color = iota
	White
)

const (
	HEIGHT = 3
	WIDTH  = 3
)

type Board struct {
	State       []int
	MoveHistory []Move
	Turn        Color
}

func NewBoard() *Board {
	b := new(Board)
	b.State = make([]int, 9)
	for i := range b.State {
		b.State[i] = 0
	}
	b.MoveHistory = nil
	b.Turn = Black
	return b
}

func (b *Board) Show() {
	s := [3]string{"   ", " o ", " x "}
	fmt.Printf("%s|%s|%s\n", s[b.State[0]], s[b.State[1]], s[b.State[2]])
	fmt.Println("-----------")
	fmt.Printf("%s|%s|%s\n", s[b.State[3]], s[b.State[4]], s[b.State[5]])
	fmt.Println("-----------")
	fmt.Printf("%s|%s|%s\n", s[b.State[6]], s[b.State[7]], s[b.State[8]])
}

func (b *Board) LegalMoves() []Move {
	moves := make([]Move, 0, 16)
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if b.State[i*WIDTH+j] == 0 {
				moves = append(moves, Move{File: WIDTH - j, Rank: i + 1, Color: b.Turn})
			}
		}
	}
	return moves
}

func (b *Board) Push(move Move) {
	if move.Color == Black {
		b.State[WIDTH*move.Rank-move.File] = 1
	} else {
		b.State[WIDTH*move.Rank-move.File] = 2
	}
	b.MoveHistory = append(b.MoveHistory, move)
	b.Turn ^= 1
}

func (b *Board) Pop() {
	move := b.MoveHistory[len(b.MoveHistory)-1]
	b.State[WIDTH*move.Rank-move.File] = 0
	b.MoveHistory = b.MoveHistory[:len(b.MoveHistory)-1]
}

func (b *Board) IsGameOver() bool {
	pieceCount := 0
	for i := range b.State {
		if b.State[i] > 0 {
			pieceCount++
		}
	}
	if pieceCount == 9 {
		return true
	}
	// horizontal
	for i := 0; i < 3; i++ {
		oCount := 0
		xCount := 0
		for j := 0; j < 3; j++ {
			if b.State[i*3+j] == 1 {
				oCount++
			} else if b.State[i*3+j] == 2 {
				xCount++
			}
		}
		if oCount == 3 || xCount == 3 {
			return true
		}
	}
	// vertical
	for i := 0; i < 3; i++ {
		oCount := 0
		xCount := 0
		for j := 0; j < 3; j++ {
			if b.State[i+3*j] == 1 {
				oCount++
			} else if b.State[i+3*j] == 2 {
				xCount++
			}
		}
		if oCount == 3 || xCount == 3 {
			return true
		}
	}
	// diagonal lower right
	oCount := 0
	xCount := 0
	for i := 0; i < 3; i++ {
		if b.State[i*4] == 1 {
			oCount++
		} else if b.State[i*4] == 2 {
			xCount++
		}
	}
	if oCount == 3 || xCount == 3 {
		return true
	}
	// diagonal upper left
	oCount = 0
	xCount = 0
	for i := 0; i < 3; i++ {
		if b.State[i*2+2] == 1 {
			oCount++
		} else if b.State[i*2+2] == 2 {
			xCount++
		}
	}
	if oCount == 3 || xCount == 3 {
		return true
	}
	return false
}
