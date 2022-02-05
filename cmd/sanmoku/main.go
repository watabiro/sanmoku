package main

import (
	"fmt"

	"github.com/watabiro/sanmoku"
)

func main() {
	b := sanmoku.NewBoard()
	var p1 sanmoku.Player
	var p2 sanmoku.Player
	// p1 = sanmoku.MinimaxPlayer{}
	p1 = *sanmoku.NewMCTSPlayer(sanmoku.White)
	p2 = sanmoku.RandomPlayer{}
	playGame(b, p1, p2)
}

func playGame(b *sanmoku.Board, p1, p2 sanmoku.Player) {
	fmt.Println("start")
	b.Show()
	var move sanmoku.Move
	for {
		move = p1.BestMove(b)
		fmt.Println("selected move: ", move)
		b.Push(move)
		b.Show()
		if b.IsGameOver() {
			break
		}
		move = p2.BestMove(b)
		fmt.Println(move)
		b.Push(move)
		b.Show()
		if b.IsGameOver() {
			break
		}
	}
	fmt.Println("game over")
	if b.IsWin(sanmoku.Black) {
		fmt.Println("Win: Black")
	} else if b.IsWin(sanmoku.White) {
		fmt.Println("Win: White")
	} else {
		fmt.Println("Draw")
	}
}
