package main

import (
	"fmt"
	"time"
	"github.com/algorithms-golang"
)

func combSort(arr []int) ([]int,int) {

	arrLen := len(arr)
	var count int
	gap := arrLen
	for gap > 1 {
		gap = gap * 10 / 13 //shrink factor is 1.3

		for i := 0; i+gap < arrLen; i++ {
			count++
			if arr[i] > arr[i+gap] {
				tmp := arr[i]
				arr[i] = arr[i+gap]
				arr[i+gap] = tmp
			}
		}
	}

	return arr,count
}



func main() {

	t1:=time.Now()
	arr:=utils.RandArray(10000)
	fmt.Printf("之前的数组：%v\n",arr)
	arr,count:=combSort(arr)
	fmt.Printf("Sort排序后的数组：%v\n",arr)
	fmt.Printf("Sort运行次数：%d\n",count)
	t2:=time.Since(t1)
	fmt.Printf("运算时间：%d",t2)

}
