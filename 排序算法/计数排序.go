package main

import (
	"fmt"
	"time"
	"github.com/algorithms-golang"
)

func getK(arr []int) int {
	if len(arr) == 0 {
		return 1
	}

	k := arr[0]

	for _, v := range arr {
		if v > k {
			k = v
		}
	}

	return k+1
}

func countingSort(arr []int) ([]int,int) {

	k := getK(arr)
	arrayOfCounts := make([]int, k)

	count:=int(0)
	ch:=make(chan int,1)

	ch<-count
	for i:= 0; i < len(arr); i++ {
		count=<- ch
		count++
		ch<-count
		arrayOfCounts[arr[i]] += 1
	}

	for i, j := 0, 0; i < k; i++ {
		for {
			count=<- ch
			count++
			ch<-count
			if arrayOfCounts[i] > 0 {
				arr[j] = i
				j += 1
				arrayOfCounts[i] -= 1
				continue
			}
			break
		}
	}
	return arr,count
}


func main() {

	t1:=time.Now()
	arr:=utils.RandArray(10)
	fmt.Printf("之前的数组：%v\n",arr)
	arr,count:=countingSort(arr)
	fmt.Printf("Sort排序后的数组：%v\n",arr)
	fmt.Printf("Sort运行次数：%d\n",count)
	t2:=time.Since(t1)
	fmt.Printf("运算时间：%d",t2)

}
