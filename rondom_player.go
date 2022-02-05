package sanmoku

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomPlayer struct {
}

func (p RandomPlayer) BestMove(b *Board) Move {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	moves := b.LegalMoves()
	fmt.Println("candidates: ", moves)
	return moves[r.Intn(len(moves))]
}
