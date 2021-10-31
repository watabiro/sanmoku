package sanmoku

import (
	"math/rand"
	"time"
)

type RandomPlayer struct {
	Color Color
}

func (p RandomPlayer) BestMove(b *Board) Move {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	moves := b.LegalMoves()
	return moves[r.Intn(len(moves))]
}
