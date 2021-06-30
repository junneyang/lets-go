package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"unicode/utf8"

	"example.com/main/calculate"
	"rsc.io/quote"
)

func swap(a int, b int) (int, int) {
	return b, a
}

func modifyArr(a [5]int) [5]int {
	a[0] = 88
	return a
}

func modifySlice(a []int) []int {
	a[0] = 88
	return a
}

func main() {
	// HelloWorld
	fmt.Println("HelloWorld")
	fmt.Println(quote.Go())

	// 基本数据类型
	var i int = 100
	var message string = "Hello, World"
	f := 1.234
	fmt.Println(i)
	fmt.Println(message)
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))

	// 字符串与UTF8编码
	str1 := "Hello"
	str2 := "Hi您好"
	fmt.Println(str1[2], string(str1[2]), reflect.TypeOf(str1[2]))
	fmt.Println(str2[2], string(str2[2]), reflect.TypeOf(str2[2]))
	fmt.Println(len(str2))
	fmt.Println(len([]rune(str2)), utf8.RuneCountInString(str2))
	fmt.Println([]rune(str2)[2], string([]rune(str2)[2]))

	// 数组与切片-数组遍历
	arr := [5]int{1, 2, 3, 4, 5}
	for _, v := range arr {
		fmt.Println(v)
	}
	for i, v := range arr {
		fmt.Printf("%T, i=>%d, v=>%v\n", i, i, v)
	}
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Println(arr)
	a, _ := swap(123, 321)
	fmt.Println(a)

	// 数组与切片-切片构造与操作
	arr = [5]int{1, 2, 3, 4, 5}
	slice := arr[:]
	fmt.Println(arr, reflect.TypeOf(arr))
	fmt.Println(slice, reflect.TypeOf(slice))
	slice1 := append(slice, 6, 7, 8)
	fmt.Println(len(slice1), cap(slice1), slice1, slice)
	copy(slice, []int{6, 7, 8})
	fmt.Println(slice)
	// 通过...自动计算数组大小
	data := [...]int{1, 2, 3, 4, 5}
	fmt.Println(data, reflect.TypeOf(data))
	// 通过make构造slice对象
	data2 := make([]int, 2, 5)
	fmt.Println(data2)
	data2 = append(data2, 1, 2, 3, 4, 5)
	fmt.Println(data2)

	// 参数均是值传递，只是某些对象是通过指针包装的对象
	arr[0] = 66
	fmt.Println(arr)
	slice[0] = 66
	fmt.Println(slice)
	arr_new := arr
	arr_new[0] = 666
	fmt.Println(arr, arr_new, reflect.TypeOf(arr_new))
	slice_new := slice
	slice_new[0] = 666
	fmt.Println(slice, slice_new, reflect.TypeOf(slice_new))
	fmt.Println(modifyArr(arr_new))
	fmt.Println(modifySlice(slice_new))
	fmt.Println(arr_new)
	fmt.Println(slice_new)

	// map哈希-构造与使用，slice、map、channel使用之前必须初始化
	m1 := make(map[string]int)
	m1["abc"] = 123
	m1["def"] = 456
	fmt.Println(m1)
	// m2 := new(map[string]int)
	// m3 := *m2
	// m3["aaa"] = 111
	// fmt.Println(m3)
	m1["abc"] += 321
	fmt.Println(m1)

	// 指针与引用
	str := "Hello"
	var p *string = &str
	*p = "World"
	fmt.Println(str)

	// 流程控制-if-else
	con := 10
	if con < 18 {
		fmt.Println("Kid")
	} else if con > 60 {
		fmt.Println("Old")
	} else {
		fmt.Println("Adult")
	}
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}
	// 流程控制-switch-case
	gender := MALE
	switch gender {
	case MALE:
		fmt.Printf("%s, %d\n", gender, MALE)
	case FEMALE:
		fmt.Printf("%s, %d\n", gender, FEMALE)
	default:
		fmt.Printf("%s, %d\n", gender, UNKNOWN)
	}
	// 流程控制-for
	hash := map[string]int{
		"ABC": 123,
		"DEF": 456,
		"HIJ": 789,
	}
	for k, v := range hash {
		fmt.Printf("%v=>%v\n", k, v)
	}

	// 函数
	fmt.Println(demo("Hello"))
	// fmt.Println(Demo("Hello"))

	// 错误处理
	lines, _ := fileOps("file.txt")
	fmt.Println(lines)
	defer defer_demo()
	// fmt.Println(slice[10])
	fmt.Println(slice)

	// 面向对象-结构体
	// stu := new(Student)
	// stu.name = "zhagsan"
	// stu.age = 18
	stu := &Student{name: "zhangsan", age: 18}
	fmt.Println(stu.Hello())

	// 面向对象-接口
	var hello HelloService = &Teacher{name: "zhangsan", age: 18, course: "english"}
	fmt.Println(hello.GetName(), hello.SayHello())
	hello = &Docter{name: "lisi", age: 18, department: "medicine"}
	fmt.Println(hello.GetName(), hello.SayHello())

	// 面向对象-反射
	t := reflect.TypeOf(Teacher{})
	o := reflect.New(t).Interface()
	tea := o.(*Teacher)
	fmt.Println(tea.GetName(), tea.SayHello())

	// 并发-无需通信
	tnum := 10
	wg := &sync.WaitGroup{}
	wg.Add(tnum)
	for i := 0; i < tnum; i++ {
		go download(wg, fmt.Sprintf("http://www.baidu.com?url=%v.php", i))
	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("download complete...")

	// 并发-GOMAXPROCS配置
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumGoroutine())

	// 并发-通过缓冲channel同步
	ch := make(chan string, tnum)
	defer close(ch)
	for i := 0; i < tnum; i++ {
		go download_ch(ch, fmt.Sprintf("http://www.baidu.com?url=%v.php", i))
	}
	for i := 0; i < tnum; i++ {
		url := <-ch
		fmt.Printf("notify to download: %v\n", url)
	}
	fmt.Println("download complete...")

	// 并发-通过阻塞channel同步
	channel := make(chan string)
	defer close(channel)
	go download_ch_once(channel, "http://www.baidu.com?url=demo.php")
	url := <-channel
	fmt.Printf("notify to download: %v\n", url)
	fmt.Println("download complete...")

	// 包管理
	fmt.Println(calculate.Sum(10, 10))
}
