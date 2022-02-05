package sanmoku

type Move struct {
	File  int
	Rank  int
	Color Color
}

func (m Move) ToIndex() int {
	return WIDTH*m.Rank - m.File
}
