package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	index int
	next  *node
	prev  *node
}

type circle struct {
	head  *node
	nodes []*node
}

type nodeSlice struct {
	nodes []*node
	head  *node
	tail  *node
}

func newSlice(ns []*node) *nodeSlice {
	return &nodeSlice{ns, ns[len(ns)-1], ns[0]}
}

func jump(n *node, N int) *node {
	var curr = n

	for i := 0; i < N; i++ {
		curr = curr.next
	}
	return curr
}

func pinch(n *node, N int) (*nodeSlice, *nodeSlice) {
	var slice1 []*node
	var slice2 []*node
	var curr = n

	for i := 0; i < N; i++ {
		slice1 = append(slice1, curr)
		curr = curr.next
	}
	for i := N; ; i++ {
		slice2 = append(slice2, curr)

		curr = curr.next
		if curr == n {
			break
		}
	}

	var cut = newSlice(slice1)
	var leftover = newSlice(slice2)

	return cut, leftover
}

func twist(cut *nodeSlice) {
	var tmp *node

	curr := cut.tail
	for {
		tmp = curr.next
		curr.next = curr.prev
		curr.prev = tmp

		curr = tmp
		if curr == cut.head {
			break
		}
	}

	tmp = cut.head
	cut.head = cut.tail
	cut.tail = tmp
}

func join(ns1 *nodeSlice, ns2 *nodeSlice) {
	ns1.head.next = ns2.tail
	ns2.head.next = ns1.tail

	ns1.tail.prev = ns2.head
	ns2.tail.prev = ns1.head
}

func getNodeIndex(ns *nodeSlice, n *node) int {
	for i, ni := range ns.nodes {
		if ni == n {
			return i
		}
	}
	return -1
}

func pushHead(ns *nodeSlice, N int) *node {
	var head = ns.tail

	for i := 0; i < N-1; i++ {
		head = head.next
	}
	return head
}

func printCycle(c circle) {
	var cur = c.head
	for {
		fmt.Printf("%d <- ", cur.prev.index)
		fmt.Printf("%d -> ", cur.index)
		fmt.Printf("%d \n", cur.next.index)

		cur = cur.next
		if cur == c.head {
			break
		}
	}
}

func printNodeSlice(c *nodeSlice) {
	var cur = c.tail
	for {

		fmt.Printf("%d <- ", cur.prev.index)
		fmt.Printf("%d -> ", cur.index)
		fmt.Printf("%d \n", cur.next.index)

		cur = cur.next
		if cur == c.head {
			break
		}
	}
}

func printCuts(s1 *nodeSlice, s2 *nodeSlice) {
	fmt.Println("Cut")
	fmt.Printf("TAIL %d <- %d -> %d\n", s1.tail.prev.index, s1.tail.index, s1.tail.next.index)
	fmt.Printf("HEAD %d <- %d -> %d\n", s1.head.prev.index, s1.head.index, s1.head.next.index)
	fmt.Println()
	fmt.Println("Leftover")
	fmt.Printf("TAIL %d <- %d -> %d\n", s2.tail.prev.index, s2.tail.index, s2.tail.next.index)
	fmt.Printf("HEAD %d <- %d -> %d\n", s2.head.prev.index, s2.head.index, s2.head.next.index)
}

func main() {
	var inputStr = "46,41,212,83,1,255,157,65,139,52,39,254,2,86,0,204"
	var cycleLen = 25

	var input []int
	for _, r := range strings.Split(inputStr, ",") {
		var i, _ = strconv.Atoi(r)
		input = append(input, i)
	}

	var nodes []*node
	for i := 0; i < cycleLen; i++ {
		if i == 0 {
			nodes = append(nodes, &node{i, nil, nil})
			continue
		}

		nodes = append(nodes, &node{i, nil, nodes[i-1]})
	}

	for i := 0; i < cycleLen; i++ {
		if i == 0 {
			nodes[i].prev = nodes[cycleLen-1]
		}

		var k = (i + 1) % cycleLen
		nodes[i].next = nodes[k]
	}

	var cycle = circle{nodes[0], nodes}
	var s1, s2 = pinch(cycle.head, 9)

	twist(s1)

	_ = s2
	//printCuts(s1, s2)
	printNodeSlice(s1)

	//join(s1, s2)

	var indexO = getNodeIndex(s1, cycle.head)
	//fmt.Println(indexO)
	if indexO != -1 {
		var d = (len(s1.nodes) / 2) - indexO
		//fmt.Println(d)
		indexO = (len(s1.nodes) / 2) + d
		//fmt.Println(indexO)
		//cycle.head = pushHead(s1, indexO)
	}

}
