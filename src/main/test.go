package main

import "fmt"


/**
-----------------------------------------------------------------
Golang基本类型:

bool (true, false )
数字内型 (有符号/无符号，有长度/无长度)
string (内建"UTF-8 string" )
array ([n]<type> )
slice (array[i:j])
map (map[<from_type>]<to_type> )
chan
error
-----------------------------------------------------------------
Golang数字类型:

无长度
int, uint
有长度
int8, int16, int32, int64
byte/uint8, uint16, uint32, uint64
float32, float64
 -----------------------------------------------------------------
string类型:
string类型是非常常见的，在Go中，它是UTF-8编码的、由双引号(”)包裹的字符序列。
一旦给变量赋值，字符串就不能修改了：在Go中字符串是不可变的。从C来的用户，下面的情况在Go中是非法的。
-----------------------------------------------------------------

*/
func main() {


	//普通定义变量方式
    var a int = 1
    fmt.Println(a)

 	//函数内部定义变量方式，:= 这种写法只能在函数内部使用
    b := 2
    fmt.Println(b)
    c, d := 3, 4 
    fmt.Println(c)
    fmt.Println(d)
    str1 := "hello world 1"
    fmt.Println(str1)
    var str2 string = "hello world 2"
    fmt.Println(str2)
    var e bool = true
    fmt.Println(e)

    // 声明却未赋值的变量，会编译出错
    // var test string   //compile error:test decleared and not used

	//如果变量声明之后，没有赋值，Golang会给它赋默认值
 	var aa int //默认值初始化 0
    var str string //默认值为空字符串
    var boo bool //默认值为false
    fmt.Println(aa)
    fmt.Println(str)
    fmt.Println(boo)
    
    //字符串
    var str_str string = "hello"
    fmt.Println(str_str)
	//s[0] = 'c'   //修改第一个字符为'c'，这会报错
	
	//If
	var x = 1
	if x > 0 {
    	fmt.Println(x)
	} else {
   	 	fmt.Println(x+1)
	}
	
	for i:=1;i<10;i++{
        fmt.Println(i)
    }
    
    var test string = "asdfghjkl"
    fmt.Println(reverse([]byte(test)))
    
    
    //数组,默认初始值都是0，长度是10
    var array1 [10]int
    array1[0] = 10
    fmt.Println("array value1s:", array1) 
    
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array values:", array2) 
	
	
	//如下情况，Go会自动统计元素个数
    arr3 := [...]int{1, 2, 3}
    fmt.Println(arr3)
    
    //二维数组,如下3中方式都可以
    //a := [2][2]int{ [2]int{1,2}, [2]int{3,4} }
    //a := [2][2]int{ [...]int{1,2}, [...]int{3,4} }
    add := [2][2]int{{1, 2}, {3, 4}}
    fmt.Println(a)
    fmt.Printf("%d\n", add[0][0]) //1
    fmt.Printf("%d\n", add[0][1]) //2
    fmt.Printf("%d\n", add[1][0]) //3
    fmt.Printf("%d\n", add[1][1]) //4
    
}

// Reverse a
// 数组同样是值类型的：将一个数组赋值给另一个数组，会复制所有的元素。
// 当向函数内传递一个数组的时候，它会获得一个数组的副本，而不是数组的指针。
func reverse(a []byte) string {
    //由于Go没有逗号表达式，而++和--是语句而不是表达式，
    //如果你想在for中执行多个变量，应当使用平行赋值。
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
    return string(a)
}
	 
