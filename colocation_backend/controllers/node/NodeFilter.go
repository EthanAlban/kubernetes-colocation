package node

import (
	"container/heap"
	v1 "k8s.io/api/core/v1"
)

type Node struct {
	CpuUsage float64
}

type NodeHeap []*Node

func buildNodePriorityQueue(nodes []*v1.Node) {
	heapNodes := NodeHeap{}
	heap.Init(&heapNodes)
}

func Filter(nodes []*v1.Node) {
	//FilteredNodes := make([]*Node,0)
	//for _, node := range nodes {
	//
	//}
}

func (h NodeHeap) Len() int { return len(h) }

func (h NodeHeap) Less(i, j int) bool { return Priority(h[i]) < Priority(h[j]) } // 这里决定 大小顶堆 现在是小顶堆

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *NodeHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(*Node))
}
