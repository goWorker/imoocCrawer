package basic1

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,9}
	var grid [4][5]int
	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	for i,v := range arr3 {
		fmt.Println(i, v)
	}
}