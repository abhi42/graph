package graph

import (
	//	"github.com/deckarep/golang-set"
	//	"fmt"
	"strconv"
)

type Vertex struct {
	name string // immutable
	// instead of a map[*Vertex]uint64, 2 slices are used since we need to obtain the neighbours in a deterministic order.
	neighbours         []*Vertex
	neighbourDistances []uint64
}

type Graph struct {
	vertices map[*Vertex]bool
}

func createVertex(vertexName string) *Vertex {
	pointerToVertex := new(Vertex)
	pointerToVertex.name = vertexName
	pointerToVertex.neighbours = make([]*Vertex, 0)
	pointerToVertex.neighbourDistances = make([]uint64, 0)
	return pointerToVertex
}

func (v *Vertex) join(neighbour *Vertex, distance uint64) {
	v.neighbours = append(v.neighbours, neighbour)
	v.neighbourDistances = append(v.neighbourDistances, distance)

	neighbour.neighbours = append(neighbour.neighbours, v)
	neighbour.neighbourDistances = append(neighbour.neighbourDistances, distance)
}

func (v *Vertex) String() string {
	if len(v.neighbours) == 0 {
		return v.name
	}
	s1 := v.name + ", neighbours with distances: "
	for i := 0; i < len(v.neighbours); i++ {
		s2 := v.neighbours[i].name + ":" + strconv.FormatUint(v.neighbourDistances[i], 10)
		if i < len(v.neighbours)-1 {
			s2 = s2 + ", "
		}
		s1 = s1 + s2
	}
	return s1
}
