package graph

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const Infinity = math.MaxUint64

type vertexComputedData struct {
	distanceFromSource                  uint64
	isVisited                           bool
	neighbourOnShortestDistanceToSource *Vertex
}

type vertexInfo struct {
	v    *Vertex
	data vertexComputedData
}

type verticesInfo []vertexInfo

func (ele verticesInfo) Len() int {
	return len(ele)
}
func (ele verticesInfo) Less(i, j int) bool {
	return ele[i].data.distanceFromSource < ele[j].data.distanceFromSource
}
func (ele verticesInfo) Swap(i, j int) {
	ele[i], ele[j] = ele[j], ele[i]
}

var vertices verticesInfo

func Dijkstra(g *Graph, source, dest *Vertex) ([]*Vertex, uint64) {
	doInit(g, source)
	process(source, dest)
	return buildShortestPath(source, dest)
}

func process(source, dest *Vertex) {
	for !areAllVerticesVisited() {
		current, offset := getUnvisitedVertexWithShortestDistanceFromSource()
		//		fmt.Println("Current: " + current.name)
		markAsVisited(current, offset)
		assignDistancesFromSourceToNeighbours(current, offset)
		printVertices()
	}
}

func areAllVerticesVisited() bool {
	for i := 0; i < len(vertices); i++ {
		if !vertices[i].data.isVisited {
			return false
		}
	}
	//	fmt.Println("All nodes have been visited")
	return true
}

func getOffsetInVertices(v *Vertex) int {
	for i := 0; i < len(vertices); i++ {
		if vertices[i].v == v {
			return i
		}
	}
	return -1
}

func assignDistancesFromSourceToNeighbours(current *Vertex, offset int) {
	distanceOfCurrentFromSource := vertices[offset].data.distanceFromSource

	for i := 0; i < len(current.neighbours); i++ {
		//		fmt.Println("Neighbour: " + current.neighbours[i].name + ", distance to " + current.name + " is " + strconv.FormatUint(current.neighbourDistances[i], 10))
		totalDistance := distanceOfCurrentFromSource + current.neighbourDistances[i]
		handleNeighbour(current, current.neighbours[i], getOffsetInVertices(current.neighbours[i]), totalDistance)
	}
}

func handleNeighbour(current, neighbour *Vertex, offsetOfNeighbour int, calculatedDistance uint64) {
	if calculatedDistance < vertices[offsetOfNeighbour].data.distanceFromSource {
		//		fmt.Println("Current: " + current.name + ", neighbour: " + neighbour.name + ", calculated distance: " + strconv.FormatUint(calculatedDistance, 10) + ", current distance: " + strconv.FormatUint(vertices[offsetOfNeighbour].data.distanceFromSource, 10))
		vertices[offsetOfNeighbour].data.distanceFromSource = calculatedDistance
		vertices[offsetOfNeighbour].data.neighbourOnShortestDistanceToSource = current
	}
}

func markAsVisited(node *Vertex, offset int) {
	vertices[offset].data.isVisited = true
}

func doInit(g *Graph, source *Vertex) {
	vertices = make(verticesInfo, len(g.vertices), len(g.vertices))
	var i = 0
	for pointerToV := range g.vertices {
		vertices[i].v = pointerToV
		vertices[i].data.isVisited = false
		if pointerToV == source {
			vertices[i].data.distanceFromSource = 0
		} else {
			vertices[i].data.distanceFromSource = Infinity
		}
		i++
	}
}

func getUnvisitedVertexWithShortestDistanceFromSource() (*Vertex, int) {
	sort.Sort(vertices)
	for i := 0; i < len(vertices); i++ {
		if !vertices[i].data.isVisited {
			return vertices[i].v, i
		}
	}
	return nil, -1
}

func (v vertexInfo) String() string {
	if v.data.neighbourOnShortestDistanceToSource != nil {
		return v.v.name + ", distance from source: " + strconv.FormatUint(v.data.distanceFromSource, 10) + ", neighbour on shortest path: " + v.data.neighbourOnShortestDistanceToSource.name
	}
	return v.v.name + ", distance from source: " + strconv.FormatUint(v.data.distanceFromSource, 10)
}

func printVertices() {
	//	fmt.Println("Printing vertices information...")
	for i := 0; i < len(vertices); i++ {
		//		fmt.Println(vertices[i].String())
	}
}

func buildShortestPath(source, dest *Vertex) ([]*Vertex, uint64) {
	shortestPath := make([]*Vertex, 0, 0)
	var shortestDistance uint64
	for i := 0; i < len(vertices); i++ {
		if vertices[i].v == dest {
			shortestPath = getShortestPathRecursively(shortestPath, vertices[i].v, i)
			shortestDistance = vertices[i].data.distanceFromSource
			break
		}
	}
	invertBuiltShortestPath(shortestPath)

	printShortestPath(shortestPath, shortestDistance)

	return shortestPath, shortestDistance
}

func printShortestPath(shortestPath []*Vertex, shortestDistance uint64) {
	fmt.Print("Shortest path (distance: ", shortestDistance, ") is: ")
	for i := 0; i < len(shortestPath); i++ {
		fmt.Print(shortestPath[i].name)
		if i < len(shortestPath)-1 {
			fmt.Print("-->")
		}
	}
	fmt.Println("")
}

func invertBuiltShortestPath(shortestPath []*Vertex) {
	for i := 0; i < len(shortestPath)-1-i; i++ {
		shortestPath[i], shortestPath[len(shortestPath)-1-i] = shortestPath[len(shortestPath)-1-i], shortestPath[i]
	}
}

func getShortestPathRecursively(shortestPath []*Vertex, v *Vertex, offset int) []*Vertex {
	shortestPath = append(shortestPath, v)
	newCurrent := vertices[offset].data.neighbourOnShortestDistanceToSource
	if newCurrent != nil {
		shortestPath = getShortestPathRecursively(shortestPath, newCurrent, getOffsetInVertices(newCurrent))
	}
	return shortestPath
}
