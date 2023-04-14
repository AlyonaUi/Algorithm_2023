package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Index int
	Value int
}

func Mergesort() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	str_arr := strings.Split(line, " ")

	arr := make([]Item, n)

	for idx, val := range str_arr {
		value, err := strconv.Atoi(val)
		if err != nil {
			break
		}
		arr[idx] = Item{Index: idx + 1, Value: value}
	}
	fmt.Println(strings.Trim(fmt.Sprint(Merge(arr)), "[]"))
}

func Merge(arr []Item) []Item {
	if len(arr) == 1 {
		return arr
	}

	left := arr[:(len(arr) / 2)]

	right := arr[(len(arr) / 2):]

	return MSort(Merge(left), Merge(right))
}

func MSort(left, right []Item) []Item {
	i, j := 0, 0
	result := make([]Item, len(left)+len(right))

	for k := 0; k < (len(left) + len(right)); k++ {
		if j == len(right) || (i < len(left) && left[i].Value <= right[j].Value) {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}

	fmt.Println(result[0].Value, result[len(result)-1].Value, result[0].Index, result[len(result)-1].Index)

	return result
}
