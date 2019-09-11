package main

import "fmt"

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Print(err)
		}
	}()

	error1 := testError1()
	fmt.Println(error1)
}
func testError1() string{
	panic("runtime error")
	return "ok"
}