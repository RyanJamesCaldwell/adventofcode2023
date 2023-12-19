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
		regex := regexp.MustCompile(`([A-Z]{3}), ([A-Z]{3})`)
		tuple := regex.FindString(line)

		node := Node{Key: key}
		node.Left = tuple[0:3]
		node.Right = tuple[5:]
		network.Schema[node.Key] = node
	}
	return network
}

func main() {
	lines := fileReader.GetLines()
	network := getNetwork(lines)

	for _, node := range network.Schema {
		fmt.Println(node.String())
	}
}
