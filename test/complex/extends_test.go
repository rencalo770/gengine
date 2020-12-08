package complex

import (
	"fmt"
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
	"github.com/google/martian/v3/log"

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

func exe(father *Father) {

	dataContext := context.NewDataContext()
	//inject struct
	dataContext.Add("Father", father)
	//rename and inject
	dataContext.Add("Sout", fmt.Println)

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(ext_rule)
	end1 := time.Now().UnixNano()

	log.Infof("rules num:%d, load rules cost time:%d ns", len(ruleBuilder.Kc.RuleEntities), end1-start1)

	if err != nil {
		log.Errorf("err:%s ", err)
	} else {
		eng := engine.NewGengine()

		start := time.Now().UnixNano()
		// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
		err := eng.Execute(ruleBuilder, true)
		end := time.Now().UnixNano()
		if err != nil {
			log.Errorf("execute rule error: %v", err)
		}
		log.Infof("execute rule cost %d ns", end-start)
	}
}

func Test_ext(t *testing.T) {
	father := &Father{
		Man: &Man{},
	}
	exe(father)
}
