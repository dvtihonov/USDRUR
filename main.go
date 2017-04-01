// USDRUR project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"time"
)

func Scan1() string {
	fmt.Println("Введите актуальный курс:")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func Scan2() string {
	fmt.Println("Введите колличество USD:")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func main() {
	KURS, _ := strconv.ParseInt(Scan1(), 10, 64)
	USD, _ := strconv.ParseInt(Scan2(), 10, 64)
	var RUR int64
	RUR = KURS * USD
	fmt.Println("Колличество в рублях : ", RUR)

	//duration := time.Duration(15) * time.Second
	//time.Sleep(duration)
}
