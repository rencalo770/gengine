package builder

import (
	"gengine/base"
	"gengine/context"
	"gengine/core/errors"
	parser "gengine/iantlr/alr"
	"gengine/iparser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"sort"
)


type RuleBuilder struct {
	Kc *base.KnowledgeContext
	Dc *context.DataContext
}

func NewRuleBuilder(kc *base.KnowledgeContext,dc *context.DataContext)*RuleBuilder{
	return &RuleBuilder{
		Kc: kc,
		Dc: dc,
	}
}

func (builder *RuleBuilder) BuildRuleFromString(ruleString string) error{
	//forbidden old rules in context
	builder.Kc.ClearRules()

	in := antlr.NewInputStream(ruleString)
	lexer := parser.NewgengineLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	listener := iparser.NewGengineParserListener(builder.Kc)

	psr := parser.NewgengineParser(stream)
	psr.BuildParseTrees = true
	//语法错误监听器
	errListener := iparser.NewGengineErrorListener()
	psr.AddErrorListener(errListener)
	antlr.ParseTreeWalkerDefault.Walk(listener, psr.Primary())

	if  len(errListener.GrammarErrors) > 0 {
		builder.Kc.ClearRules()
		return errors.Errorf("%v", errListener.GrammarErrors)
	}

	if len(listener.ParseErrors) > 0 {
		builder.Kc.ClearRules()
		return errors.Errorf("%v", listener.ParseErrors)
	}

	//initial
	for _,v := range builder.Kc.RuleEntities {
		v.Initialize(builder.Kc, builder.Dc)
	}

	//sort
	for _,v := range builder.Kc.RuleEntities {
		builder.Kc.SortRules = append(builder.Kc.SortRules, v)
	}
	if len(builder.Kc.SortRules) > 1 {
		sort.SliceStable(builder.Kc.SortRules, func(i, j int) bool {
			return builder.Kc.SortRules[i].Salience > builder.Kc.SortRules[j].Salience
		})
	}
	return nil
}