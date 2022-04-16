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

	scan1 := Scan1()
	KURS, err := strconv.ParseInt(scan1, 10, 64)
	fmt.Println("Введен актуальный курс USD: ", scan1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка! Чувак ты мне хрень написал, а не число - ", scan1)
		return
	}

	scan2 := Scan2()
	USD, err := strconv.ParseInt(scan2, 10, 64)
	fmt.Println("Введено колличество USD: ", scan2)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка! Чувак ты мне хрень написал, а не число - ", scan2)
		return
	}

	var RUR int64
	RUR = KURS * USD
	fmt.Println("Колличество в рублях : ", RUR)

	//duration := time.Duration(15) * time.Second
	//time.Sleep(duration)
}
