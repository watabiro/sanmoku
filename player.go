package sanmoku

type Player interface {
	BestMove(b *Board) Move
}
