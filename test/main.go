package main

import "fmt"

//古典问题：有一对兔子，从出生后第3个月起每个月都生一对兔子，
//小兔子长到第三个月后每个月又生一对兔子，
//假如兔子都不死，问每个月的兔子总数为多少？
func getRabbitNum(n int) int {
	rabbitNumMap := map[int]int{
		3: 0,
		2: 0,
		1: 2,
	}
	for i := 0; i < n; i++ {
		rabbitNumMap[3], rabbitNumMap[2], rabbitNumMap[1] =
			rabbitNumMap[2]+rabbitNumMap[3], rabbitNumMap[1], rabbitNumMap[3]/2*2
		//newRabbitNum := rabbitNumMap[3]/2 *2
		//rabbitNumMap[3] = rabbitNumMap[2]+rabbitNumMap[3]
		//rabbitNumMap[2] = rabbitNumMap[1]
		//rabbitNumMap[1] = newRabbitNum
	}
	return rabbitNumMap[1] + rabbitNumMap[2] + rabbitNumMap[3]
}

func main() {
	fmt.Println(getRabbitNum(10))
}
