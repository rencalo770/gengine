package test

import (
	"fmt"
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
	"reflect"
	"testing"
)

func Test_hello(t *testing.T) {

	type Result struct {
		RiskLevel string
	}

	type CTX struct {
		Result Result
	}

	result := Result{
		RiskLevel:"hello",
	}
	ctx := &CTX{Result:result}

	dataContext := context.NewDataContext()
	dataContext.Add("ctx", ctx)
	dataContext.Add("println", fmt.Println)

	ruleBuilder := builder.NewRuleBuilder(dataContext)
	e1 := ruleBuilder.BuildRuleFromString(`
rule "1" 
begin
result = ctx.Result
//result.RiskLevel = "E"
end
`)
	if e1 != nil {
		panic(e1)
	}

	gengine := engine.NewGengine()
	e := gengine.Execute(ruleBuilder, true)
	if e != nil {
		panic(e)
	}

	println("---x->",ctx.Result.RiskLevel)



}

func Test_in_in(t *testing.T)  {
	type Result struct {
		RiskLevel string
	}

	type CTX struct {
		Result Result
	}

	result := Result{
		RiskLevel:"hello",
	}
	ctx := &CTX{Result:result}
	reflect.ValueOf(ctx).Elem().FieldByName("Result").FieldByName("RiskLevel").SetString("yyyyyyy")
	println("---x->",ctx.Result.RiskLevel)


}