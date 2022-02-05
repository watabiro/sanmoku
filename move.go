package sanmoku

import (
	"fmt"
)

type Move struct {
	File  int
	Rank  int
	Color Color
}

func (m Move) ToIndex() int {
	return WIDTH*m.Rank - m.File
}

func (m Move) String() string {
	color := "W"
	if m.Color == Black {
		color = "B"
	}
	return fmt.Sprintf("%d%d%s", m.File, m.Rank, color)
}
