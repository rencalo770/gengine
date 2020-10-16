package server

import (
	"fmt"
	"gengine/engine"
	"testing"
)

//业务规则
const service_rules string = `
rule "1" "1"
begin
	resp.At = room.GetAttention()
	println("rule 1...")
end 

rule "2" "2"
begin
	resp.Num = room.GetNum()
	println("rule 2...")
end
`

//业务接口
type MyService struct {
	//gengine pool
	Pool *engine.GenginePool

	//other params
}

//request
type Request struct {
	Rid       int64
	RuleNames []string
	//other params
}

//resp
type Response struct {
	At  int64
	Num int64
	//other params
}

//特定的场景服务
type Room struct {
}

func (r *Room) GetAttention( /*params*/ ) int64 {
	// logic
	return 100
}

func (r *Room) GetNum( /*params*/ ) int64 {
	//logic
	return 111
}

//初始化业务服务
func NewMyService(poolMinLen, poolMaxLen int64, em int, rulesStr string, apiOuter map[string]interface{}) *MyService {
	pool, e := engine.NewGenginePool(poolMinLen, poolMaxLen, em, rulesStr, apiOuter)
	if e != nil {
		panic(fmt.Sprintf("初始化gengine失败，err:%+v", e))
	}

	myService := &MyService{Pool: pool}
	return myService
}

//service
func (ms *MyService) Service(req *Request) (*Response, error) {

	resp := &Response{}

	//基于需要注入接口或数据
	data := make(map[string]interface{})
	data["req"] = req
	data["resp"] = resp

	//模块化业务逻辑,api
	room := &Room{}
	data["room"] = room

	e := ms.Pool.ExecuteSelectedRulesWithMultiInput(data, req.RuleNames)
	if e != nil {
		println(fmt.Sprintf("pool execute rules error: %+v", e))
		return nil, e
	}

	return resp, nil
}

//模拟调用
func Test_run(t *testing.T) {

	//初始化
	//注入api，请确保注入的API属于并发安全
	apis := make(map[string]interface{})
	apis["println"] = fmt.Println
	msr := NewMyService(10, 20, 1, service_rules, apis)


	//调用
	req := &Request{
		Rid:       123,
		RuleNames: []string{"1", "2"},
	}
	response, e := msr.Service(req)
	if e != nil {
		println(fmt.Sprintf("service err:%+v", e))
		return
	}

	println("resp result = ", response.At, response.Num)
}

