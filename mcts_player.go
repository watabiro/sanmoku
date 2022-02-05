package sanmoku

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	threshold = 0
	c         = 0
)

type Node struct {
	VisitCount    int
	Win           float64
	LegalMoves    []Move
	ChildIndicies []int
}

func NewNode(b *Board) *Node {
	node := new(Node)
	node.VisitCount = 0.0
	node.Win = 0.0
	node.LegalMoves = b.LegalMoves()
	node.ChildIndicies = nil
	return node
}

type MCTSPlayer struct {
	Color Color
	Nodes []*Node
}

func NewMCTSPlayer(color Color) *MCTSPlayer {
	p := new(MCTSPlayer)
	p.Nodes = make([]*Node, 1<<16)
	p.Color = color
	return p
}

func (p MCTSPlayer) ExpandNode(node *Node, b *Board) {
	node.ChildIndicies = make([]int, len(node.LegalMoves))
	for i, v := range node.LegalMoves {
		b.Push(v)
		index := p.findIndexFromBoard(b)
		node.ChildIndicies[i] = index
		if p.Nodes[index] == nil {
			p.Nodes[index] = NewNode(b)
		}
		b.Pop()
	}
}

func (p MCTSPlayer) BestMove(b *Board) Move {
	// 現局面のノード選択
	rootindex := p.findIndexFromBoard(b)
	if p.Nodes[rootindex] == nil {
		p.Nodes[rootindex] = NewNode(b)
	}
	root := p.Nodes[rootindex]

	// 子ノードの展開
	p.ExpandNode(root, b)

	// 探索
	for i := 0; i < 300000; i++ {
		p.search(root, b)
	}

	// 手の選択
	maxval := 0
	argmax := 0
	for i, index := range root.ChildIndicies {
		cnt := p.Nodes[index].VisitCount
		if cnt >= maxval {
			maxval = cnt
			argmax = i
		}
	}

	return root.LegalMoves[argmax]
}

func (p *MCTSPlayer) search(node *Node, b *Board) float64 {
	if b.IsGameOver() {
		if b.IsWin(b.Turn) {
			node.Win = 1.0
			return node.Win
		} else if b.IsDraw() {
			node.Win = 0.5
			return node.Win
		} else {
			node.Win = 0.0
			return node.Win
		}
	}
	if node.VisitCount < threshold {
		// rollout
		node.VisitCount++
		node.Win = rollout(b)
		return node.Win
	} else {
		// select next move
		// search again
		node.VisitCount++
		if node.ChildIndicies == nil {
			p.ExpandNode(node, b)
		}
		nextIdx := p.selectNextNode(node)
		nextMove := node.LegalMoves[nextIdx]
		nextNode := p.Nodes[node.ChildIndicies[nextIdx]]
		b.Push(nextMove)
		node.Win = 1.0 - p.search(nextNode, b)
		b.Pop()
		return node.Win
	}
}

func (p *MCTSPlayer) selectNextNode(node *Node) int {
	N := 0
	v := make([]float64, len(node.ChildIndicies))
	u := make([]float64, len(node.ChildIndicies))
	ucb := make([]float64, len(node.ChildIndicies))
	for _, idx := range node.ChildIndicies {
		N += p.Nodes[idx].VisitCount
	}
	for i, idx := range node.ChildIndicies {
		child := p.Nodes[idx]
		if child.VisitCount == 0 {
			v[i] = 0.0
			u[i] = 0.0
		} else {
			v[i] = child.Win / float64(child.VisitCount)
			u[i] = math.Sqrt(math.Log(float64(N)) / float64(child.VisitCount))
		}
		ucb[i] = v[i] + c*u[i]
		// fmt.Println(v[i], u[i], ucb[i])
	}
	return argMax(ucb)
}

func rollout(b *Board) float64 {
	turn := b.Turn
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 0
	for !b.IsGameOver() {
		// select move randomly
		moves := b.LegalMoves()
		move := moves[r.Intn(len(moves))]
		b.Push(move)
		count++
	}
	lose := true
	draw := false
	if b.IsWin(turn) {
		lose = false
	}
	if b.IsDraw() {
		draw = true
	}
	for i := 0; i < count; i++ {
		b.Pop()
	}
	if lose {
		return 0.0
	}
	if draw {
		return 0.5
	}
	return 1.0
}

func (p *MCTSPlayer) findIndexFromBoard(b *Board) int {
	idx := 0
	for i := 0; i < len(b.State); i++ {
		idx *= 3
		idx += b.State[i]
	}
	return idx
}

func argMax(values []float64) int {
	if len(values) < 1 {
		panic(fmt.Errorf("argMax needs non empty list"))
	}
	argmax := 0
	maxval := values[0]
	for i, v := range values {
		if maxval >= v {
			maxval = v
			argmax = i
		}
	}
	return argmax
}
