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

	arr := make([]int, n)

	for idx, valarr := range str_arr {
		arr[idx], err = strconv.Atoi(valarr)
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

	arr1 := make([]int, k)

	for idx, valarr := range str_arr1 {
		arr1[idx], err = strconv.Atoi(valarr)
		if err != nil {
			break
		}
	}

	counting := make([]int, n)
	for _, val := range arr1 {
		counting[val-1]++
	}

	for j := 0; j < n; j++ {
		if arr[j] >= counting[j] {
			fmt.Println(`no`)
		} else {
			fmt.Println(`yes`)
		}
	}
}
