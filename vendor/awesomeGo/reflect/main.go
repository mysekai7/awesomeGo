package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName(str string) string {
	this.name = str
	return this.name + "==="
}

func main() {

	// 备注: reflect.Indirect -> 如果是指针则返回 Elem()
	// 首先，reflect包有两个数据类型我们必须知道，一个是Type，一个是Value。
	// Type就是定义的类型的一个数据类型，Value是值的类型

	// 对象
	s := "this is string"

	// 获取对象类型 (string)
	fmt.Println(reflect.TypeOf(s))

	// 获取对象值 (this is string)
	fmt.Println(reflect.ValueOf(s))

	// 对象
	var x float64 = 3.4

	// 获取对象值 (<float64 Value>)
	fmt.Println(reflect.ValueOf(x))

	// 对象
	a := &MyStruct{name: "nljb"}

	// 返回对象的方法的数量 (1)
	fmt.Println(reflect.TypeOf(a).NumMethod())	// 1
	fmt.Println(reflect.ValueOf(a).Elem())	// {nljb}

	// 遍历对象中的方法
	for m := 0; m < reflect.TypeOf(a).NumMethod(); m++ {
		fmt.Println(reflect.TypeOf(a))	// *main.MyStruct

		method := reflect.TypeOf(a).Method(m)
		fmt.Println(method.Type)         // func(*main.MyStruct) string
		fmt.Println(method.Name)         // GetName
		fmt.Println(method.Type.NumIn()) // 参数个数
		fmt.Println(method.Type.In(1))   // 参数类型
	}

	// 获取对象值 (<*main.MyStruct Value>)
	fmt.Println(reflect.ValueOf(a))		// &{nljb}

	fmt.Println(reflect.Indirect(reflect.ValueOf(a)))		// {nljb}

	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).Type())		// main.MyStruct

	// 获取对象名称
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).Type().Name())	// MyStruct

	// 参数
	i := "Hello"
	v := make([]reflect.Value, 0)
	fmt.Println(v)	//[]

	fmt.Println(reflect.ValueOf(i))	// Hello

	v = append(v, reflect.ValueOf(i))
	fmt.Println(v)	//[Hello]

	// 通过对象值中的方法名称调用方法 ([nljb]) (返回数组因为Go支持多值返回)
	fmt.Println(reflect.ValueOf(a).MethodByName("GetName").Call(v))	// [Hello===]

	// 通过对值中的子对象名称获取值 (nljb)
	fmt.Println(reflect.ValueOf(a))	// &{Hello}
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)))	// {Hello}
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).FieldByName("name"))	// Hello


	// 是否可以改变这个值 (false)
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)).FieldByName("name").CanSet())	//false

	// 是否可以改变这个值 (true)
	fmt.Println(&(a.name))	//0xc42000e210
	fmt.Println(reflect.ValueOf(&(a.name)))	//0xc42007a1f0
	fmt.Println(reflect.ValueOf(&(a.name)).CanSet())	//false
	fmt.Println(reflect.Indirect(reflect.ValueOf(&(a.name))))	//Hello
	fmt.Println(reflect.Indirect(reflect.ValueOf(&(a.name))).CanSet())	//true

	// 不可改变 (false)
	fmt.Println(reflect.ValueOf(s))	//this is string
	fmt.Println(reflect.ValueOf(s).CanSet())//false
	fmt.Println(reflect.Indirect(reflect.ValueOf(s)))	// this is string
	fmt.Println(reflect.Indirect(reflect.ValueOf(s)).CanSet())//false


	// 可以改变
	// reflect.Indirect(reflect.ValueOf(&s)).SetString("jbnl")
	fmt.Println(reflect.ValueOf(&s))	//0xc42000e1e0
	fmt.Println(reflect.ValueOf(&s).CanSet())	//false
	fmt.Println(reflect.Indirect(reflect.ValueOf(&s)))//this is string
	fmt.Println(reflect.Indirect(reflect.ValueOf(&s)).CanSet())//true

}