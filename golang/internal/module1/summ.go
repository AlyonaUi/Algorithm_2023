package module1

import "fmt"

func Summ() {
	var a, b int
	_, err := fmt.Scanln(&a, &b)
    if err != nil {
        return
    }
	summ := a + b
	fmt.Println(summ)
}