package main

import "fmt"

func main() {
	//value6:=interface{}(byte(127))

	defer func() {
		err := recover()
		if err != nil {
			fmt.Print(err)
		}
	}()

	error1 := testError1()
	fmt.Println(error1)
}
func testError1() string {
	return testError2()
}

func testError2() string {
	panic("runtime error")
	return "ok"
}
