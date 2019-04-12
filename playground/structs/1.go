package main
import ("fmt")

//Linked List Node Structure
type node struct {
	data int
	next *node
}

func NewNode(currentnode *node,data int,next *node) node {
	newnode := node{data, next}
	fmt.Println(currentnode)
	currentnode.next = &newnode
	return newnode
}

func main(){
	n := node{1, nil}
	n1 := node {2, nil}
	n.next = &n1
	fmt.Println(n, n1)
	nn := NewNode(&n1, 11, nil)
	fmt.Println(n, n1, nn)
	fmt.Println(&n.data, &n1.data, &nn.data)
}
