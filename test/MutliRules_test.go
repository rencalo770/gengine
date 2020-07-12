package test

import (
	"bytes"
	"fmt"
	"gengine/base"
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
	"time"
)

const rule1 = `
rule "name test" "i can"  salience 0
begin
		if 7 == User.GetNum(7){
			User.Age = User.GetNum(89767) + 10000000
			User.Print("6666")
		}else{
			User.Name = "yyyy"
		}
end
rule "姓名测试1" "我可以的"  salience 0
begin
		if 7 == User.GetNum(7){
			User.Age = User.GetNum(89767) + 100000000
			User.Print("6666")
		}else{
			User.Name = "yyyy"
		}
end
`

func Test_Multi(t *testing.T){
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}

	dataContext := context.NewDataContext()
	dataContext.Add("User", user)

	//init rule engine
	knowledgeContext := base.NewKnowledgeContext()
	ruleBuilder := builder.NewRuleBuilder(knowledgeContext,dataContext)


	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(rule1) //string(bs)
	end1 := time.Now().UnixNano()

	logrus.Infof("rules num:%d, load rules cost time:%d", len(knowledgeContext.RuleEntities), end1-start1 )

	if err != nil{
		logrus.Errorf("err:%s ", err)
	}else{
		eng := engine.NewGengine()

		start := time.Now().UnixNano()
		err := eng.Execute(ruleBuilder,true)
		println(user.Age)
		end := time.Now().UnixNano()
		if err != nil{
			logrus.Errorf("execute rule error: %v", err)
		}
		logrus.Infof("execute rule cost %d ns",end-start)
		logrus.Infof("user.Age=%d,Name=%s,Male=%t", user.Age, user.Name, user.Male)
	}
}

func Test_Read(t*testing.T){

	bytes, e := ioutil.ReadFile("/path/to/file")
	if e != nil {
		panic(e)
	}
	fmt.Println(string(bytes))

}

func Test_Write(t *testing.T){

	r :=`rule "TTTTTTT" "我可以的"  salience 0
		begin
		if 7 == User.GetNum(7){
			User.Age = User.GetNum(89767) + 10000000
			User.Print("6666")
		}else{
			User.Name = "yyyy"
		}
	    end
`

	var buffer bytes.Buffer
	for i := 0;i < 100; i++ {
		i2 := strconv.Itoa(i)
		rep := strings.Replace(r, "TTTTTTT", "姓名测试"+i2, -1)
		buffer.WriteString(rep)
	}

	cont := buffer.String()


	fileName := "/Users/renyunyi/go/src/gengine/test/file"
	err := ioutil.WriteFile(fileName, []byte(cont), 0664)
	if err != nil {
		panic(err)
	}
	fmt.Println("写入文件成功!")
}
