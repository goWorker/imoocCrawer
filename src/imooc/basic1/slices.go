package basic1

import "fmt"
func Updateslice(s []int){
	s[0] = 100
}
func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	//s := arr[2:6]
	fmt.Println("arr[2:6] = ",arr[2:6])
	s1 := arr[2:]
	fmt.Println("S1 = ",s1)
	s2 := arr[:]
	fmt.Println("S2 = ",s2)
	fmt.Println("After updateslice(s1)",)
	Updateslice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateslice(s2)",)
	Updateslice(s2)
	fmt.Println(s2)
	fmt.Println(arr)
}

