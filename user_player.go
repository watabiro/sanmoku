package sanmoku

import "fmt"

type UserPlayer struct {
	Color Color
}

func (p UserPlayer) BestMove(b *Board) Move {
	var file, rank int
	fmt.Printf("input file: ")
	_, err := fmt.Scan(&file)
	if err != nil {
		file = 0
	}
	for file < 1 || file > 3 {
		fmt.Printf("input file: ")
		_, err = fmt.Scan(&file)
		if err != nil {
			file = 0
		}
	}
	fmt.Printf("input rank: ")
	_, err = fmt.Scan(&rank)
	if err != nil {
		rank = 0
	}
	for rank < 1 || rank > 3 {
		fmt.Printf("input rank: ")
		fmt.Scan(&rank)
		_, err = fmt.Scan(&rank)
		if err != nil {
			rank = 0
		}

	}
	return Move{File: file, Rank: rank, Color: p.Color}
}
