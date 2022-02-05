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
	NoStone = iota
	BlackStone
	WhiteStone
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
	if b.Turn == Black {
		fmt.Println("next turn is black")
	} else {
		fmt.Println("next turn is white")
	}
}

func (b *Board) LegalMoves() []Move {
	moves := make([]Move, 0, 16)
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if b.State[i*WIDTH+j] == NoStone {
				moves = append(moves, Move{File: WIDTH - j, Rank: i + 1, Color: b.Turn})
			}
		}
	}
	return moves
}

func (b *Board) Push(move Move) {
	if move.Color == Black {
		b.State[move.ToIndex()] = BlackStone
	} else {
		b.State[move.ToIndex()] = WhiteStone
	}
	b.MoveHistory = append(b.MoveHistory, move)
	if b.Turn == Black {
		b.Turn = White
	} else {
		b.Turn = Black
	}
}

func (b *Board) Pop() {
	move := b.MoveHistory[len(b.MoveHistory)-1]
	b.State[WIDTH*move.Rank-move.File] = NoStone
	b.MoveHistory = b.MoveHistory[:len(b.MoveHistory)-1]
	if b.Turn == Black {
		b.Turn = White
	} else {
		b.Turn = Black
	}
}

func (b *Board) IsWin(color Color) bool {
	targets := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	var stone int
	if color == Black {
		stone = BlackStone
	} else {
		stone = WhiteStone
	}
	for _, target := range targets {
		count := 0
		for _, i := range target {
			if b.State[i] == stone {
				count++
			}
		}
		if count == 3 {
			return true
		}
	}
	return false
}

func (b *Board) IsDraw() bool {
	if b.IsWin(Black) || b.IsWin(White) {
		return false
	}
	count := 0
	for i := range b.State {
		if b.State[i] != NoStone {
			count++
		}
	}
	return count == 9
}

func (b *Board) IsGameOver() bool {
	if b.IsWin(Black) || b.IsWin(White) || b.IsDraw() {
		return true
	}
	return false
}
