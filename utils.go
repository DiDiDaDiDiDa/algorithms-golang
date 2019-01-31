package utils

import "math/rand"

func RandArray(n int) []int {

    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = rand.Intn(n)
    }

    return arr
}
