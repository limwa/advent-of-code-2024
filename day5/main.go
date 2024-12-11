package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/limwa/advent-of-code-2024/lib/cast"
	"github.com/limwa/advent-of-code-2024/lib/util"
)

type Node struct {
	value int
	after []*Node
	enabled bool
	visitDepth int
}

type Graph struct {
	nodes []*Node
	values_to_node map[int]int
}

func createGraph() Graph {
	return Graph{
		nodes: []*Node{},
		values_to_node: map[int]int{},
	}
}

func (g *Graph) getNode(value int) (bool, *Node) {
	if index, ok := g.values_to_node[value]; ok {
		return true, g.nodes[index]
	}

	return false, nil
}

func (g *Graph) getNodeOrCreate(value int) *Node {
	if found, node := g.getNode(value); found {
		return node
	}

	index := len(g.nodes)
	g.values_to_node[value] = index

	node_ptr := &Node{
		value: value,
		after: []*Node{},
		enabled: true,
		visitDepth: -1,
	}

	g.nodes = append(g.nodes, node_ptr)
	return node_ptr
}

func (g *Graph) addDependency(before int, after int) {
	before_node := g.getNodeOrCreate(before)
	after_node := g.getNodeOrCreate(after)

	before_node.after = append(before_node.after, after_node)
}

func (g *Graph) dfs(value int, maxDepth int) {
	g._dfs(value, 0, maxDepth)
}

func (g *Graph) _dfs(value int, currentDepth int, maxDepth int) {
	if currentDepth == maxDepth { 
		return
	}

	_, node := g.getNode(value)
	node.visitDepth = currentDepth

	for _, after := range node.after {
		if !after.enabled || (after.is_visited() && after.visitDepth > currentDepth) {
			continue
		}

		g._dfs(after.value, currentDepth + 1, maxDepth)
	}
}

func (g *Graph) resetVisits() {
	for _, node := range g.nodes {
		node.visitDepth = -1
	}
}

func (g *Graph) resetMask() {
	for _, node := range g.nodes {
		node.enabled = true
	}
}

func (g *Graph) mask(values []int) {
	for _, node := range g.nodes {
		node.enabled = false
	}

	for _, value := range values {
		_, node := g.getNode(value)
		node.enabled = true
	}
}

func (g *Graph) isVisited(value int) bool {
	_, node := g.getNode(value);
	return node.is_visited()
}

func (n *Node) is_visited() bool {
	return n.visitDepth != -1
}

type Rule struct {
	before int
	after int
}

func parseInput(input string) ([]Rule, [][]int) {
	input_parts := strings.Split(input, "\n\n")
	rules_part, updates_part := input_parts[0], input_parts[1]

	rules := []Rule{}
	for _, rule := range strings.Split(rules_part, "\n") {
		rule_parts := strings.Split(rule, "|")
		rules = append(rules, Rule{
			before: cast.ToInt(rule_parts[0]),
			after: cast.ToInt(rule_parts[1]),
		})
	}

	updates := [][]int{}
	for _, update := range strings.Split(updates_part, "\n") {
		pages := []int{}
		for _, page := range strings.Split(update, ",") {
			pages = append(pages, cast.ToInt(page))
		}

		updates = append(updates, pages)
	}

	return rules, updates
}

func createGraphFromRules(rules []Rule) Graph {
	graph := createGraph()
	for _, rule := range rules {
		graph.addDependency(rule.before, rule.after)
	}
	return graph
}

func isOrdered(graph *Graph, update []int) bool {
	graph.resetVisits()
	graph.resetMask()

	for i := len(update) - 1; i >= 0; i-- {
		value := update[i]
	
		if graph.isVisited(value) {
			return false
		}

		// Starting node and children
		graph.dfs(value, 2)
	}

	return true
}

func solvePart1(input string) string {
	rules, updates := parseInput(input)
	graph := createGraphFromRules(rules)
	
	sum := 0
	for _, update := range updates {
		if isOrdered(&graph, update) {
			mid := update[len(update) / 2]
			sum += mid
		}
	}
	
	return cast.ToString(sum)
}

func solvePart2(input string) string {
	rules, updates := parseInput(input)
	graph := createGraphFromRules(rules)

	sum := 0
	for _, update := range updates {
		if !isOrdered(&graph, update) {
			graph.resetVisits()

			graph.resetMask()
			graph.mask(update)

			lastDfs := -1
			for _, value := range update {
				if graph.isVisited(value) {
					continue
				}
				
				lastDfs = value
				graph.dfs(value, -1)
			}

			graph.resetVisits()
			graph.dfs(lastDfs, -1)

			middleDepth := len(update) / 2
			for _, value := range update {
				_, node := graph.getNode(value)
				if node.visitDepth == middleDepth {
					sum += value
					break
				}
			}
		}
	}

	return cast.ToString(sum)
}

// START template

//go:embed input.txt
var _input string

func init() {
	util.NormalizeInput(&_input)
}

func measure(execute func() string) string {
    start := time.Now()
	answer := execute()
	fmt.Printf("Execution time: %v\n", time.Since(start))
	return answer
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part to solve")
	flag.Parse()

	answer := measure(func() string {
		if part == 1 {
			return solvePart1(_input)
		} else if part == 2 {
			return solvePart2(_input)
		} else {
			panic("a valid part must be specified")
		}
	})
	
	util.CopyToClipboard(answer)
	fmt.Printf("Answer for part %d:\n%s\n", part, answer)
}

// END template
