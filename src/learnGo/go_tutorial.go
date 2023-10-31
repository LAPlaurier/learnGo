package main

// 分组导入形式
import (
	"fmt"
	my "learnGo/selfpack"
	"math"
	"time"
)

func testDefer() {
	fmt.Println("in func")

	t := time.Now()
	fmt.Println(t.Date())

	defer fmt.Println(t.Year())
	defer fmt.Println(t.Month())
	defer fmt.Println(t.Day())

	fmt.Println("out func")
}

func add(x int, y int) int {
	return x + y
}

// 函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

// x := 12

func main() {

	// fmt.Println(math.pi) // build Error: undefined
	// go的导出机制：
	// 只有首字母大写的变量、函数、结构体、接口等才是公开的（可以被其他包访问）
	// 其他的都是私有的（只能在包内访问）
	fmt.Println(math.Pi)
	fmt.Printf("%f\n", math.Pi)

	fmt.Println(my.MyFunc())

	fmt.Println("3 + 2 =", add(3, 2))

	fmt.Println(swap("hello", "world"))

}
