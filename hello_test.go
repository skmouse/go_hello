package main

import "testing"

/**
 * 编写go测试程序
 * 测试方法名以Test开头 eg: func TestXXX(t *testTing .T){}
 */

func TestFirstTry(t *testing.T) {
	t.Log("hello world")
}

/**
 * go变量定义有三种
 * 1.var a =1 todo 自动推断类型
 * 2. a:=1  todo :=这种声明左侧必须是没有申明过的变量
 * 3 var a int
 * 4. var (m bool, n int) todo 一般用于全局变量的声明
 */
func TestVar(t *testing.T)  {
	var a int
	//var d int  会发生错误
	//d :=4
	a = 1
	var b=2
	c :=3
	t.Log(a,b,c)
}
/**
 * 常量
 */
func TestConst(t *testing.T)  {
	const Length int =10
	const Width int = 20
	area := Length *Width
	t.Log(area)
}

/**
 * var a[3] int
 * a[0]=1
 * b:=[3] int {1,2,3}
 * arra3 = [...]{1,2,3}
 */
func TestArray(t *testing.T)  {
	var b =[...]int{1,2,3,4,5,6,7,8}
	//for i,index :=range a{
	//	t.Log(i,index)
	//}
	//for n :=0; n<len(b);n++ {
	//	t.Log(b[n])
	//}
	t.Log(b[3:])
	t.Log(b[2:len(b)])
}

/**
 * var s []int
 * var s1 = []int {1,2,3}
 * var s2 = make([]int ,3,5)
 * 切片是2倍在增长 所以 s= append[2,i]
 */
func TestSlice(t *testing.T)  {
	var s1= [] int{1,2,3}
	var s2 = make([]int, 3,4)
	t.Log(len(s1),cap(s1))
	t.Log(len(s2),cap(s2))
	t.Log(s1[0],s1[1],s1[2])
}

/**
 * map底层采用hash
 */
func TestMap(t *testing.T)  {
	var a = map[string] int {"test":1,"test1":2,"test2":3}
	for i,index :=range a{
		t.Log(i,index)

	}

}








