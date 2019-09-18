package main

import (
	"container/list"
	"flag"
	"fmt"
	"math"
	"strings"
	"unsafe"
	//_ "github.com/goinaction/code/chapter2/sample/matchers"
)

var name string = "global"
var container = []string{"zero", "one", "two"}

type MyString = string //MyString是string类型的别名类型
//byte是uint8的别名类型，而rune是int32的别名类型。
type MyString2 string // 注意，这里没有等号。 最后的代表实际类型,称为潜在类型
//潜在类型的含义是，某个类型在本质上是哪个类型。
type MyStrings []MyString2

type Object struct {
	name string
}

func main() {
	//pac.Printstr("eqwe")

	//testSelect()
	//testFunc()
	//testRange()
	//testInterface()
	//testError()
	//test()
	//testPanic()

	testString()
}

func testFlag() {
	//var name string
	//flag.StringVar(&name, "name", "everyone", "The greeting object.") // [2]

	//var name = *flag.String("name", "everyone", "The greeting object.")

	name := *flag.String("name", "everyone", "The greeting object.")

	//name := *getTheFlag()

	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
}

func testPoint() {
	object := Object{name: "123"}
	fmt.Println(object.name)
	obj := &object
	fmt.Println((*obj).name)
	ptr := uintptr(unsafe.Pointer(obj))
	o := (*Object)(unsafe.Pointer(ptr))
	fmt.Println(o.name)
	fmt.Println(ptr)
	ref := int(ptr)
	fmt.Println(ref)
}

func testList() {
	var l list.List
	l.PushBack("123")
	addList(&l)
	addElement(l.Front())
	fmt.Println(l.Front().Value)
}

func addList(l *list.List) {
	(*l).PushFront("222")
}
func addElement(e *list.Element) {
	(*e).Value = "222"
}

func testInt() {
	var a = 1111111
	u := uint32(a) & 0x3f
	fmt.Printf("%b", u)

	var srcInt = int16(-255)
	dstInt := int8(srcInt)
	fmt.Printf("%v", dstInt)

	i2 := math.MaxInt64
	i := int32(i2)
	i1 := int32(i2 >> 32)
	fmt.Println(i2)
	fmt.Println(i)
	fmt.Println(i1)
	i3 := int64(uint32(i1))<<32 | int64(uint32(i))
	fmt.Println(i3)

	var tmp1 int32 = 123424021
	var tmp2 int16
	var tmp3 uint16
	tmp2 = int16(tmp1)
	tmp3 = uint16(tmp1)
	fmt.Printf("0x%x,%b,%d\n", tmp1, tmp1, tmp1)
	fmt.Printf("0x%x,%b,%d\n", tmp2, tmp2, tmp2)
	fmt.Printf("0x%x,%b,%d\n", tmp3, tmp3, tmp3)

	fmt.Printf("====================================\n")

	var t1 int32 = 123456789
	var t2 int16
	var t3 uint16
	t2 = int16(t1)
	t3 = uint16(t1)
	fmt.Printf("0x%x,%b,%d\n", t1, t1, t1)
	fmt.Printf("0x%x,%b,%d\n", t2, t2, t2)
	fmt.Printf("0x%x,%b,%d\n", t3, t3, t3)
}

func testString() {
	ss := "foo"
	b := []byte(ss)
	fmt.Println(b)

	//UTF-8编码的三个字节\xe4、\xbd和\xa0合在一起才能代表字符'你'，而\xe5、\xa5和\xbd合
	//在一起才能代表字符'好'。

	i := string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'}) // 你好
	//一个值在从string类型向[]rune类型转换时代表着字符串会被拆分成一个个Unicode字符。
	s := string([]rune{'\u4F60', '\u597D'}) // 你好
	fmt.Println(i)
	fmt.Println(s)

	indexRune := strings.IndexRune("sdadsad;dsada", ';')
	fmt.Println(indexRune)
	indexByte := strings.IndexByte("sdadsad;dsada", ';')
	fmt.Println(indexByte)
}

func getTheFlag() *string {
	return flag.String("name", "everybody", "The greeting object.")
}

func testType() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	//类型断言表达式的语法形式是x.(T)
	//interface{}代表空接口，任何类型都是它的实现类型。
	//{} 代表空代码块 或者是 不包含任何内容的数据结构
	value, ok := interface{}(container).([]string)
	if ok {
		fmt.Printf("Hello, %v!\n", value[0])
	}

}

func testSlice() {
	//strings := [3]string{"a", "b", "c"}
	//strings1 := []string{"a", "b", "c"}//长度可变

	s1 := make([]int, 5) //切片长度为5
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8) //切片长度为5,容量为8
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
}

func testSelect() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

func testFunc() {
	/* 声明函数变量 */
	getSquareRoot := func(x float64, y func(float64) float64) float64 {
		return y(math.Sqrt(x))
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9, func(f float64) float64 {
		fmt.Println(f)
		return f
	}))
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 值传递和引用传递  slice和arrat的区别
func testValueTransform() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	arrayModify(array)
	sliceModify(slice)
	fmt.Println(array)
	fmt.Println(slice)
}

func arrayModify(array [5]int) {
	newArray := array
	newArray[0] = 88
}
func sliceModify(slice []int) {
	newSlice := slice
	newSlice[0] = 88
}

func testRange() {
	nums := []int{2, 3, 4}
	sum := 0
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i])
	}

	for index, num := range nums {
		fmt.Printf("%d -> %d\n", index, num)
		sum += num
	}
	fmt.Println(sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}

type fruit interface {
	getName() string
	setName(name string)
}
type apple struct {
	name string
}

func (a *apple) getName() string {
	return a.name
}
func (a *apple) setName(name string) {
	a.name = name
}

//编译期检查是否实现接口
var _ fruit = &apple{}
var _ fruit = new(apple)

func testInterface() {
	a := apple{"红富士"}
	fmt.Print(a.getName())
	a.setName("树顶红")
	fmt.Print(a.getName())
	f, ok := interface{}(a).(fruit)
	if ok {
		fmt.Println(f.getName())
	}

}

func testPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		panic("three")
	}()
	defer func() {
		panic("two")
	}()
	panic("one")
}

func except() {
	err := recover()
	if err != nil {
		fmt.Print(err)
	}

}
func testError() {
	defer except()
	panic("runtime error")
}

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 使用值接收者实现了一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail 使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

// main 是应用程序的入口
func test() {
	// user 类型的值可以用来调用
	// 使用值接收者声明的方法
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// 指向user 类型值的指针也可以用来调用
	// 使用值接收者声明的方法
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify() //这边会自动的转为(*lisa)

	// user 类型的值可以用来调用
	// 使用指针接收者声明的方法, 其实这边会被自动的 加上&
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// 指向user 类型值的指针可以用来调用
	// 使用指针接收者声明的方法
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
