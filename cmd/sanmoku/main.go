package main

import (
	"fmt"

	"github.com/watabiro/sanmoku"
)

func main() {
	b := sanmoku.NewBoard()
	b.Show()
	var p1 sanmoku.Player
	var p2 sanmoku.Player
	// p1 := sanmoku.UserPlayer{Color: sanmoku.Black}
	// p2 := sanmoku.UserPlayer{Color: sanmoku.White}
	// p1 = sanmoku.RandomPlayer{Color: sanmoku.Black}
	p1 = sanmoku.NewMCTSPlayer(sanmoku.Black)
	p2 = sanmoku.RandomPlayer{Color: sanmoku.White}
	// p2 = sanmoku.NewMCTSPlayer()
	playGame(b, p1, p2)
}

func playGame(b *sanmoku.Board, p1, p2 sanmoku.Player) {

	for {
		fmt.Println("Black Turn")
		b.Push(p1.BestMove(b))
		b.Show()
		if b.IsGameOver() {
			break
		}
		fmt.Println("White Turn")
		b.Push(p2.BestMove(b))
		b.Show()
		if b.IsGameOver() {
			break
		}
	}
	fmt.Println("game over")
	if b.Turn == sanmoku.Black {
		fmt.Println("Win: White")
	} else {
		fmt.Println("Win: Black")
	}
}
