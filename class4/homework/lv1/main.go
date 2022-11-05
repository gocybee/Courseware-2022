package main

import (
	"fmt"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type Element struct {
	name     string
	priority int
}

func byPriority(a, b interface{}) int {
	priorityA := a.(Element).priority
	priorityB := b.(Element).priority
	return -utils.IntComparator(priorityA, priorityB)
}

func main() {
	vis := map[int]bool{}
	cnt := map[int]int{}
	q := pq.NewWith(utils.IntComparator)

	var n int
	fmt.Scanf("%d\n", &n)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Scanf("%d\n", &x)
		if !vis[x] {
			q.Enqueue(x)
			vis[x] = true
		}
		cnt[x]++
	}
	for !q.Empty() {
		v, _ := q.Dequeue()
		fmt.Printf("%d %d\n", v, cnt[v.(int)])
	}
}
