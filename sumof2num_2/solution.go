package sumof2num_2

type ListNode struct {
     Val int
     Next *ListNode
}
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := ListNode{0, nil}
	s := &ln
	//_flag := false
	i,j := l1, l2
	for  {
		if i != nil && j != nil {
			ln.Val = i.Val + j.Val
			ln.Next = &ListNode{0, nil}
			ln = *ln.Next
		} else if i != nil {

		} else if j != nil {

		} else {

		}
		i, j=i.Next, j.Next
	}
return s
}

func Run() {

}


