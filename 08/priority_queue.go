package main

type ConnectionsPriorityQueue []*Connection

func (pq ConnectionsPriorityQueue) Len() int {
    return len(pq)
}

func (pq ConnectionsPriorityQueue) Less(i, j int) bool {
    return pq[i].dist < pq[j].dist
}

func (pq ConnectionsPriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *ConnectionsPriorityQueue) Push(x any) {
    posDist := x.(*Connection)
    *pq = append(*pq, posDist)
}

func (pq *ConnectionsPriorityQueue) Pop() any {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil
    *pq = old[0: n-1]
    return item
}

