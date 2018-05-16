package piecetable

type Line struct {
	Buffer string
	parent *PieceTable
	mods   map[int]bool
}

func NewLine(data string, parent *PieceTable) *Line {
	return &Line{
		data, parent, map[int]bool{},
	}
}

func (l *Line) AppendNode(node *PieceNode) {
	l.mods[len(l.parent.nodes)] = true
	l.parent.nodes = append(l.parent.nodes, node)
}

func (l *Line) Split(idx int) string {
	rightPart := l.Buffer[idx:]
	l.Buffer = l.Buffer[:idx]
	return rightPart
}

func (l *Line) Substr(idx, length int) string {
	return l.Buffer[idx : idx+length]
}

func (l *Line) Len() int {
	return len(l.String())
}

func (l *Line) String() string {
	data := l.Buffer

	for index, val := range l.mods {
		if !val {
			continue
		}

		mod := l.parent.nodes[index]

		if mod.Length > 0 {
			data = data[:mod.Start] + mod.Data + data[mod.Start:]
		} else {
			data = data[:mod.Start-1] + data[mod.Start:]
		}
	}

	return data
}
