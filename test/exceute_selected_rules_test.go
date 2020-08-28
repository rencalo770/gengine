package test

import (
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

const rule5 = `
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
	Print("333")	
END
rule "444" "444" salience 11
BEGIN
	Print("444")
END
`

func Test_selected_sort(t *testing.T){
	dataContext := context.NewDataContext()
	dataContext.Add("Print", Print)

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(rule5)
	end1 := time.Now().UnixNano()

	logrus.Infof("rules num:%d, load rules cost time:%d", len(ruleBuilder.Kc.RuleEntities), end1-start1 )

	if err != nil{
		logrus.Errorf("err:%s ", err)
	}else{
		eng := engine.NewGengine()
		start := time.Now().UnixNano()
		eng.ExecuteSelectedRules(ruleBuilder, []string{"444","555","111","222", "333"})
		end := time.Now().UnixNano()
		logrus.Infof("execute rule cost %d ns",end-start)
	}
}

func Test_selected_concurrent(t *testing.T){
	dataContext := context.NewDataContext()
	dataContext.Add("Print", Print)

	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(rule5)
	end1 := time.Now().UnixNano()

	logrus.Infof("rules num:%d, load rules cost time:%d", len(ruleBuilder.Kc.RuleEntities), end1-start1 )

	if err != nil{
		logrus.Errorf("err:%s ", err)
	}else{
		eng := engine.NewGengine()
		start := time.Now().UnixNano()
		eng.ExecuteSelectedRulesConcurrent(ruleBuilder, []string{"444","111","222","777"})
		end := time.Now().UnixNano()
		logrus.Infof("execute rule cost %d ns",end-start)
	}
}












