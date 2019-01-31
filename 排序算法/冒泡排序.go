package main

import (
	"fmt"
	"github.com/algorithms-golang"
	"time"
)

//最初写法
func bubbleSort1(arr []int) ([]int,int) {

	length:=len(arr)
	var count int
	for i:=0;i<length;i++ {
		for j:=i+1;j<length;j++ {
			count++
			if arr[i]>arr[j] {
				tem:=arr[i]
				arr[i]=arr[j]
				arr[j]=tem
			}
		}
	}

	return arr,count
}

//常见写法
//伪代码
//function bubble_sort (array, length) {
//	var i, j;
//	for(i from 1 to length-1){
//		for(j from 0 to length-1-i){
//			if (array[j] > array[j+1])
//				swap(array[j], array[j+1])
//			}
//		}
//	}
func bubbleSort2(arr []int) ([]int,int) {

	length:=len(arr)
	var count int
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			count++
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	return arr,count
}

func main() {

	t1:=time.Now()
 	arr:=utils.RandArray(10000)
	fmt.Printf("之前的数组：%v\n",arr)
	arr1,count1:=bubbleSort1(arr)
	fmt.Printf("Sort1排序后的数组：%v\n",arr1)
	fmt.Printf("Sort1运行次数：%d\n",count1)

	t2:=time.Since(t1)
	fmt.Printf("运算时间：%d\n",t2)


	arr2,count2:=bubbleSort2(arr)
	fmt.Printf("Sort2排序后的数组：%v\n",arr2)
	fmt.Printf("Sort2运行次数：%d\n",count2)

}