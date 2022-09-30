package node

func Priority(node *Node) int {
	return int(node.CpuUsage*100) % 10000
}
