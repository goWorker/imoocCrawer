package basic1
import (
	"fmt"
)
func main() {
	// 准备一个字符串类型
	var description = "This is the pointer example 001"
	// 对字符串取地址, pointer类型为*string
	pointer := &description
	// 打印pointer的类型
	fmt.Printf("pointer type: %T\n", pointer)
	// 打印pointer的指针地址
	fmt.Printf("address: %p\n", pointer)
	// 对指针进行取值操作
	value := *pointer
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)

	var ip *int
	var sp *string

	if ip == nil {
		fmt.Println("ip == nil")
	}
	if sp == nil {
		fmt.Println("sp == nil")
	}
}
