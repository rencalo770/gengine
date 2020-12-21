<div align="center">
  <img src="gengine.png">
</div>

[![license](https://img.shields.io/badge/license-BSD-blue.svg)]()
[![Documentation](https://img.shields.io/badge/api-reference-blue.svg)](https://rencalo770.github.io/gengine_en) 

# Gengine
- [简体中文](README_zh.md) 使用交流Q群1132683357

## the rule engine based on golang 
- this is a rule engine(or code dynamic load framework) named **Gengine** based on golang and AST, it can help you to load your code(rules) to run while you did not need to restart your application.  
- Gengine's code structure is Modular design, logic is easy to understand, and necessary testing！
- it is also a high performance engine, support many execute-model and rules pool for business, it is easy to use in distribute system. 

## English Doc
- because gengine is designed to be extremely easy to use(only 4 APIs), so there are a lot programmers to use gengine without reading document!
- I sincerely hope you can read the document before you use gengine! The doc can greatly help you to use it well
- https://rencalo770.github.io/gengine_en

## supported the execute model of rules
 ![avatar](exe_model.jpg)

## use 
- go mod or go vendor 
- use the newest version! 

## use pool
- for user to use the gengine pool, we made the methods one to one between gengine.go and gengine_pool.go 

| methods in ```gengine.go```(up) <br/> methods in ```gengine_pool.go```(down)| method explain | 
| -------- | -------- | 
|```func (g *Gengine) Execute(rb *builder.RuleBuilder, b bool) error```<br/> ****``` func (gp *GenginePool) Execute(data map[string]interface{}, b bool) (error, map[string]interface{}) ```**** | Sequential execution mode, <br/>b=true means that even if a rule fails, the remaining rules will continue to be executed; <br/>if b=false means that if a rule fails, the subsequent rules will not be executed |
|```func (g *Gengine) ExecuteWithStopTagDirect(rb *builder.RuleBuilder, b bool, sTag *Stag) error ``` <br/> ```func (gp *GenginePool)ExecuteWithStopTagDirect(data map[string]interface{}, b bool, sTag *Stag) (error, map[string]interface{}) ```|All rules are executed in order.<br/>b=true means that even if a rule fails, the remaining rules will continue to be executed; <br/>if b=false means that if a rule fails, the following rules will not be executed; <br/>If sTag=true, then do not continue to execute the following rules, if sTag=false, continue to execute the following rules|
|```func (g *Gengine) ExecuteConcurrent(rb *builder.RuleBuilder) error  ``` <br/> ``` func (gp *GenginePool)ExecuteConcurrent(data map[string]interface{})  (error, map[string]interface{}) ```|Concurrent mode, execute all rules concurrently|
|```func (g *Gengine) ExecuteMixModel(rb *builder.RuleBuilder) error ``` <br/> ```func  (gp *GenginePool)ExecuteMixModel(data map[string]interface{})  (error, map[string]interface{})  ```|Mixed mode, first execute a rule with the highest priority, and then execute all remaining rules concurrently|
|```func (g *Gengine) ExecuteMixModelWithStopTagDirect(rb *builder.RuleBuilder, sTag *Stag) error  ``` <br/> ```func  (gp *GenginePool) ExecuteMixModelWithStopTagDirect(data map[string]interface{}, sTag *Stag)  (error, map[string]interface{}) ```|Mixed mode, <br/>first execute a rule with the highest priority, and then execute all the remaining rules concurrently; <br/>if sTag is set to true in the first executed rule, the rest will not be executed the rule of|
|```func (g *Gengine) ExecuteSelectedRules(rb *builder.RuleBuilder, names []string) error  ``` <br/> ``` func (gp *GenginePool) ExecuteSelectedRules(data map[string]interface{}, names []string)  (error, map[string]interface{}) ```|Select a batch of rules based on the rule name and execute them in sequential mode.<br/>Even if there is a rule execution error during the execution process, it will continue to execute all the remaining rules|
|```func (g *Gengine) ExecuteSelectedRulesWithControl(rb *builder.RuleBuilder, b bool, names []string) error ``` <br/> ```func (gp *GenginePool) ExecuteSelectedRulesWithControl(data map[string]interface{}, b bool, names []string)  (error, map[string]interface{}) ```|Select a batch of rules based on the rule name and execute them in sequential mode.<br/>Even if there is a rule execution error during the execution, all the remaining rules will continue to be executed; <br/>b=true means that even if a rule goes wrong, Continue to execute the remaining rules; <br/>If b=false means that if a rule fails, then the following rules will not be executed|
|```func (g *Gengine) ExecuteSelectedRulesWithControlAndStopTag(rb *builder.RuleBuilder, b bool, sTag *Stag, names []string) error ``` <br/> ```func  (gp *GenginePool)ExecuteSelectedRulesWithControlAndStopTag(data map[string]interface{}, b bool, sTag *Stag, names []string)  (error, map[string]interface{}) ```|Select a batch of rules based on the rule name and execute them in sequential mode.<br/>Even if there is a rule execution error during the execution process, all the remaining rules will continue to be executed; <br/>b=true means that even if a rule goes wrong, the Continue to execute the remaining rules; <br/> if b=false means that if a rule fails, the following rules will not continue; <br/>If sTag is set to true in the first executed rule, the remaining rules will not be executed Rules under|
|```func (g *Gengine) ExecuteSelectedRulesConcurrent(rb *builder.RuleBuilder, names []string) error  ``` <br/> ```func (gp *GenginePool) ExecuteSelectedRulesConcurrent(data map[string]interface{}, names []string) (error, map[string]interface{}) ```|Select a batch of rules based on the rule name, and execute the selected rules concurrently|
|```func (g *Gengine) ExecuteSelectedRulesMixModel(rb *builder.RuleBuilder, names []string) error ``` <br/> ```func (gp *GenginePool) ExecuteSelectedRulesMixModel(data map[string]interface{}, names []string)  (error, map[string]interface{}) ```|Select a batch of rules based on the rule name, and execute the selected rules in mixed mode|
|```func (g *Gengine) ExecuteInverseMixModel(rb *builder.RuleBuilder) error ```  <br/> ``` func (gp *GenginePool) ExecuteInverseMixModel(data map[string]interface{})  (error, map[string]interface{})  ```  |Inverse mixed execution mode, execute all rules|
|```func (g *Gengine) ExecuteSelectedRulesInverseMixModel(rb *builder.RuleBuilder, names []string) error ``` <br/> ``` func (gp *GenginePool)ExecuteSelectedRulesInverseMixModel(data map[string]interface{}, names []string) (error, map[string]interface{})  ```|基于规则名选择一批规则,逆混合执行模式，执行选中的规则|
|```func (g *Gengine) ExecuteNSortMConcurrent(nSort, mConcurrent int, rb *builder.RuleBuilder, b bool) error  ``` <br/> ```func (gp *GenginePool) ExecuteNSortMConcurrent(nSort, mConcurrent int, b bool, data map[string]interface{}) (error, map[string]interface{})  ```|N-M模式,前N个规则顺序执行，接下来的M个规则并发执行, 且N+M大于等于规则集的长度|
|```func (g *Gengine) ExecuteNConcurrentMSort(nConcurrent, mSort int, rb *builder.RuleBuilder, b bool) error  ``` <br/> ```func (gp *GenginePool) ExecuteNConcurrentMSort(nSort, mConcurrent int, b bool, data map[string]interface{}) (error, map[string]interface{}) ```|N-M模式,前N个规则并发执行，接下来的M个规则顺序执行, 且N+M大于等于规则集的长度|
|```func (g *Gengine) ExecuteNConcurrentMConcurrent(nConcurrent, mConcurrent int, rb *builder.RuleBuilder, b bool) error  ```<br/> ```func (gp *GenginePool) ExecuteNConcurrentMConcurrent(nSort, mConcurrent int, b bool, data map[string]interface{}) (error, map[string]interface{})  ```|N-M模式,前N个规则并发执行，接下来的M个规则也并发执行, 且N+M大于等于规则集的长度|
|```func (g *Gengine) ExecuteSelectedNSortMConcurrent(nSort, mConcurrent int, rb *builder.RuleBuilder, b bool, names []string) error  ```<br/> ```func (gp *GenginePool) ExecuteSelectedNSortMConcurrent(nSort, mConcurrent int, b bool, names []string, data map[string]interface{}) (error, map[string]interface{}) ```|基于规则名选择一批规则,N-M模式,前N个规则顺序执行，接下来的M个规则并发执行, 且N+M大于等于选定的规则集的长度 |
|```func (g *Gengine) ExecuteSelectedNConcurrentMSort(nConcurrent, mSort int, rb *builder.RuleBuilder, b bool, names []string) error ``` <br/> ```func (gp *GenginePool) ExecuteSelectedNConcurrentMSort(nSort, mConcurrent int, b bool, names []string, data map[string]interface{}) (error, map[string]interface{})  ```  |基于规则名选择一批规则, N-M模式,前N个规则并发执行，接下来的M个规则顺序执行, 且N+M大于等于选定的规则集的长度|
|```func (g *Gengine) ExecuteSelectedNConcurrentMConcurrent(nConcurrent, mConcurrent int, rb *builder.RuleBuilder, b bool, names []string) error  ```<br/>```func (gp *GenginePool) ExecuteSelectedNConcurrentMConcurrent(nSort, mConcurrent int, b bool, names []string, data map[string]interface{}) (error, map[string]interface{})  ```|基于规则名选择一批规则,N-M模式,前N个规则并发执行，接下来的M个规则也并发执行, 且N+M大于等于选定的规则集的长度|






## Question Connection
- write issue or connect
- renyunyi@bilibili.com (become some reason,this mail box may can't get mail)
- M201476117@alumni.hust.edu.cn
