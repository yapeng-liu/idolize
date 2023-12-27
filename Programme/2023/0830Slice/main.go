package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	testSlices := make([]int, 10)
	testSlices[0] = 1
	testSlices[2] = 2
	fmt.Println(testSlices)
	var toAccount = []int64{1, 2, 3, 4, 5}
	//使用批量单聊
	size := 3 //官方每次批量发送的数量限制为500，分多次发送
	total := len(toAccount)
	//动态适应用户数量创建切片数量
	slicesNum := (total / size) + 1
	slices := make([][]string, slicesNum)
	currentIndex := 0
	for i := 0; i < slicesNum; i++ {
		creatSize := size
		//如果当前是最后一个切片，并且有剩余数量
		if i == slicesNum-1 && total%size != 0 {
			creatSize = total - currentIndex
		}
		slices[i] = make([]string, creatSize)
		for j := 0; j < creatSize; j++ {
			slices[i][j] = strconv.FormatInt(toAccount[currentIndex], 10)
			currentIndex++
		}
		func() {
			fmt.Println("协程", slices[i])
		}()
	}
	fmt.Println(slices)
	fmt.Printf("%v 数据类型 %T\n", slices[0], slices[0])
	time.Sleep(5 * time.Second)
	fmt.Println("123 ", time.Now().Format("2006-01-02 15:04:05"))
}
