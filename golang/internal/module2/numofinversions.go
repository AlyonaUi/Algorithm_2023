package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NumofInversions() {
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

	fmt.Println(MergeSortInv(arr, 0, len(arr)-1))
}

func MergeSortInv(arr []int, l_idx, r_idx int) int {
	var count_inv int = 0

	if l_idx < r_idx {
		var m_idx int = (l_idx + r_idx) / 2

		count_inv += MergeSortInv(arr, l_idx, m_idx)

		count_inv += MergeSortInv(arr, m_idx+1, r_idx)

		count_inv += MSortForInv(arr, l_idx, m_idx, r_idx)
	}
	return count_inv
}

func MSortForInv(arr []int, l_idx, m_idx, r_idx int) int {
	left := make([]int, 0, m_idx-l_idx+1)
	for i := l_idx; i <= m_idx; i++ {
		left = append(left, arr[i])
	}

	right := make([]int, 0, r_idx-m_idx)
	for j := m_idx + 1; j <= r_idx; j++ {
		right = append(right, arr[j])
	}

	i, j, inv := 0, 0, 0

	for k := l_idx; k < (len(arr) - (len(arr) - r_idx) + 1); k++ {
		if j == len(right) || (i < len(left) && left[i] <= right[j]) {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
			if i < len(left) {
				inv += len(left) - i
			}
		}
	}
	return inv
}
