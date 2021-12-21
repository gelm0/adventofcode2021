package dumbooctopus

type NeighbourNode struct {
	Value int
	// Neighbours contains coordinates to neighbours in the array
	Neighbours []int
	Flashed    bool
}

func InitializeNodes(points []int, height int) []NeighbourNode {
	nodes := make([]NeighbourNode, len(points))
	length := len(points) / height
	var neighbours []int
	for i, val := range points {
		row := i / length
		col := i % length
		if row == height-1 {
			if col == 0 {
				neighbours = []int{i + 1, i - length, i - length + 1}
			} else if col == length-1 {
				neighbours = []int{i - 1, i - length, i - length - 1}
			} else {
				neighbours = []int{i - 1, i + 1, i - length, i - length + 1, i - length - 1}
			}
		} else if row == 0 {
			if col == 0 {
				neighbours = []int{i + 1, i + length, i + length + 1}
			} else if col == length-1 {
				neighbours = []int{i - 1, i + length, i + length - 1}
			} else {
				neighbours = []int{i - 1, i + 1, i + length, i + length + 1, i + length - 1}
			}
		} else if col == 0 {
			neighbours = []int{i + 1, i + length, i - length, i + length + 1, i - length + 1}
		} else if col == length-1 {
			neighbours = []int{i - 1, i + length, i - length, i + length - 1, i - length - 1}
		} else {
			neighbours = []int{i - 1, i + 1, i + length, i + length + 1, i + length - 1, i - length, i - length + 1, i - length - 1}
		}
		nodes[i] = NeighbourNode{
			Value:      val,
			Neighbours: neighbours,
		}
	}
	return nodes
}

func resetNodes(nodes []NeighbourNode) {
	for i, _ := range nodes {
		nodes[i].Flashed = false
	}
}

func flashNodes(nodes []NeighbourNode, flashedNodes []int, flashes int) int {
	if len(flashedNodes) == 0 {
		return flashes
	}
	var neighBourList []int
	for _, index := range flashedNodes {
		if nodes[index].Value >= 9 && !nodes[index].Flashed {
			// Nodes which have flashed need to have their neighbours updated
			neighBourList = append(neighBourList, nodes[index].Neighbours...)
			nodes[index].Value = 0
			nodes[index].Flashed = true
			flashes += 1
		}
		flashedNodes = flashedNodes[1:]
	}
	// Update neighbours to nodes who have flashed
	for _, nnIndex := range neighBourList {
		// Avoid nodes who have flashed or should flash next round
		if nodes[nnIndex].Value != 0 && nodes[nnIndex].Value < 9 {
			nodes[nnIndex].Value += 1
		} else {
			// Nodes to still be flashed if node.Flashed != true
			flashedNodes = append(flashedNodes, nnIndex)
		}
	}
	// Recursively update til all nodes are flashed or we have no neighbours left
	return flashNodes(nodes, flashedNodes, flashes)
}

func Step(steps int, nodes []NeighbourNode) (flashes int) {
	var flashedNodes []int
	for step := 0; step < steps; step++ {
		for i, _ := range nodes {
			if nodes[i].Value >= 9 {
				// Start round by adding the node plus it's neighbours
				flashedNodes = append(flashedNodes, i)
				flashedNodes = append(flashedNodes, nodes[i].Neighbours...)
			} else {
				// Otherwise just update each value if not flashed
				nodes[i].Value += 1
			}
		}
		// Reset nodes
		flashes += flashNodes(nodes, flashedNodes, 0)
		flashedNodes = []int{}
		resetNodes(nodes)
	}
	return
}

func synchronized(nodes []NeighbourNode) bool {
	for _, n := range nodes {
		if n.Value != 0 {
			return false
		}
	}
	return true
}

func FindFirstSynchronized(nodes []NeighbourNode) int {
	stepsTaken := 0
	for {
		Step(1, nodes)
		stepsTaken += 1
		if synchronized(nodes) {
			break
		}
	}
	return stepsTaken
}
