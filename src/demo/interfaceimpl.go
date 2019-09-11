package main

import "fmt"

/*
接口实现 指针和不是指针
方法集
 */

type notifier interface {
	notify()
}

type user1 struct {
	name  string
	email string
}

func (u user1) get() string {
	return u.name
}
func (u *user1) set(name string) {
	u.name = name
}

//使用湿疹接受者来实现一个接口,那么只有指向那个类型的指针才能够实现对应的接口
func (u *user1) notify() {
	fmt.Printf("send user email to %s<%s>\n", u.name, u.email)
}
func (u *admin) notify() {
	fmt.Printf("send admin email to %s<%s>\n", u.name, u.email)
}

//如果使用只接受者来实现一个借口,那么那个类型的值和指针都能够实现对应的接口
//func (u user1) notify(){
//	fmt.Printf("send user email to %s<%s>\n", u.name,u.email)
//}

func main() {
	admin := admin{
		user1: user1{name: "admin", email: "111"},
		level: "1",
	}

	user := user1{name: "123", email: "333"}
	point := &user
	user.notify()
	point.notify()
	notify1(&user)
	notify1(point)

	admin.user1.notify()
	admin.notify()
	notify1(&admin)
	notify1(&(admin.user1))
}

func notify1(n notifier) {
	n.notify()
}

type admin struct {
	user1
	level string
}
