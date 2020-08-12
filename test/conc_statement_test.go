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

const rule_conc_statement  = `
rule "conc_test" "test" 
begin
	conc  { // this should be used in time-consuming operation, such as the operation contains network connection (get data from remote based on network) 
		println("hihihi")
		a = 3.0
		b = 4
		c = 6.8
		d = "hello world"
        e = "you will be happy here!"
		sout("heheheheh")
	}
	println(a, b, c, d, e)

end
`

func Sout(str string)  {
	println("----",str)
}

func Test_conc_statement(t *testing.T) {

	dataContext := context.NewDataContext()
	dataContext.Add("println",fmt.Println)
	dataContext.Add("sout", Sout)

	//init rule engine
	knowledgeContext := base.NewKnowledgeContext()
	ruleBuilder := builder.NewRuleBuilder(knowledgeContext, dataContext)

	//resolve rules from string
	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(rule_conc_statement)
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

