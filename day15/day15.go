package day15

import (
	"math"
	"strconv"
	"strings"
)

// Let the node at which we are starting at be called the initial node.
// Let the distance of node Y be the distance from the initial node to Y.
// Dijkstra's algorithm will initially start with infinite distances and will try to improve them step by step.

// Mark all nodes unvisited. Create a set of all the unvisited nodes called the unvisited set.
// Assign to every node a tentative distance value: set it to zero for our initial node and to infinity for all other nodes.
// The tentative distance of a node v is the length of the shortest path discovered so far between the node v and the starting node.
// Since initially no path is known to any other vertex than the source itself (which is a path of length zero), all other tentative distances are initially set to infinity. Set the initial node as current.[15]
// For the current node, consider all of its unvisited neighbors and calculate their tentative distances through the current node.
// Compare the newly calculated tentative distance to the current assigned value and assign the smaller one.
// For example, if the current node A is marked with a distance of 6, and the edge connecting it with a neighbor B has length 2,
// then the distance to B through A will be 6 + 2 = 8. If B was previously marked with a distance greater than 8 then change it to 8.
// Otherwise, the current value will be kept.
// When we are done considering all of the unvisited neighbors of the current node, mark the current node as visited and remove it from the unvisited set. A visited node will never be checked again.
// If the destination node has been marked visited (when planning a route between two specific nodes) or
// if the smallest tentative distance among the nodes in the unvisited set is infinity (when planning a complete traversal; occurs when there is no connection between the initial node and remaining unvisited nodes), then stop.
// The algorithm has finished.
// Otherwise, select the unvisited node that is marked with the smallest tentative distance, set it as the new current node, and go back to step 3.
// When planning a route, it is actually not necessary to wait until the destination node is "visited" as above:
// the algorithm can stop once the destination node has the smallest tentative distance among all "unvisited" nodes (and thus could be selected as the next "current").

type Node struct {
	x                   int
	y                   int
	visited             bool
	difficultyFromStart int
	difficulty          int
}

type Nodes []*Node

var CURRENT_NODE *Node

var TARGET_NODE *Node

func parseInput(input []string) Nodes {
	nodes := Nodes{}
	for x, line := range input {
		//38438
		slc := strings.Split(line, "")
		for y, strNum := range slc {
			num, _ := strconv.Atoi(strNum)
			node := Node{
				x:                   x,
				y:                   y,
				visited:             false,
				difficultyFromStart: math.MaxInt,
				difficulty:          num,
			}
			if x == 0 && y == 0 {
				node.difficultyFromStart = 0
				CURRENT_NODE = &node
			} else {
				nodes = append(nodes, &node)
			}
			if x == len(input)-1 && y == len(slc)-1 {
				TARGET_NODE = &node
			}

		}
	}
	return nodes
}

func findNodeWithSmallestTentativeDistance(nodes Nodes) *Node {
	smallest := Node{
		difficultyFromStart: math.MaxInt,
	}
	foundIdx := 0
	for idx, n := range nodes {
		if n.difficultyFromStart < smallest.difficultyFromStart {
			smallest = *n
			foundIdx = idx
		}

	}

	return nodes[foundIdx]
}

func compareNeighbour(nodeA *Node, nodeB *Node) {
	difficultyFromStartToBThroughA := nodeA.difficultyFromStart + nodeB.difficulty
	if nodeB.difficultyFromStart >= difficultyFromStartToBThroughA {
		nodeB.difficultyFromStart = difficultyFromStartToBThroughA
	}
}

func getNeighbours(node *Node, nodes Nodes) Nodes {
	neighbours := Nodes{}

	for _, n := range nodes {
		if (n.x == node.x+1 || n.x == node.x-1) && n.y == node.y {
			neighbours = append(neighbours, n)
		} else if (n.y == node.y+1 || n.y == node.y-1) && n.x == node.x {
			neighbours = append(neighbours, n)
		}
	}
	return neighbours
}

func compareAllNeighbours(node *Node, nodes Nodes) {
	neighbours := getNeighbours(node, nodes)
	for _, n := range neighbours {
		compareNeighbour(node, n)
	}
}

func removeNode(node Node, nodes Nodes) Nodes {
	for idx, n := range nodes {
		if node.x == n.x && node.y == n.y {
			nodes[idx] = nodes[len(nodes)-1]
			return nodes[:len(nodes)-1]
		}
	}
	return nodes
}

func djikstra(nodes Nodes) int {
	for {
		neighbours := getNeighbours(CURRENT_NODE, nodes)
		compareAllNeighbours(CURRENT_NODE, neighbours)
		CURRENT_NODE.visited = true
		removeNode(*CURRENT_NODE, nodes)
		CURRENT_NODE = findNodeWithSmallestTentativeDistance(nodes)
		if TARGET_NODE.visited {
			return TARGET_NODE.difficultyFromStart
		}
	}
}

func Task1(input []string) int {
	nodes := parseInput(input)
	return djikstra(nodes)
}

func Task2() {

}
