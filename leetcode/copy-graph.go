package main

/**
 * Definition for a Node.
 */

type Node struct {
    Val int
    Neighbors []*Node
 }

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    nodeMapping := make(map[int][]*Node)
    createClones(node, make(map[int]bool), nodeMapping)
    attachNeighbors(node, make(map[int]bool), nodeMapping)

    return nodeMapping[node.Val][1]
}

func createClones(node *Node, visited map[int]bool, nodeMapping map[int][]*Node) {
    visited[node.Val] = true

    clonedNode := &Node{Val: node.Val}
    nodeMapping[node.Val] = []*Node{node, clonedNode}

    for _, neighbor := range node.Neighbors {
        if !visited[neighbor.Val] {
            createClones(neighbor, visited, nodeMapping)
        }
    }
}

func attachNeighbors(node *Node, visited map[int]bool, nodeMapping map[int][]*Node){
    visited[node.Val] = true

    clonedNode := nodeMapping[node.Val][1]
    for _, neighbor := range node.Neighbors {
        clonedNode.Neighbors = append(clonedNode.Neighbors, nodeMapping[neighbor.Val][1])

        if !visited[neighbor.Val] {
            attachNeighbors(neighbor, visited, nodeMapping)
        }
    }
}