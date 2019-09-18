package main

import (
	"fmt"
	"reflect"
)

type user2 struct {
	name string
}

func (self *user2) getName() string {
	return self.name
}
func (self user2) getName2() string {
	return self.name
}

func getName() string {
	return "12"
}

func main() {
	user2 := &user2{name: "123"}

	value := reflect.ValueOf(user2)
	fmt.Println(value)
	fmt.Println(value.Elem())
	fmt.Println(value.Method(0).Call(nil))
	value.Field(0).SetString("333")
	fmt.Println(value.Field(0).String())
	type1 := reflect.TypeOf(user2)
	fmt.Println(type1)
	fmt.Println(type1.Name())
	fmt.Println(type1.Method(0))
	fmt.Println(type1.NumIn())

}
