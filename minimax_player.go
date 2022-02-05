package sanmoku

import "fmt"

type MinimaxPlayer struct {
}

func (p MinimaxPlayer) BestMove(b *Board) Move {
	var val = -100000
	var res Move
	for _, move := range b.LegalMoves() {
		b.Push(move)
		tmp := -p.negaMax(b)
		fmt.Println(move, tmp)
		b.Pop()
		if tmp > val {
			val = tmp
			res = move
		}
	}
	return res
}
func (p MinimaxPlayer) negaMax(b *Board) int {
	if b.IsDraw() {
		return 0
	}
	color := b.Turn
	if b.IsWin(color) {
		return 10000
	}
	if b.IsGameOver() {
		return -10000
	}
	// 未決着
	val := -200000
	for _, move := range b.LegalMoves() {
		b.Push(move)
		tmp := -p.negaMax(b)
		b.Pop()
		if tmp > val {
			val = tmp
		}
	}
	return val
}
