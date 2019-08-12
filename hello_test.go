package main

import (
	"errors"
	"fmt"
	"testing"
)

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
	var b= map[int] int{}
	b[0] =1
	b[1] =2
	t.Log(b)
	t.Log(Returntest())
}
func Returntest() (int, int){
	return 100,200
}
type student struct {
	name string
	age int
}
/**
 * go面向对象
 * 初始化
 * 1. student {1,2,3}
 * 2. student {k1:v1,k2:v2}
 * 3. new (student)
 * 4. & student{"test",1}
 */

func TestObject (t *testing.T)  {
	a := student{"fengwnewne",2}
	fmt.Println("a address is %x",&a.name)
	t.Log(a.string())
	t.Log(a.zhizhen())
	}

func (s student) string() string   {
	fmt.Println("address is %x",&s.name)
	return fmt.Sprintf("name is %s,age is %d",s.name,s.age)
}

func (s *student) zhizhen() string {
	fmt.Println("address is x%",&s.name)
	return fmt.Sprintf("name is %s,age is %d",s.name,s.age)
}

/**
 * go接口
 */
type Prog interface {
	WriteHelloWord() string
}
type GoProg struct {
	
}

func (Go *GoProg) WriteHelloWord() string  {
	return "hello world"
}

func TestInterface(t *testing.T)  {
	var b Prog
	b = new(GoProg)
	t.Log(b.WriteHelloWord())
}
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func TestCainaio(t *testing.T) {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone1 := new(IPhone)
	phone1.call()

}

type pet struct {

}

func (p *pet) Speak()  {
	fmt.Println("....")
}
func (p *pet) SpeakTo(host string)  {
	p.Speak()
	fmt.Println("", host)
}

type Dog struct {
	 p* pet
}

func (d *Dog)Speak()  {
	fmt.Println("wang")
}


func TestOxtend(t *testing.T)  {
	var d Dog
	d.Speak()
}

func TestError(t *testing.T)  {
	t.Log(errors.New("error"))
}







