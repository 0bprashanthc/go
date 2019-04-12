package main
import ("fmt")

type point struct{
	x int
	y int
}

type square struct{
	center point
	length int
}

func (p *point) move(dx int, dy int){
	p.x = dx
	p.y = dy
}

func (s *square) area() (int, error){
	return s.center.x*s.center.y, nil
}

func NewSquare(x int, y int, length int) (*square){
	np := point{x,y}
	ns := square{np, length}
	return &ns
}

func main(){
	p := point{10,10}
	s := square{p,100}
	p.move(12,12)
	fmt.Println(s.area())
	ns := NewSquare(1,1,10)
	fmt.Println(ns)
	fmt.Println(&ns.center.x, &ns.center.y)
}
