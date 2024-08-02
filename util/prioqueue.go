package util

import "container/heap"

type PrioQueueItem[T any] struct {
	Value    T
	Index    int
	Priority int
}

type PrioQueue[T any] []*PrioQueueItem[T]

func (pq PrioQueue[T]) Len() int { return len(pq) }

func (pq PrioQueue[T]) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PrioQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PrioQueue[T]) Push(x any) {
	todo := x.(*PrioQueueItem[T])
	*pq = append(*pq, todo)
}

func (pq *PrioQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func (pq *PrioQueue[T]) update(item *PrioQueueItem[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
