package main

import (
	"fmt"
	"math"
)

func unims() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,javascript,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)

}

func triangle() {
	var a, b int = 3,4
	fmt.Println(calcTriangle(a,b))
}

func calcTriangle(a,b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func main() {
	unims()
}
