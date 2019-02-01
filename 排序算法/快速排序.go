package main

import (
	"time"
	"fmt"
	"github.com/algorithms-golang"
)

//交换元素位置
func swap(arr []int, i, k int){
	tem:=arr[i]
	arr[i]=arr[k]
	arr[k]=tem
}

func partition(arr []int,left,right int) int {
	storeIndex:=left
	pivot:=arr[right]
	for i:=left;i<right ;i++  {
		if arr[i]<pivot {
			swap(arr,storeIndex,i)
			storeIndex++
		}
	}
	swap(arr,right,storeIndex)
	return storeIndex
}

func sort(arr []int,left,right int)  {
	if left>right {
		return
	}
	storeIndex := partition(arr, left, right)
	sort(arr, left, storeIndex - 1)
	sort(arr, storeIndex + 1, right)

}

func quickSort(arr []int) []int{

	length:=len(arr)
	sort(arr, 0, length - 1)
	return arr
}

func main() {

	t1:=time.Now()
	arr:=utils.RandArray(10)
	fmt.Printf("之前的数组：%v\n",arr)
	arr=quickSort(arr)
	fmt.Printf("Sort排序后的数组：%v\n",arr)
	t2:=time.Since(t1)
	fmt.Printf("运算时间：%d",t2)

}