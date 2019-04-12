package main
import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func (n *node) setData(data int){
	n.data = data
}

func (n  *node) setNext(newnode *node){
	n.next = newnode
}

func main(){
	a := node{1,nil}
	a.setData(111)
	a1 := node{2,nil}
	a.setNext(&a1)
	fmt.Println(a)
	fmt.Println(a.data)
	fmt.Println(a.next)
	fmt.Println("==========")
	fmt.Println(&a1)
	fmt.Println(&a1.data)
	fmt.Println(&a1.next)
	fmt.Println(&a.next.data)
}
