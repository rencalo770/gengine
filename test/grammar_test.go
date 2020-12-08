package test

import (
	"fmt"
	"gengine/builder"
	"gengine/context"
	"gengine/internal/base"
	parser "gengine/internal/iantlr/alr"
	"gengine/internal/iparser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/google/martian/v3/log"

	"testing"
)

const (
	rule3 = `
rule "name test" "i can" salience 0
BEGIN
		if 7 == User.GetNum(7){
			User.Age = User.GetNum(89767) + 10000000
			User.Print("6666")
		}else{
			User.Name = "yyyy"
		}
END`
)

func Test_grammar(t *testing.T) {
	in := antlr.NewInputStream(rule3)
	lexer := parser.NewgengineLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	knowledgeBase := base.NewKnowledgeContext()
	listener := iparser.NewGengineParserListener(knowledgeBase)

	psr := parser.NewgengineParser(stream)
	psr.BuildParseTrees = true
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Primary())
	if len(listener.ParseErrors) > 0 {
		fmt.Println(listener.ParseErrors)
	}
}

func Test_base_msg(t *testing.T) {

	dataContext := context.NewDataContext()
	//init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	//读取规则
	err := ruleBuilder.BuildRuleFromString(rule3)
	if err != nil {
		log.Errorf("______%v", err)
	}

	rule := ruleBuilder.Kc.RuleEntities
	for k, v := range rule {
		fmt.Println(k)
		fmt.Println(v.RuleName)
		fmt.Println(v.RuleDescription)
		fmt.Println(v.Salience)
		fmt.Println(v.RuleContent)
	}
}

/**
测试语法错误
*/
func Test_err(t *testing.T) {

	in := antlr.NewInputStream(rule3)
	lexer := parser.NewgengineLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	knowledgeBase := base.NewKnowledgeContext()
	listener := iparser.NewGengineParserListener(knowledgeBase)

	psr := parser.NewgengineParser(stream)
	psr.BuildParseTrees = true

	//test -- here
	errListener := iparser.NewGengineErrorListener()
	psr.AddErrorListener(errListener)
	//test -- here

	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Primary())
	fmt.Println(errListener.GrammarErrors)
}
