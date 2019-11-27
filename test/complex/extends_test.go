package complex

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

const ext_rule = `
rule "extends test" "extends test" 
begin
	Father.Son = "tom"
	Sout(Father.Son)
end
`

func exe(father *Father){

	dataContext := context.NewDataContext()
	//注入结构体指针
	dataContext.Add("Father", father)
	//重命名函数,并注入
	dataContext.Add("Sout",fmt.Println)

	//初始化规则引擎
	knowledgeContext := base.NewKnowledgeContext()
	ruleBuilder := builder.NewRuleBuilder(knowledgeContext, dataContext)

	//读取规则
	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(ext_rule)
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


func Test_ext(t *testing.T) {
	father := &Father{}
	exe(father)
}
