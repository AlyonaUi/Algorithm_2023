package module2

import (
	"fmt"
	"strings"
)

func Bubble_sort() {
	var size int
	_, err := fmt.Scanln(&size)
	if err != nil {
		return
	}
	arr := make([]int, 0, size)
	count := 0
	for {
		var rec int
		fmt.Scan(&rec)
		arr = append(arr, rec)
		if len(arr) == size {
			break
		}
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				count += 1
				fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
			}
		}
	}
	if count == 0 {
		fmt.Println("0")
	}
}
