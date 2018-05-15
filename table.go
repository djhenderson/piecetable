package piecetable

import (
	"fmt"
	"strings"
)

type PieceTable struct {
	Lines []*Line
}

func MakePieceTable(data string) *PieceTable {
	readStrings := strings.Split(data, "\n")

	lines := make([]*Line, len(readStrings))
	for idx, data := range readStrings {
		lines[idx] = &Line{
			data,
			[]*PieceNode{},
		}
	}

	return &PieceTable{
		lines,
	}
}

func (p *PieceTable) Insert(val string, line int, idx int) {
	node := NewPiece(val, line, idx)
	p.Lines[line].AppendNode(node)
}

func (p *PieceTable) Line(idx int) string {
	return p.Lines[idx].String()
}

func (p *PieceTable) String() string {
	var result string
	for idx, line := range p.Lines {
		if idx > 0 {
			result += string('\n')
		}
		result += fmt.Sprintf(line.String())
	}
	return result
}

func (p *PieceTable) Print() {
	for _, line := range p.Lines {
		fmt.Println(line.Buffer)
	}
}
