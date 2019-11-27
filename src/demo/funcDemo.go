package main

import "fmt"

type iinterface interface {
	f(string) string
}

type ifunc func(string) (string)

// 使用函数去实现某个接口
func (i ifunc) f(s string) string{
	return i(s)
}



func demofunc (s string) string{
	fmt.Println(s)
	return "hello "+s
}

func dunctest(f func(s string) string, s string){
	fmt.Println(f(s))
}

func dunctest2(i ifunc, s string){
	fmt.Println(i(s))
}

func dunctest3(i iinterface, s string){
	fmt.Println(i.f(s))
}

func main() {
	d := func (s string) string{
		fmt.Println(s)
		return "hello "+s
	}("111")

	dunctest(demofunc, "123")
	dunctest2(demofunc, "123")
	dunctest2(func (s string) string{
				fmt.Println(s)
				return "hello "+s
			}, "123")
	dunctest3(ifunc(demofunc), "!23")

}