package main
import "fmt"
import "math"

type Point struct {
	x, y float64
}

func (p *Point) AddX(val float64) {
	p.x = p.x + val
}

func (p *Point) DistToOrigin () (float64){
	res := p.x*p.x + p.y*p.y
	return math.Sqrt(res)
}

func main() {
	var p = &Point{3.6, 4.8}
	p.AddX(1.2)
	fmt.Println(p) // {4.8 4.8}
	fmt.Println(p.DistToOrigin())
}