package piecetable

type Line struct {
	Buffer string
	nodes  []*PieceNode
}

func (l *Line) AppendNode(node *PieceNode) {
	l.nodes = append(l.nodes, node)
}

func (l *Line) String() string {
	data := l.Buffer

	for _, mods := range l.nodes {
		data = data[:mods.Start] + mods.Data + data[mods.Start:]
	}

	return data
}
