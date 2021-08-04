package deletenownode_m0203

import "fmt"
import "leetcode-go/common"



func deleteNode(node *common.Node) {

}

func Run() {

	var (
		ln0  = ListNode{0, nil}
		ln1  = ListNode{1, nil}
		ln2  = ListNode{2, nil}
		ln3  = ListNode{3, nil}
		ln4  = ListNode{4, nil}
		ln5  = ListNode{5, nil}
		ln6  = ListNode{6, nil}
		ln7  = ListNode{7, nil}
		)
	ln0.Next = &ln1
	ln1.Next = &ln2
	ln2.Next = &ln3
	ln3.Next = &ln4
	ln4.Next = &ln5
	ln5.Next = &ln6
	ln6.Next = &ln7

	deleteNode(&ln3)


	fmt.Println(ln0)
}