package main

import (
	"fmt"
	"time"
	"github.com/algorithms-golang"
)

func findMin(arr []int) ([]int,int) {
	length:=len(arr)
	index:=0
	min:=arr[index]
	for i:=1;i<length;i++{
		if arr[i]<min {
			min=arr[i]
		}
	}
	arr=append(arr[:index],arr[index+1:]...)
	return arr,min
}

func selectSort(arr []int) []int {
	length:=len(arr)
	newArr:=make([]int,length)
	for i:=0;i<length;i++{
		arr,newArr[i]=findMin(arr)
		}
	return newArr
}


func main() {

	t1:=time.Now()
	arr:=utils.RandArray(10)
	fmt.Printf("之前的数组：%v\n",arr)
	arr=selectSort(arr)
	fmt.Printf("Sort排序后的数组：%v\n",arr)
	t2:=time.Since(t1)
	fmt.Printf("运算时间：%d",t2)

}