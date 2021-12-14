package day12

import (
	"fmt"
	"strings"
)

type Node struct {
	name         string
	connections  []string
	big          bool
	canBeVisited bool
}

var NodeMap = make(map[string]Node)

type NodeList struct {
	nodes     []string
	completed bool
	valid     bool
}

var Paths []NodeList

// type Connector interface {
// 	findConnections()
// }

// func (nl NodeList) findConnections() {

// }

func isLowerCase(str string) bool {
	//checks if a string is equivalent to its lowercase version
	return str == strings.ToLower(str)
}

func parseInput(input []string) {
	for _, line := range input {
		split := strings.Split(line, "-")
		_, exists := NodeMap[split[0]]
		if !exists {
			NodeMap[split[0]] = Node{
				name:         split[0],
				connections:  []string{},
				canBeVisited: true,
				big:          !isLowerCase(split[0]),
			}
		}

		_, exists2 := NodeMap[split[1]]
		if !exists2 {
			NodeMap[split[1]] = Node{
				name:         split[1],
				connections:  []string{},
				canBeVisited: true,
				big:          !isLowerCase(split[1]),
			}
		}
		temp := NodeMap[split[0]]
		temp.connections = append(temp.connections, NodeMap[split[1]].name)
		NodeMap[split[0]] = temp
		temp = NodeMap[split[1]]
		temp.connections = append(temp.connections, NodeMap[split[0]].name)
		NodeMap[split[1]] = temp
	}
}

func start() {
	for _, connection := range NodeMap["start"].connections {
		nl := NodeList{
			nodes:     []string{"start", connection},
			completed: false,
			valid:     true,
		}
		Paths = append(Paths, nl)
	}
}

func allPathsComplete() bool {
	for _, nl := range Paths {
		if nl.completed == false {
			return false
		}
	}
	return true
}

func containsRepeat(slc []string) bool {
	for i := 0; i < len(slc); i++ {
		for j := 0; j < len(slc); j++ {
			if j == i {
				continue
			}
			if slc[i] == slc[j] {
				return true
			}
		}
	}
	return false
}

func isPathValid(connections []string) bool {
	lowercase := []string{}
	for _, str := range connections {
		if isLowerCase(str) {
			lowercase = append(lowercase, str)
		}
	}
	return !containsRepeat(lowercase)
}

func formNextConnection() {
	temp := []NodeList{}
	for i := 0; i < len(Paths); i++ {
		lastConn := Paths[i].nodes[len(Paths[i].nodes)-1]
		if lastConn == "end" {
			Paths[i].completed = true
		}
		if !isPathValid(Paths[i].nodes) {
			Paths[i].valid = false
		}
		if !Paths[i].valid || Paths[i].completed {
			temp = append(temp, Paths[i])
		} else {
			possibleConnections := NodeMap[lastConn].connections
			for _, conn := range possibleConnections {
				new := NodeList{
					completed: false,
					valid:     true,
					nodes:     NodeMap[conn].connections,
				}
				temp = append(temp, new)
			}

		}
	}
	Paths = temp
}

func generatePaths() {
	start()

	// arePathsCompleted := allPathsComplete()
	// for !arePathsCompleted {
	// 	formNextConnection()
	// 	arePathsCompleted = allPathsComplete()
	// }
	fmt.Println(Paths)
	formNextConnection()
	fmt.Println(Paths)
	formNextConnection()
	fmt.Println(Paths)
	formNextConnection()
	fmt.Println(Paths)
}

func Task1(input []string) int {
	parseInput(input)
	generatePaths()

	return 0
}
