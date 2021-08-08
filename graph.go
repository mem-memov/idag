package idag

type Graph struct {
	entries []uint
}

func (g *Graph) Append(adjacentVertices []uint) uint {
	vertex := uint(len(g.entries))
	for _, adjacentVertex := range adjacentVertices {
		g.entries = append(g.entries, adjacentVertex)
	}
	g.entries = append(g.entries, vertex)
	return vertex
}

func (g *Graph) Read(vertex uint) []uint {
	adjacentVertices := make([]uint, 0)
	var i uint = 0
	for {
		adjacentVertex := g.entries[vertex + i]
		if adjacentVertex == vertex {
			break
		}
		adjacentVertices = append(adjacentVertices, adjacentVertex)
		i++
	}
	return adjacentVertices
}