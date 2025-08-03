package gramma_learning

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 类型转换

type FT float64

// SameTypeChange Go语言中，类型转换是将一个类型的值转换为另一个类型的值。
// 对于一个类型T，都有一个对应的类型转换操作T(x)将值x转换为类型T。如果两个类型具有相同的底层类型或二者都是指向相同底层类型变量的未命名指针类型，则二者是可以相互转换。
func SameTypeChange() {
	f := 3.14                          // f 是 float64 类型
	var t FT = FT(f)                   // 将 float64 转换为 FT 类型
	fmt.Printf("f: %T, t: %T\n", f, t) // 输出：f: float64, t: gramma_learning.FT
}

// ByteAndRune Go语言中，字符串是不可变的（immutable），这意味着一旦创建，字符串的内容就不能被修改。
/*
byte类型，uint8的别名，代表一个ASCII码字符。
rune类型，int32的别名，代表一个UTF-8字符。
字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的，字符串是由byte字节组成，所以字符串的长度是byte字节的长度。 rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成，所以字符串也可以和[]rune类型相互转换。

*/
func ByteAndRune() {
	// 要修改字符串，需要先将其转换成`[]rune或[]byte`，完成后再转换为`string`。
	// 无论哪种转换，都会重新分配内存，并复制字节数组。
	bs := "hello"
	fmt.Printf("%x\n", &bs) // 14000096010
	bytes := []byte(bs)     // 强制类型转换为[]byte 类型

	bytes[0] = 'H' // 修改字符串第一个字符为 H
	rbs := string(bytes)
	fmt.Printf("%x\n", &rbs) // 14000096020
	fmt.Println(rbs)         // 输出：Hello 强类型转换为 string

	rs := "你好"
	runers := []rune(rs)        // 强制类型转换为[]rune 类型
	runers[0] = '狗'             // 修改字符串第一个字符为 狗
	fmt.Println(string(runers)) // 输出：狗好   强类型转换为 string
	fmt.Printf("runers 长度%d\n", len(runers))
	fmt.Printf("rs 长度,string %d\n", len(rs))

	rsbyte := []byte(rs) // 转换成字节数组
	fmt.Printf("rsbyte 长度,[]bytes:%d\n", len(rsbyte))
	fmt.Printf("rsbyte 长度,string:%d\n", len(string(rsbyte)))

}

// 类型断言
// 类型断言是一个使用在接口值上的操作。语法上它看起来像x.(T)被称为断言类型，这里x表示一个接口的类型，T表示一个类型。
// JudgeType 断言
func JudgeType() {
	// 成功：第一个返回值为x的动态值，第二个返回值为true
	// 失败：第一个返回值为x的类型零值，第二个返回值为false
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File)      // success:  ok, f == os.Stdout
	b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil

	fmt.Println("f:", f, "ok:", ok) // 输出：f: &{0xc00000a080} ok: true
	fmt.Println("b:", b, "ok:", ok) // 输出：b: <nil> ok: false
}
