package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	arr := make([]int, n)

	for idx, val := range str_arr {
		arr[idx], err = strconv.Atoi(val)
		if err != nil {
			break
		}
	}

	index := make([]int, 0, n)
	for ind := 1; ind <= n; ind++ {
		index = append(index, ind)
	}
	fmt.Println(strings.Trim(fmt.Sprint(Merge(arr, index)), "[]"))
}

func Merge(arr, index []int) []int {
	if len(arr) == 1 {
		return arr
	}
	if len(index) == 1 {
		return index
	}

	left := arr[:(len(arr) / 2)]
	lind := index[:(len(index) / 2)]

	right := arr[(len(arr) / 2):]
	rind := index[(len(index) / 2):]

	return MSort(Merge(left, lind), Merge(right, rind), lind, rind)
}

func MSort(left, right, lind, rind []int) []int {
	i, j := 0, 0
	result := make([]int, len(left)+len(right))
	//
	for k := 0; k < (len(left) + len(right)); k++ {
		if j == len(right) || (i < len(left) && left[i] <= right[j]) {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}

	fmt.Println(lind[0], rind[len(rind)-1], result[0], result[len(result)-1])

	return result
}
