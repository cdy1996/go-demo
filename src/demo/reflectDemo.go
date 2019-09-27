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
	structTest()
}

func typeOf() {
	dep := "some dependency"
	of := reflect.TypeOf(dep)
	fmt.Println(of)
	dep1 := "some dependency1"
	typeOf := reflect.TypeOf(dep1)
	fmt.Println(typeOf)
	fmt.Println(typeOf == of)

}

func valueOf() {
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

func structTest() {
	//创建一个结构体Student
	type Student struct {
		Name   string
		School string `level:"national" id:"1"` //注明有两个标签level和id，用于给字段添加自定义信息，方便其他模块根据信息进行不同功能的处理
		//注意，注明标签的时候不要随便添加空格，否则会因为格式的问题在调用Get()函数时得不到结果
	}
	stu := Student{"xiaoming", "peking university"}

	//获取结构体实例的反射类型对象
	typeStu := reflect.TypeOf(stu)

	fmt.Println(typeStu.NumField()) //2
	//遍历获得结构体的所有字段
	for i := 0; i < typeStu.NumField(); i++ {
		//获取对应的字段类型
		fieldType := typeStu.Field(i)
		fmt.Printf("fieldType: %v\n", fieldType)
		//循环两次返回：
		//fieldType: {Name  string  0 [0] false}
		//fieldType: {School  string level:"national" id:"1" 16 [1] false}

		//打印获取的字段名以及其注明的标签信息
		fmt.Printf("name : %v tag : %v\n", fieldType.Name, fieldType.Tag) //fieldType.Tag返回StructTag类型的值，然后可以调用Get()函数得到具体的tag值
		//循环两次返回：
		//name : Name tag :
		//name : School tag : level:"national" id:"1"
	}
	//通过字段名，找到字段的类型信息
	if studentSchool, ok := typeStu.FieldByName("School"); ok {
		// fmt.Printf("studentSchool: %v\n", studentSchool)//studentSchool: {School  string level:"national" id:"1" 16 [1] false}

		//使用Get()获取标签的value
		fmt.Println(studentSchool.Tag.Get("level"), studentSchool.Tag.Get("id")) //national 1
	}

}

func ptrTo() {
	//创建一个结构体Student
	type Student struct {
		name   string
		school string
	}
	//创建一个Student实例指针
	stuPtr1 := &Student{"xiaoming", "peking university"}
	typeStuPtr1 := reflect.TypeOf(stuPtr1)
	fmt.Println(typeStuPtr1) //*main.Student

	fmt.Println(typeStuPtr1.Name()) //为空
	fmt.Println(typeStuPtr1.Kind()) //ptr

	//获取该指针的元素
	typeStu1 := typeStuPtr1.Elem()

	fmt.Println(typeStu1.Name()) //Student
	fmt.Println(typeStu1.Kind()) //struct

	//上面直接传入指针给TypeOf和下面使用PtrTo将其转成指针类型的效果是一样的
	stu2 := Student{"zhangwei", "qinghua university"}
	typeStu2 := reflect.TypeOf(stu2)
	fmt.Println(typeStu2) //main.Student
	typeStuPtr2 := reflect.PtrTo(typeStu2)
	fmt.Println(typeStuPtr2) //*main.Student

	fmt.Println(typeStuPtr2.Name()) //为空
	fmt.Println(typeStuPtr2.Kind()) //ptr

	//获取该指针的元素
	typeStu22 := typeStuPtr2.Elem()
	fmt.Println(typeStu22.Name()) //Student
	fmt.Println(typeStu22.Kind()) //struct
}
