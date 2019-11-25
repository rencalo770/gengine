package test

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

const rule4 = `
rule "111" "111" salience 80
BEGIN
		Print("111")
END
rule "222" "222" salience 10
BEGIN
		Print("222")
END
rule "333" "333" salience 11
BEGIN
	if false {
		Print("333")
	}else{
		Print("444")
	}
END

`

func Print(s string){
	fmt.Println(s)
}

func Test_Priority(t *testing.T){
	dataContext := context.NewDataContext()
	dataContext.Add("Print", Print)

	//初始化规则引擎
	knowledgeContext := base.NewKnowledgeContext()
	ruleBuilder := builder.NewRuleBuilder(knowledgeContext,dataContext)

	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(rule4)
	end1 := time.Now().UnixNano()

	logrus.Infof("规则个数:%d, 加载规则耗时:%d", len(knowledgeContext.RuleEntities), end1-start1 )

	if err != nil{
		logrus.Errorf("err:%s ", err)
	}else{
		eng := engine.NewGengine()
		start := time.Now().UnixNano()
		err := eng.Execute(ruleBuilder, true)
		end := time.Now().UnixNano()
		if err != nil{
			logrus.Errorf("execute rule error: %v", err)
		}
		logrus.Infof("execute rule cost %d ns",end-start)
	}
}





