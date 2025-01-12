package main

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

func main() {
	func1()
	fmt.Println()
	func2()
	fmt.Println()
	func3()

	a, b := 1, 3
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)

	var a1 uint = 1
	var b1 uint = 2
	fmt.Println(a1, b1)
	fmt.Println(a1 - b1)

	orderId := "AAAC-2412302123"
	serialNumbers := strings.Split(orderId, "-")
	sn := serialNumbers[1][6:]
	fmt.Println(sn)
}

func func1() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", j, i, j*i)
		}
		fmt.Println()
	}
}

func func2() {
	str := "this is a string"

	len := utf8.RuneCountInString(str)

	for i := 0; i < len; i++ {

		fmt.Printf("%s\n", string(str[i]))
	}

	//for _, s := range str {
	//	fmt.Printf("%s\n", string(s))
	//}
}

func func3() {
	books := map[string]map[string]int{
		"四书": map[string]int{"论语": 80, "大学": 66, "中庸": 60, "孟子": 70},
		"五经": map[string]int{"周易": 90, "诗书": 80, "礼记": 88, "尚书": 78, "春秋": 99},
		"书法": map[string]int{"兰亭集序": 66, "九成宫碑": 68, "多宝塔": 56},
	}
	for key, value := range books {
		slice := []string{}

		for v := range value {
			slice = append(slice, v)
		}
		fmt.Printf("%s : %s\n", key, strings.Join(slice, ","))

	}
}

func func4(s []int, b ...int) {}

type Person struct {
	Name string
	Age  int
}

// 使用 *Person 作为 receiver 参数的方法
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

func (p *Person) UpdateName(name string) {
	p.Name = name
}

type MyStruct struct {
	privateField int
}

type MyMethods struct{}

func (m *MyMethods) MyMethod() string {
	return "Hello, World!"
}
func main2() {
	p := &Person{Name: "Forest", Age: 24}

	// 使用 *Person 类型变量调用 UpdateAge 方法
	p.UpdateAge(30)
	fmt.Println("Age after update:", p.Age) // 输出: Age after update: 30

	// 使用 Person 类型变量调用 UpdateAge 方法会导致编译错误
	p.UpdateName("aabb")                     // 编译错误: cannot use p (type Person) as type *Person in function argument
	fmt.Println("Age after update:", p.Name) // 输出: Age after update: 30

	ctx := context.TODO()
	withValue := context.WithValue(ctx, "aaa", "bbb")

	value := withValue.Value("aaa")
	fmt.Println("aaa=", value)

	v := reflect.ValueOf(p).Elem().FieldByName("Name")
	fmt.Println("v name", v.String())

	s := MyStruct{privateField: 1}
	v = reflect.ValueOf(&s)

	field := v.Elem().FieldByName("privateField")
	fmt.Println("Private Field:", field.Int()) // 输出: Private Field: 1
	func5()
	func6()
	func7()

}

func func5() {
	var x float64 = 3.4
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(&x)

	fmt.Println("Type:", t)              // 输出: Type: float64
	fmt.Println("Value:", v.Interface()) // 输出: Value: 3.4

	// 检查是否可设置
	//if v.Kind() == reflect.Ptr {
	//	v = v.Elem() // 获取指针指向的实际值
	//	if v.CanSet() {
	//		// 修改变量的值
	//		v.SetFloat(7.8)
	//	} else {
	//		fmt.Println("Value is not settable")
	//	}
	//}

	v.Elem().SetFloat(7.1)
	fmt.Println(x)

}

func func6() {
	var i interface{} = "hello"
	v := reflect.ValueOf(i)

	if s, ok := v.Interface().(string); ok {
		fmt.Println(s) // 输出: hello
	}
}

func func7() {
	obj := &MyMethods{}
	method := reflect.ValueOf(obj).MethodByName("MyMethod")
	result := method.Call(nil)
	fmt.Println("Method Result:", result[0].Interface()) // 输出: Method Result: Hello, World!

	//arr := []int{1, 2, 3, 4, 5}
	var arr [5]int
	fmt.Println(arr)

}
