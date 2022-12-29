package main

import "fmt"

type List struct {
	Val  int
	Next *List
}

func main() {
	head := List{Val: 1, Next: &List{Val: 2, Next: &List{Val: 3, Next: &List{Val: 4, Next: &List{Val: 5}}}}}
	// head1 := List{Val: 1, Next: &List{Val: 2, Next: &List{Val: 3, Next: &List{Val: 4}}}}
	fmt.Println(head)

	deleteGivenNode(head.Next)
	// r1 := iterativeKthToLast(&head1, 3)
	fmt.Println(head.Next)
	// fmt.Println(r1)

}

func deleteGivenNode(node *List) { // Assuming given node is middle
	if node == nil || node.Next == nil {
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
