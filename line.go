package piecetable

type Line struct {
	Buffer string
	parent *PieceTable
	mods   map[int]bool
}

func (l *Line) AppendNode(node *PieceNode) {
	l.mods[len(l.parent.nodes)] = true
	l.parent.nodes = append(l.parent.nodes, node)
}

func (l *Line) String() string {
	data := l.Buffer

	for index, _ := range l.mods {
		mod := l.parent.nodes[index]
		if mod.Length > 0 {
			data = data[:mod.Start] + mod.Data + data[mod.Start:]
		} else {
			data = data[:mod.Start-1] + data[mod.Start:]
		}
	}

	return data
}
