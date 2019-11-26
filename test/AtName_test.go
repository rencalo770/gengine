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

func PrintName(name string){
	fmt.Println(name)
}

/**
use '@name',you can get rule name in rule content
 */
const atname_rule  =`
rule "测试规则名称1" "规则描述"
begin
  va = @name
  PrintName(va)
  PrintName(@name)
end
rule "测试规则名称2" "规则描述"
begin
  va = @name
  PrintName(va)
  PrintName(@name)
end
`

func exec(){
	dataContext := context.NewDataContext()
	dataContext.Add("PrintName",PrintName)

	//初始化规则引擎
	knowledgeContext := base.NewKnowledgeContext()
	ruleBuilder := builder.NewRuleBuilder(knowledgeContext, dataContext)

	//读取规则
	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(atname_rule)
	end1 := time.Now().UnixNano()

	logrus.Infof("规则个数:%d, 加载规则耗时:%d ns", len(knowledgeContext.RuleEntities), end1-start1 )

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
	}
}

func Test_AtName(t *testing.T){
	exec()
}





