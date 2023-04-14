package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Storage() {
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

	arr1 := make([]int, n)

	for idx, val_arr := range str_arr {
		arr1[idx], err = strconv.Atoi(val_arr)
		if err != nil {
			break
		}
	}
	var k int
	_, err = fmt.Scanln(&k)
	if err != nil {
		return
	}
	reader = bufio.NewReader(os.Stdin)
	line, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	str_arr1 := strings.Split(line, " ")

	arr2 := make([]int, k)

	for idx, val_arr := range str_arr1 {
		arr2[idx], err = strconv.Atoi(val_arr)
		if err != nil {
			break
		}
	}

	counting := make([]int, n)
	for _, val := range arr1 {
		counting[val-1]++
	}

	for j := 0; j < n; j++ {
		if arr1[j] >= counting[j] {
			fmt.Println(`no`)
		} else {
			fmt.Println(`yes`)
		}
	}
}
