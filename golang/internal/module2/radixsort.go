package module2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RadixSort() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	arr := make([]int, n)

	for idx := 0; idx < n; idx++ {
		scanner.Scan()
		arr[idx], err = strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
	}
	Sort(arr, n)
}

func Sort(arr []int, n int) {
	fmt.Println("Initial array:")

	for b, v := range arr {
		if b < len(arr)-1 {
			fmt.Printf("%d, ", v)
		} else {
			fmt.Printf("%d\n", v)
		}
	}
	fmt.Println("**********")
	digit := 1
	sorted_arr := make([]int, n)

	for phase := 1; phase <= len(strconv.Itoa(arr[0])); phase++ {
		fmt.Printf("Phase %d\n", phase)

		bucket := [10]int{0}

		for i := 0; i < n; i++ {
			bucket[(arr[i]/digit)%10]++
		}
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			bucket[(arr[i]/digit)%10]--
			sorted_arr[bucket[(arr[i]/digit)%10]] = arr[i]
		}
		for i := 0; i < n; i++ {
			arr[i] = sorted_arr[i]
		}

		for num := 0; num < 10; num++ {
			bag := make([]int, 0, n)
			for _, value := range arr {
				if (value/digit)%10 == num {
					bag = append(bag, value)
				}
			}
			if len(bag) != 0 {
				fmt.Printf("Bucket %d: %v\n", num, strings.Trim(fmt.Sprint(bag), "[]"))
			} else {
				fmt.Printf("Bucket %d: empty\n", num)
			}
		}
		digit *= 10
		fmt.Println("**********")
	}
	fmt.Println("Sorted array:")

	for b, v := range arr {
		if b < len(arr)-1 {
			fmt.Printf("%d, ", v)
		} else {
			fmt.Printf("%d\n", v)
		}
	}
}
