//struct in Go
package main

import (
	"fmt"
)

type employee struct{
	name string
	id uint64
}

func (e *employee) username() string {
	username := "00"+e.name
	return  username
}

func Newemployee(name string, id uint64) *employee{
	emp := employee{
		name: name,
		id: id,
		}
	return &emp
}

func main(){
	emp1 := employee{"Prashanth", 378047}
	fmt.Println(emp1)
	fmt.Println(emp1.id)
	emp2 := employee{
		name: "employee2",
		id: 378089,
		}
	emp2_username := emp2.username()
	fmt.Println(emp2_username)
	fmt.Println(emp2)
	emp3 := employee{}
	fmt.Println(emp3)
	emp4 := Newemployee("employee3", 378901)
	fmt.Println(emp4)
}
