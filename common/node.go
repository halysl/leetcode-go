package common

type Node struct {
	Val int
	Next *Node
}

type NodeList struct {
	size uint64 // 车辆数量
	// head *Node  // 车头
	tail *Node  // 车尾
}

func (nodeList *NodeList) Init() {
	(*nodeList).size = 0
	(*nodeList).tail = nil
}

func (nodeList *NodeList) Append (node *Node) {
	nodeList.tail = node
}


