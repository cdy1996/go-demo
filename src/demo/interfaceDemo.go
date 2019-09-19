package main

import "fmt"

type Dog struct {
	name string
}

type Pet interface {
	Name() string
	Category() string
}

func (self *Dog) Name() string {
	return self.name
}
func (self *Dog) Category() string {
	return "dog"
}

func (self *Dog) SetName(name string) {
	self.name = name
}

func main() {
	dog1 := Dog{"little pig"}
	dog2 := dog1
	dog3 := &dog1
	dog1.name = "monster"
	fmt.Println(dog2.name)
	fmt.Println(dog3.name)

	dog := Dog{"little pig"}
	//如果我们使用一个变量给另外一个变量赋值，那么真正赋给后者的，并不是前者持有的那个值，而是该值的一个副本。
	//var pet Pet = dog //这里要把实现的接口方法改为没有指针的
	var pet Pet = &dog
	dog.SetName("monster")
	fmt.Println(pet.Name())
}
