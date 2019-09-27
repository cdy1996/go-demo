package main

import (
	"fmt"
	"github.com/codegangsta/inject"
	"reflect"
)

type SpecialString interface {
}

type TestStruct struct {
	Dep1 string        `inject:"t" json:"-"`
	Dep2 SpecialString `inject`
	Dep3 string
}

type Greeter struct {
	Name string
}

func (g *Greeter) String() string {
	return "Hello, My name is" + g.Name
}

func main() {
	//injectorApply()
	injectorGet()
}

func invokeTest() {
	injector := inject.New()
	dep := "some dependency"
	injector.Map(dep)
	dep2 := "another dep"
	injector.MapTo(dep2, (*SpecialString)(nil))

	dep3 := make(chan *SpecialString)
	dep4 := make(chan *SpecialString)

	// 返回方向为<- 的chain类型
	typRecv := reflect.ChanOf(reflect.RecvDir, reflect.TypeOf(dep3).Elem())
	typSend := reflect.ChanOf(reflect.SendDir, reflect.TypeOf(dep4).Elem())

	injector.Set(typRecv, reflect.ValueOf(dep3))
	injector.Set(typSend, reflect.ValueOf(dep4))

	_, _ = injector.Invoke(func(d1 string, d2 SpecialString, d3 <-chan *SpecialString, d4 chan<- *SpecialString) {
		fmt.Println(d1, dep)
		fmt.Println(d2, dep2)
		fmt.Println(reflect.TypeOf(d3).Elem(), reflect.TypeOf(dep3).Elem())
		fmt.Println(reflect.TypeOf(d4).Elem(), reflect.TypeOf(dep4).Elem())
		fmt.Println(reflect.TypeOf(d3).ChanDir(), reflect.RecvDir)
		fmt.Println(reflect.TypeOf(d4).ChanDir(), reflect.SendDir)
	})
}

func injectorApply() {
	injector := inject.New()

	injector.Map("a dep").MapTo("another dep", (*SpecialString)(nil))

	s := TestStruct{}
	_ = injector.Apply(&s)
	fmt.Println(s.Dep1, "a dep")
	fmt.Println(s.Dep2, "another dep")
	fmt.Println(s.Dep3, "")
}

func interfaceOf() {
	iType := inject.InterfaceOf((*SpecialString)(nil))
	fmt.Println(iType.Kind(), reflect.Interface)

	iType = inject.InterfaceOf((**SpecialString)(nil))
	fmt.Println(iType.Kind(), reflect.Interface)

}

func injectorSet() {
	injector := inject.New()
	typ := reflect.TypeOf("string")
	typSend := reflect.ChanOf(reflect.SendDir, typ)
	typRecv := reflect.ChanOf(reflect.RecvDir, typ)

	// instantiating unidirectional channels is not possible using reflect
	// http://golang.org/src/pkg/reflect/value.go?s=60463:60504#L2064
	chanRecv := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, typ), 0)
	chanSend := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, typ), 0)

	injector.Set(typSend, chanSend)
	injector.Set(typRecv, chanRecv)

	fmt.Println(injector.Get(typSend).IsValid(), true)
	fmt.Println(injector.Get(typRecv).IsValid(), true)
	fmt.Println(injector.Get(chanSend.Type()).IsValid(), false)
}

func injectorGet() {
	injector := inject.New()

	injector.Map("some dependency")

	fmt.Println(injector.Get(reflect.TypeOf("string")).String())
	fmt.Println(injector.Get(reflect.TypeOf("string")).IsValid(), true)
	fmt.Println(injector.Get(reflect.TypeOf(11)).IsValid(), false)
}

func injectorSetParent() {
	injector := inject.New()
	injector.MapTo("another dep", (*SpecialString)(nil))

	injector2 := inject.New()
	injector2.SetParent(injector)

	fmt.Println(injector2.Get(inject.InterfaceOf((*SpecialString)(nil))).IsValid(), true)
}

func injectImplementors() {
	injector := inject.New()
	g := &Greeter{"Jeremy"}
	injector.Map(g)

	fmt.Println(injector.Get(inject.InterfaceOf((*fmt.Stringer)(nil))).IsValid(), true)
}
