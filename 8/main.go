package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ryanjamescaldwell/adventofcode2023/fileReader"
)

type Node struct {
	Key   string
	Left  string
	Right string
}

func (n *Node) String() string {
	return fmt.Sprintf("Node %s: %s, %s", n.Key, n.Left, n.Right)
}

// Network is a struct that holds all the nodes and instructions
// to get from AAA to ZZZ.
type Network struct {
	Schema       map[string]Node
	Instructions []string
}

func getNetwork(lines []string) Network {
	network := Network{}
	network.Schema = make(map[string]Node)
	network.Instructions = strings.Split(lines[0], "")

	for _, line := range lines[2:] {
		key := line[0:3]
		regex := regexp.MustCompile(`([A-Z0-9]{3}), ([A-Z0-9]{3})`)
		tuple := regex.FindString(line)

		node := Node{Key: key}
		node.Left = tuple[0:3]
		node.Right = tuple[5:]
		network.Schema[node.Key] = node
	}
	return network
}

func totalHopsToNavigateNetwork(network Network, startNode Node) int {
	hops := 0
	currentNode := startNode

	for i := 0; i < len(network.Instructions); i++ {
		instruction := network.Instructions[i]

		if instruction == "R" {
			currentNode = network.Schema[currentNode.Right]
		} else {
			currentNode = network.Schema[currentNode.Left]
		}
		hops++

		if strings.HasSuffix(currentNode.Key, "Z") {
			return hops
		}

		// if we still haven't found the node we're looking for,
		// we need to repeat the instructions
		if i+1 >= len(network.Instructions) {
			i = -1 // -1 because the loop will increment it
		}
	}

	return hops
}

func getStartNodes(network Network) []Node {
	startNodes := []Node{}

	for _, node := range network.Schema {
		if strings.HasSuffix(node.Key, "A") {
			startNodes = append(startNodes, node)
		}
	}

	return startNodes
}

func totalHopsToNavigateNetworkConcurrently(network Network, startNodes []Node) int {
	nodeHops := []int{}
	for _, startNode := range startNodes {
		nodeHops = append(nodeHops, totalHopsToNavigateNetwork(network, startNode))
	}

	return LCM(nodeHops[0], nodeHops[1], nodeHops[2:]...)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	lines := fileReader.GetLines()
	network := getNetwork(lines)

	// fmt.Println("Part 1: ", totalHopsToNavigateNetwork(network))

	// Part 2
	startNodes := getStartNodes(network)
	fmt.Println("Part 2: ", totalHopsToNavigateNetworkConcurrently(network, startNodes))
}
