package basic1

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("Unsupport operation: %v", op)

	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//第一个参数是一个有两个参数返回一个int的函数，后面跟两个int参数，最终返回一个int
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d,%d)\n", opName, a, b)
	return op(a, b)
}

func incr(p *int) int {
	*p++ //p是一个整型指针，*p代表指针所指向的值
	fmt.Printf("value type: %T\n", *p)
	return *p
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, _ := div(12, 5)
	fmt.Println(q)
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	v := 1
	incr(&v)              // &v 是1所在的地址
	fmt.Println(&v)       //打印2的地址，结果为 0xc420084008
	fmt.Println(incr(&v)) //现在v的值为3

}
