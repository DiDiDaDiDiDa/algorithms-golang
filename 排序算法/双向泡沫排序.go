package main

import(
	"fmt"
	"github.com/algorithms-golang"
	"time"
)

//双向泡沫排序
func cocktailSort(arr []int) ([]int,int) {

	var count int
	for i := 0; i < len(arr)/2; i++ {
		left := 0
		right := len(arr) - 1

		for left <= right {

			count++
			if arr[left] > arr[left+1] {
				tmp := arr[left]
				arr[left] = arr[left+1]
				arr[left+1] = tmp
			}

			left++

			if arr[right-1] > arr[right] {
				tmp := arr[right-1]
				arr[right-1] = arr[right]
				arr[right] = tmp
			}

			right--
		}
	}
	return arr,count
}


func main() {

	t1:=time.Now()
	arr:=utils.RandArray(10000)
	fmt.Printf("之前的数组：%v\n",arr)
	arr,count:=cocktailSort(arr)
	fmt.Printf("Sort排序后的数组：%v\n",arr)
	fmt.Printf("Sort运行次数：%d\n",count)
	t2:=time.Since(t1)
	fmt.Printf("运算时间：%d",t2)

}