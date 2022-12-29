package main

import "fmt"

type List struct {
	Val  int
	Next *List
}

func main() {
	head := List{Val: 1, Next: &List{Val: 2, Next: &List{Val: 3, Next: &List{Val: 4}}}}
	head1 := List{Val: 1, Next: &List{Val: 2, Next: &List{Val: 3, Next: &List{Val: 4}}}}
	fmt.Println(head)

	r := returnKthToLast(&head, 3)
	r1 := iterativeKthToLast(&head1, 3)
	fmt.Println(r)
	fmt.Println(r1)

}

func iterativeKthToLast(head *List, k int) *List {
	p1, p2 := head, head

	for i := 0; i < k; i++ {
		if p1 == nil {
			return nil
		}
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

func returnKthToLast(head *List, k int) *List {
	curr := head
	n := 0

	for curr != nil {
		n++
		curr = curr.Next
	}

	curr = head

	for i := 0; i < (n - k); i++ {
		curr = curr.Next
	}

	return curr
}
