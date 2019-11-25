# Gengine [中文说明](#基于golang的规则引擎)

## the rule engine based on golang 
- this is a rule engine named **Gengine** based on golang and AST, it can help you to load your code(rules) to run while you did not need to restart your application.  
- Gengine's code structure is Modular design, logic is easy to understand, and it passed the necessary testing！    

## Grammar support by Gengine
- support the priority of the rules list and the priority scope is -int64 to int64
- support rule's description
- support define a local variable in a rule, and it invisible between rules
- support 'if..else' and it's Nested structure
- support complex logic operationa
- support complex Arithmetic (+ -  * /)
- support method of golang's structure 
- support single line comment(//)
- support elegant check error, if there is an error in one rule, gengine will not load the rules to run to forbidden the harm to data 
- to make it easy to use,Gengine just supports one return value function's or method's Assignment, but support call multi return value function
- support directly inject function to run, and rename function
- support switch to help user to decide when a rule execute error in the list whether continue to execute the last rules 

## Gengine not support grammar
- not support 'else if', beacase the creator hates 'else if'
- not support Multi-level call such as 'user.ip.ipRisk()'，because it not meet the "Dimit rule", and multi-level call make it hard to understand, so it just support this call type: 'ip.ipRisk()'
- not support multi line comment (/*  comment */)
- not support nil 

## something need your attention
- if you not declare the rules' priority, the rules will be execute in unknown sort 
- every rule's name should be different from each other 


## support data type
- string 
- bool  
- int, int8, int16, int32, int64   
- float32, float64

## support logic operation
- &&  
- ||  
- !   

## support compared operation
- ==   
- !=   
- \>   
-  <   
-  \>=  
- <=   

## support math operation
-  \+ 
-  \- 
-  \* 
-  /  
- support string and string's plus


## Gengine rule example
```go
//rule
rule "测试" "测试描述"  salience 0 
begin
		// rename function test
		Sout("XXXXXXXXXX")
		// common function test
		Hello()
		//struct's method  test
		User.Say()
		// if
		if !(7 == User.GetNum(7)) || !(7 > 8)  {
			//define variable and string's plus
			variable = "hello" + (" world" + "zeze")
			// inner function
			User.Name = "hhh" + strconv.FormatBool(true)
			//struct's field 
			User.Age = User.GetNum(8976) / 1000+ 3*(1+1) 
			//bool set test
			User.Male = false
			//use inner variable test
			User.Print(variable)
			//float test		
			f = 9.56			
			PrintReal(f)
			//if-else test
			if false	{
				Sout("sout true")
			}else{
				Sout("sout false")
			}
		}else{ //else
			//struct field set value test
			User.Name = "yyyy"
		}
end
```

##  Gengine complete test 
- you can find all code in test package
```go

import (
	"fmt"
	"gengine/base"
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

type User struct {
	Name string
	Age  int
	Male bool
}

func (u *User)GetNum(i int64) int64 {
	return i
}

func (u *User)Print(s string){
	fmt.Println(s)
}

func (u *User)Say(){
	fmt.Println("hello world")
}

const (
	rule2 = `
rule "测试" "测试描述"  salience 0 
begin
		// 重命名函数 测试
		Sout("XXXXXXXXXX")
		// 普通函数 测试
		Hello()
		//结构提方法 测试
		User.Say()
		// if
		if 7 == User.GetNum(7){
			//自定义变量 和 加法 测试
			variable = "hello" + " world"
			// 加法 与 内建函数 测试
			User.Name = "hhh" + strconv.FormatBool(true)
			//结构体属性、方法调用 和 除法 测试
			User.Age = User.GetNum(89767999999) / 10000000
			//布尔值设置 测试
			User.Male = false
			//规则内自定义变量调用 测试
			User.Print(variable)
			//float测试	也支持科学计数法		
			f = 9.56			
			PrintReal(f)
			//嵌套if-else测试
			if false	{
				Sout("嵌套if测试")
			}else{
				Sout("嵌套else测试")
			}
		}else{ //else
			//字符串设置 测试
			User.Name = "yyyy"
		}
end`)

func Hello()  {
	fmt.Println("hello")
}

func PrintReal(real float64){
	fmt.Println(real)
}

func exe(user *User){
	/**
	 不要注入除函数和结构体指针以外的其他类型(如变量)
	 */
	dataContext := context.NewDataContext()
	//注入结构体指针
	dataContext.Add("User", user)
	//重命名函数,并注入
	dataContext.Add("Sout",fmt.Println)
	//直接注入函数
	dataContext.Add("Hello",Hello)
	dataContext.Add("PrintReal",PrintReal)

	//初始化规则引擎
	knowledgeContext := base.NewKnowledgeContext()
	ruleBuilder := builder.NewRuleBuilder(knowledgeContext, dataContext)

	//读取规则
	err := ruleBuilder.BuildRuleFromString(rule2)
	if err != nil{
		logrus.Errorf("err:%s ", err)
	}else{
		eng := engine.NewGengine()

		start := time.Now().UnixNano()
		// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
		err := eng.Execute(ruleBuilder, true)
		end := time.Now().UnixNano()
		if err != nil{
			logrus.Errorf("execute rule error: %v", err)
		}
		logrus.Infof("execute rule cost %d ns",end-start)
		logrus.Infof("user.Age=%d,Name=%s,Male=%t", user.Age, user.Name, user.Male)
	}
}

func Test_Base(t *testing.T){
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}
	exe(user)
}

```

## Connection
- renyunyi@bilibili.com

## Licence
- MIT



## 基于golang的规则引擎
- **Gengine**是一款基于AST(Abstract Syntax Tree)和golang语言实现的规则引擎。能够让你在golang这种静态语言上，在不停服务的情况下实现动态加载与配置规则。
- **代码结构松散，逻辑极其简单，但经过了必要且详尽的测试**
- Gengine所支持的规则，就是一门**DSL**(领域专用语言)

## Gengine支持的规则语法
- 支持规则优先级和规则执行条件，优先级高的先执行，优先级低的后执行；
- 支持的优先级范围 -int64 ～ int64 
- 支持中文规则名与中文规则描述
- 支持规则内定义变量，但规则内定义的变量在规则与规则之间的不可见
- 支持 if../if..else.. 代码块和其代码块嵌套
- 支持复杂逻辑运算
- 支持复杂数学四则运算(+ -  * /)
- 支持结构体方法调用
- 支持单行注释(//)  
- 支持优雅的规则检错机制(是的,你没看错,就是检错,不是检测!):如果待加载的一批规则中有一个规则有语法错误，那么规则引擎就不会加载这批规则去执行,防止对数据造成不可预知的危害
- 支持仅有一个返回值的函数赋值，且返回值为基础类型；支持多返回值函数的调用，但无法处理其返回值
- 支持直接注入函数并执行，并允许函数重命名
- 支持规则链中有规则执行失败时，是否继续执行后续规则开关
- 支持一些内置函数

## Gengine不支持的规则语法
- 不支持else if, 因为作者讨厌使用else if
- 不支持形如user.ip.ipRisk()这种多级调用，因为它不符合"迪米特法则"，并且会使代码变得难以理解；只支持ip.ipRisk()这种单级调用
- 不支持函数多个返回值，并且不支持返回类型为非基础类型
- 不支持多行注释，因为不想支持
- 不支持nil 

## 书写规则需要注意的事情
- 如果规则的优先级不指定，多个规则将以未知次序执行
- 同一批待加载的规则中不能有重名规则

## 支持的基本数据类型
- string 
- bool  
- int, int8, int16, int32, int64   
- float32, float64

## 支持的逻辑运算符
- &&  且 
- ||  或
- !   非

## 支持的比较运算符
- ==   等于
- !=   不等于
- \>   大于
-  <   小于
-  \>= 大于等于 
- <=   小于等于

## 支持的算术运算符
-  \+ 加
-  \- 减
-  \* 乘
-  /  除
- 支持int,uint,float任意两者之间的加减乘除, 以及string与string之间的加法

## 支持的括号
- 圆括号

## 运算优先级
- 单目运算符非(!) > 算术运算符 > 比较运算符 > 逻辑运算符

## Gengine规则示例
```go
//规则
rule "测试" "测试描述"  salience 0 
begin
		// 重命名函数 测试
		Sout("XXXXXXXXXX")
		// 普通函数 测试
		Hello()
		//结构提方法 测试
		User.Say()
		// if
		if !(7 == User.GetNum(7)) || !(7 > 8)  {
			//自定义变量 和 加法 测试
			variable = "hello" + (" world" + "zeze")
			// 加法 与 内建函数 测试
			User.Name = "hhh" + strconv.FormatBool(true)
			//结构体属性、方法调用 和 除法 测试
			User.Age = User.GetNum(8976) / 1000+ 3*(1+1) 
			//布尔值设置 测试
			User.Male = false
			//规则内自定义变量调用 测试
			User.Print(variable)
			//float测试	也支持科学计数法		
			f = 9.56			
			PrintReal(f)
			//嵌套if-else测试
			if false	{
				Sout("嵌套if测试")
			}else{
				Sout("嵌套else测试")
			}
		}else{ //else
			//字符串设置 测试
			User.Name = "yyyy"
		}
end
```

## Gengine完整的规则加载并执行的代码示例
- 所有代码，在test测试包可见
```go

import (
	"fmt"
	"gengine/base"
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

type User struct {
	Name string
	Age  int
	Male bool
}

func (u *User)GetNum(i int64) int64 {
	return i
}

func (u *User)Print(s string){
	fmt.Println(s)
}

func (u *User)Say(){
	fmt.Println("hello world")
}

const (
	rule2 = `
rule "测试" "测试描述"  salience 0 
begin
		// 重命名函数 测试
		Sout("XXXXXXXXXX")
		// 普通函数 测试
		Hello()
		//结构提方法 测试
		User.Say()
		// if
		if 7 == User.GetNum(7){
			//自定义变量 和 加法 测试
			variable = "hello" + " world"
			// 加法 与 内建函数 测试
			User.Name = "hhh" + strconv.FormatBool(true)
			//结构体属性、方法调用 和 除法 测试
			User.Age = User.GetNum(89767999999) / 10000000
			//布尔值设置 测试
			User.Male = false
			//规则内自定义变量调用 测试
			User.Print(variable)
			//float测试	也支持科学计数法		
			f = 9.56			
			PrintReal(f)
			//嵌套if-else测试
			if false	{
				Sout("嵌套if测试")
			}else{
				Sout("嵌套else测试")
			}
		}else{ //else
			//字符串设置 测试
			User.Name = "yyyy"
		}
end`)

func Hello()  {
	fmt.Println("hello")
}

func PrintReal(real float64){
	fmt.Println(real)
}

func exe(user *User){
	/**
	 不要注入除函数和结构体指针以外的其他类型(如变量)
	 */
	dataContext := context.NewDataContext()
	//注入结构体指针
	dataContext.Add("User", user)
	//重命名函数,并注入
	dataContext.Add("Sout",fmt.Println)
	//直接注入函数
	dataContext.Add("Hello",Hello)
	dataContext.Add("PrintReal",PrintReal)

	//初始化规则引擎
	knowledgeContext := base.NewKnowledgeContext()
	ruleBuilder := builder.NewRuleBuilder(knowledgeContext, dataContext)

	//读取规则
	err := ruleBuilder.BuildRuleFromString(rule2)
	if err != nil{
		logrus.Errorf("err:%s ", err)
	}else{
		eng := engine.NewGengine()

		start := time.Now().UnixNano()
		// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
		err := eng.Execute(ruleBuilder, true)
		end := time.Now().UnixNano()
		if err != nil{
			logrus.Errorf("execute rule error: %v", err)
		}
		logrus.Infof("execute rule cost %d ns",end-start)
		logrus.Infof("user.Age=%d,Name=%s,Male=%t", user.Age, user.Name, user.Male)
	}
}

func Test_Base(t *testing.T){
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}
	exe(user)
}

```

## 联系
- renyunyi@bilibili.com


## 许可证
- MIT

