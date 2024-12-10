package pathfinding

import (
	"github.com/limwa/advent-of-code-2024/lib/spatial"
)

type PathfindingItem struct {
	Value int

	// Additional information, not relevant for pathfinding
	Position spatial.Vec2D
}

type PathfindingHeap []*PathfindingItem

func (h *PathfindingHeap) Len() int {
	return len(*h)
}

func (h *PathfindingHeap) Less(i, j int) bool {
	heap := *h
	return heap[i].Value > heap[j].Value
}

func (h *PathfindingHeap) Swap(i, j int) {
	heap := *h
	heap[i], heap[j] = heap[j], heap[i]
}

func (h *PathfindingHeap) Push(x any) {
	item := x.(*PathfindingItem)
	*h = append(*h, item)
}

func (h *PathfindingHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
