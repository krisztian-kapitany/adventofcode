package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type node struct {
	name           string
	leaf           bool
	finished       bool
	weight         int
	originalWeight int
	leaves         []*node
	leafStr        string
}

type nTree struct {
	root *node
}

func newNode(line string) *node {
	w, _ := strconv.Atoi(line[strings.Index(line, "(")+1 : strings.Index(line, ")")])

	var lstr string

	if strings.Contains(line, " -> ") {
		lstr = line[strings.Index(line, "> ")+2:]
	}

	lstr = line[strings.Index(line, "> ")+2:]

	return &node{
		name:           line[:strings.Index(line, " (")],
		leaf:           !(strings.Contains(line, " -> ")),
		finished:       !(strings.Contains(line, " -> ")),
		weight:         w,
		originalWeight: w,
		leafStr:        lstr}
}

func setW(n *node) int {
	if n.leaf || n.finished {
		return n.weight
	}

	sum := n.weight
	for _, r := range n.leaves {
		sum += setW(r)
	}

	n.weight = sum
	n.finished = true

	return sum
}

func getBalanceMap(n *node) map[int]string {

	m := make(map[int]string)

	for _, r := range n.leaves {
		if _, ok := m[r.weight]; ok {
			m[r.weight] = "__DUPLICATE__"
			continue
		}

		m[r.weight] = r.name
	}

	return m
}

func walkTree(t *nTree) (*node, *node) {

	var current = t.root
	var parent = t.root

	for {
		if current.leaves == nil {
			break
		}

		var m = getBalanceMap(current)
		var next *node

		for _, name := range m {
			if name != "__DUPLICATE__" {
				next = getNodeByName(current.leaves, name)
				break
			}
		}

		if next == nil {
			return current, parent
		}

		parent = current
		current = next
	}

	return nil, nil

}

func getNodeByName(list []*node, name string) *node {
	for _, r := range list {
		if r.name == name {
			return r
		}
	}
	return nil
}

func setLeaves(n *node, nodeList []*node) bool {
	if n == nil {
		return false
	}

	for _, r := range strings.Split(n.leafStr, ", ") {

		nn := getNodeByName(nodeList, r)
		n.leaves = append(n.leaves, nn)
	}

	for _, il := range n.leaves {
		setLeaves(il, nodeList)
	}

	return true
}

func main() {
	var lines, _ = readLines("input7.txt")

	clues := make(map[string]string)

	for _, line := range lines {
		if strings.Contains(line, " -> ") {
			clues[line[:strings.Index(line, " (")]] = line[strings.Index(line, "> ")+2:]
		}
		continue
	}

	var invalidRoots []string
	for root := range clues {
		for _, leaves := range clues {
			if strings.Contains(leaves, root) {
				invalidRoots = append(invalidRoots, root)
			}
		}
	}

	for _, root := range invalidRoots {
		delete(clues, root)
	}

	fmt.Println("The root node is:")
	fmt.Println(clues)
	fmt.Println("")

	var nodelist []*node
	var tree nTree

	for _, line := range lines {

		n := newNode(line)
		nodelist = append(nodelist, n)

		if n.name == "mwzaxaj" {
			tree = nTree{
				root: n}
		}
	}

	setLeaves(tree.root, nodelist)
	setW(tree.root)

	var lastNode, lastParent = walkTree(&tree)
	var equalNode *node

	for _, r := range lastParent.leaves {
		if r.name != lastNode.name {
			equalNode = r
			break
		}
	}
	var diff = equalNode.weight - lastNode.weight

	fmt.Printf("Last nodes weight should be: %d \n", lastNode.originalWeight+diff)

}
