package graph

import (
//	"fmt"
)

func CreateEmptyGraph() *Graph {
	pointerToGraph := new(Graph)
	pointerToGraph.vertices = make(map[*Vertex]bool)
	return pointerToGraph
}

func CreateVertexInGraph(g *Graph, vertexName string) *Vertex {
	pointerToVertex := CreateVertex(vertexName)
	g.vertices[pointerToVertex] = true
	return pointerToVertex
}

func CreateVertex(vertexName string) *Vertex {
	return createVertex(vertexName)
}

func Join(v1, v2 *Vertex, distance uint64) {
	v1.join(v2, distance)
}

func GetShortestPathBetween(v1, v2 *Vertex, g *Graph) ([]*Vertex, uint64) {
	return Dijkstra(g, v1, v2)
}
