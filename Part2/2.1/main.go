package main

type List struct {
	Val  int
	Next *List
}

func main() {
	// var head List
	// head.Val = 1
	// head.Next = &List{Val: 3, Next: &List{Val: 1, Next: &List{Val: 1}}}
	// res := deleteDuplicateWithReturn(&head)
	// fmt.Println("\nres:", res)

}

func deleteDuplicateWithReturn(head *List) *List {
	curr := head
	m := make(map[int]int)
	var prev *List

	for curr != nil {
		if _, ok := m[curr.Val]; ok {
			prev.Next = curr.Next
		} else {
			m[head.Val]++
			prev = curr
		}
		curr = curr.Next
	}
	return head
}

func deleteDuplicateNoReturn(head *List) {

	m := make(map[int]int)
	var prev *List

	for head != nil {
		if _, ok := m[head.Val]; ok {
			prev.Next = head.Next
		} else {
			m[head.Val]++
			prev = head
		}
		head = head.Next
	}
}
