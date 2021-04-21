package main

import "fmt"

func main() {
	result := myfunc()
	fmt.Println("result = ", result)
}
func myfunc(args ...int) (result int) {
	//建议使用(result int)方式     args可以传多个，也可以传0个
	for _, data := range args {
		//i是args的下标，data是args的元素
		fmt.Println(data)
		result += data
		//	等价于result = result+data
	}
	return
	//	默认返回result
}
