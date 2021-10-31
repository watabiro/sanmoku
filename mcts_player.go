package sanmoku

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	threshold = 10
	c         = 1
)

type Node struct {
	VisitCount    int
	WinCount      int
	LegalMoves    []Move
	ChildIndicies []int
}

func NewNode(b *Board) *Node {
	node := new(Node)
	node.VisitCount = 0
	node.WinCount = 0
	node.LegalMoves = b.LegalMoves()
	node.ChildIndicies = make([]int, len(node.LegalMoves))
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

func (p MCTSPlayer) BestMove(b *Board) Move {
	// 現局面のノード選択
	rootindex := p.findIndexFromBoard(b)
	if p.Nodes[rootindex] == nil {
		p.Nodes[rootindex] = NewNode(b)
	}
	root := p.Nodes[rootindex]

	// 子ノードの展開
	for i, v := range root.LegalMoves {
		b.Push(v)
		index := p.findIndexFromBoard(b)
		root.ChildIndicies[i] = index
		if p.Nodes[index] == nil {
			p.Nodes[index] = NewNode(b)
		}
		b.Pop()
	}

	// 探索
	for i := 0; i < 100000; i++ {
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
	if node.VisitCount < threshold {
		// rollout
		return 1.0 - rollout(b)
	} else {
		// select next move
		// search again
		next := p.selectNextNode(node)
		return p.search(next, b)
	}
}

func (p *MCTSPlayer) selectNextNode(node *Node) *Node {
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
			v[i] = float64(child.WinCount) / float64(child.VisitCount)
			u[i] = math.Sqrt(math.Log(float64(N)) / float64(child.VisitCount))
		}
		ucb[i] = v[i] + c*u[i]
	}
	argmax := argMax(ucb)

	return p.Nodes[node.ChildIndicies[argmax]]
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
	for i := 0; i < count; i++ {
		b.Pop()
	}
	// rollout開始時の手番とゲーム終了時の手番が同じ場合は負け
	// TODO: これじゃ引き分け判定できていない
	if turn == b.Turn {
		return 0.0
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
	argmax := len(values)
	maxval := values[0]
	for i, v := range values {
		if maxval >= v {
			maxval = v
			argmax = i
		}
	}
	return argmax
}
