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

	maps := make(map[int][]*Node)
	buildNodes(node, make(map[int]bool), maps)
	buildNeighbourNodes(node, make(map[int]bool), maps)
	
	return maps[node.Val][1]
}

func buildNodes(node *Node, visited map[int]bool, maps map[int][]*Node) {
	visited[node.Val] = true

	copyNode := Node{
		Val: node.Val, Neighbors: nil,
	}

	maps[node.Val] = []*Node{node, &copyNode}

	for _, p := range node.Neighbors {
		_, exist := visited[p.Val]
		if !exist {
			buildNodes(p, visited, maps)
		}
	}
}

func buildNeighbourNodes(node *Node, visited map[int]bool, maps map[int][]*Node){
	visited[node.Val] = true

	copiedNodePointer := maps[node.Val][1];
	for _, p := range node.Neighbors {
		copiedNodePointer.Neighbors = append(copiedNodePointer.Neighbors, maps[p.Val][1])
		
		_, exist := visited[p.Val]
		if !exist {
			buildNeighbourNodes(p, visited, maps)
		}
	}
}