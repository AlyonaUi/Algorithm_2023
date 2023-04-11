package module2

import (
	"fmt"
	"strings"
)

func Sortofpairs() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}
	type pair struct {
		num  int
		cost int
	}
	arr := make([]pair, 0, 2)

	for {
		var num, cost int
		_, err := fmt.Scanln(&num, &cost)
		if err != nil {
			break
		}
		arr = append(arr, pair{num, cost})
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j].cost < arr[j+1].cost {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else if arr[j].cost == arr[j+1].cost {
				if arr[j].num > arr[j+1].num {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}

	for _, val := range arr {
		fmt.Println(strings.Trim(fmt.Sprint(val), "{}"))
	}
}
