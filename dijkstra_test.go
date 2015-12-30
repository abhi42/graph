package graph

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestDijkstraWith3Nodes(t *testing.T) {
	g := CreateEmptyGraph()
	v1 := CreateVertexInGraph(g, "v1")
	v2 := CreateVertexInGraph(g, "v2")
	v3 := CreateVertexInGraph(g, "v3")
	Join(v1, v2, 7)
	Join(v1, v3, 2)
	Join(v3, v2, 3)

	expectedShortestPath := []*Vertex{v1, v3, v2}
	fmt.Println("*****Expected shortest path*****")
	printVertexSlice(expectedShortestPath)

	shortestPath, shortestDistance := GetShortestPathBetween(v1, v2, g)
	fmt.Println("*****Calculated shortest path*****")
	printVertexSlice(shortestPath)

	assertShortestPath(t, expectedShortestPath, shortestPath, 5, shortestDistance)
}

func TestDijkstraWith5Nodes(t *testing.T) {
	g := CreateEmptyGraph()
	v1 := CreateVertexInGraph(g, "v1")
	v2 := CreateVertexInGraph(g, "v2")
	v3 := CreateVertexInGraph(g, "v3")
	v4 := CreateVertexInGraph(g, "v4")
	v5 := CreateVertexInGraph(g, "v5")
	Join(v1, v2, 7)
	Join(v1, v3, 2)
	Join(v3, v2, 3)
	Join(v2, v5, 1)
	Join(v3, v4, 4)
	Join(v4, v5, 3)

	expectedShortestPath := []*Vertex{v1, v3, v2, v5}
	fmt.Println("*****Expected shortest path*****")
	printVertexSlice(expectedShortestPath)

	shortestPath, shortestDistance := GetShortestPathBetween(v1, v5, g)
	fmt.Println("*****Calculated shortest path*****")
	printVertexSlice(shortestPath)

	assertShortestPath(t, expectedShortestPath, shortestPath, 6, shortestDistance)
}

func assertShortestPath(t *testing.T, expectedPath, calculatedPath []*Vertex, expectedDistance, calculatedDistance uint64) {
	if !areSlicesEqual(calculatedPath, expectedPath) {
		t.Errorf("Incorrect shortest path, expected ", expectedPath, ", got ", calculatedPath)
	}
	if calculatedDistance != expectedDistance {
		t.Errorf("Incorrect shortest distance, expected ", expectedDistance, ", got ", calculatedDistance)
	}
}

func printVertexSlice(s []*Vertex) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func areSlicesEqual(s1, s2 []*Vertex) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			fmt.Println(s1[i])
			fmt.Println(s2[i])
			return false
		}
	}
	return true
}

func TestSet(t *testing.T) {
	m := mapset.NewSet()
	aStr := "abc"
	anInt := 123
	m.Add(aStr)
	m.Add(anInt)

	if !m.Contains(aStr) {
		t.Errorf("Set should have contained %s", aStr)
	}

	if !m.Contains(anInt) {
		t.Errorf("Set should have contained %i", anInt)
	}
}

func TestStringComparison(t *testing.T) {
	str1 := "abc"
	str2 := "abc"
	if str1 != str2 {
		t.Errorf("Strings should have been equal")
	}
	ptr1 := &str1
	ptr2 := &str1

	if ptr1 == ptr2 {
		t.Logf("same pointer")
	} else {
		t.Logf("diff pointer")
	}
}

func TestPointers(t *testing.T) {
	pointerToV1 := new(Vertex)
	pointerToV1.name = "v1"
	t.Logf("v1 name: %s", pointerToV1.name)
	fmt.Println(reflect.TypeOf(pointerToV1))
	pointerToGraph := new(Graph)
	pointerToGraph.vertices = make(map[*Vertex]bool)
	pointerToGraph.vertices[pointerToV1] = true

	mapOfVertices := pointerToGraph.vertices

	fmt.Println("Printing map...")
	for key, _ := range mapOfVertices {
		fmt.Println(key.name)
		fmt.Println(pointerToV1 == key)
	}
}

func TestCompare(t *testing.T) {
	arr := make(SortableArr, 5, 5)
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = len(arr) - i
	}
	fmt.Println(arr)
	sort.Sort(arr)
	fmt.Println(arr)
}

func TestCompareVertices(t *testing.T) {
	input := make(SortableVertices, 5, 5)
	for i := 0; i < len(input); i++ {
		aVertex := CreateVertex("v" + strconv.Itoa(i+1))
		input[i].v = *aVertex
		input[i].distance = uint64(len(input) - i)
	}
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
	}
	sort.Sort(input)
	fmt.Println("After sorting")
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
	}
}

type Vertices struct {
	v        Vertex
	distance uint64
}

type SortableVertices []Vertices

func (v SortableVertices) Len() int {
	return len(v)
}
func (v SortableVertices) Less(i, j int) bool {
	return v[i].distance < v[j].distance
}
func (v SortableVertices) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (vertexStruct Vertices) String() string {
	return vertexStruct.v.name + ":" + strconv.FormatUint(vertexStruct.distance, 10)
}

type SortableArr []int

func (a SortableArr) Len() int {
	return len(a)
}

func (a SortableArr) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a SortableArr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
