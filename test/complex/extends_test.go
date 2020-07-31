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
	Father.Eat= "apple"
	Sout(Father.Eat)
end
`

func exe(father *Father){

	dataContext := context.NewDataContext()
	//inject struct
	dataContext.Add("Father", father)
	//rename and inject
	dataContext.Add("Sout",fmt.Println)

	//init rule engine
	knowledgeContext := base.NewKnowledgeContext()
	ruleBuilder := builder.NewRuleBuilder(knowledgeContext, dataContext)

	//读取规则
	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(ext_rule)
	end1 := time.Now().UnixNano()

	logrus.Infof("rules num:%d, load rules cost time:%d ns", len(knowledgeContext.RuleEntities), end1-start1 )

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
	father := &Father{
		Man:&Man{},
	}
	exe(father)
}
