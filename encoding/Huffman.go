package encoding

import (
	"container/heap"
	"fmt"
	"io"
)

type treeNode struct {
	left     *treeNode
	right    *treeNode
	priority float64
	index    int
}

type PriorityQueue []*treeNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return (*pq[i]).priority < (*pq[j]).priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*treeNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func generateCodes(result *[]string, node *treeNode, curCode string) {
	if node.index > -1 {
		(*result)[node.index] = curCode
		return
	}
	if node.left != nil {
		generateCodes(result, node.left, curCode+"0")
	}
	if node.right != nil {
		generateCodes(result, node.right, curCode+"1")
	}

}

func HuffmanCodes(r io.Reader, w io.Writer) {
	var n int
	_, _ = fmt.Fscan(r, &n)

	var p []float64
	p = make([]float64, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(r, &p[i])
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i, value := range p {
		heap.Push(&pq, &treeNode{
			priority: value,
			index:    i,
			left:     nil,
			right:    nil,
		})
	}

	for pq.Len() > 1 {
		node1 := heap.Pop(&pq).(*treeNode)
		node2 := heap.Pop(&pq).(*treeNode)
		heap.Push(&pq, &treeNode{
			left:     node1,
			right:    node2,
			priority: node1.priority + node2.priority,
			index:    -1,
		})
	}
	var codes = make([]string, n)
	generateCodes(&codes, heap.Pop(&pq).(*treeNode), "")

	for i := 0; i < n; i++ {
		_, _ = fmt.Fprintln(w, codes[i])
	}

}
