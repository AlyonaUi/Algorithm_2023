package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NumOfDifferent() {
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

	count := 1
	dif := MergeD(arr)
	for idx := 0; idx < len(dif)-1; idx++ {
		if dif[idx] != dif[idx+1] {
			count++
		}
	}

	fmt.Println(count)
}

func MergeD(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	left := arr[:(len(arr) / 2)]

	right := arr[(len(arr) / 2):]

	return MSortD(MergeD(left), MergeD(right))
}

func MSortD(left, right []int) []int {
	i, j := 0, 0
	result := make([]int, len(left)+len(right))

	for k := 0; k < (len(left) + len(right)); k++ {
		if j == len(right) || (i < len(left) && left[i] <= right[j]) {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}

	return result
}
