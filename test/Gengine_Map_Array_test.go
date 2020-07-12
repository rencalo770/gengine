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


type MapArray struct {
	Mx map[string]bool
	Ax [3]int
	Sx []string
}

const ma_rule  =`
rule "测试规则" "rule desc"
begin
  
  x = Ma.Mx["hello"]
  
  PrintName(x)
  
  Ma.Mx["hello"] = false
  b = "your"
  
  Ma.Mx[b]= true
  
  y = Ma.Mx	
  PrintName("------",y["hello"])

  if x {
     PrintName("Single data")
  }	

  if 2 == 2 {
     PrintName("true == true")
  }
   
  if x == true {
   PrintName("haha")
  } 

  if !x {
     PrintName("haha")
  }else{
     PrintName("!x")
  }
  
  xx = Ma.Ax[2]
  PrintName(xx) 
  //Ma.Ax[2] = 3000001
 
  yy = Ma.Ax
  PrintName(yy[1]) 
  
  
  if yy[2] == 20000 {
     PrintName("20000")
  }

  z = Ma.Sx[1]
  PrintName(z) 

  zz = Ma.Sx
  if zz[2] == "kkkk"{
     PrintName("kkkk") 
  }

  a = 2
  zz[a] = "MMMM"
  PrintName(zz[a]) 

end
`


func Test_map_array(t *testing.T) {

	Ma := &MapArray{
		Mx: map[string]bool{"hello": true},
		Ax : [3]int{1000,20000,300},
		Sx: []string{"jjj", "lll", "kkkk"},
	}


	dataContext := context.NewDataContext()
	dataContext.Add("PrintName",fmt.Println)
	dataContext.Add("Ma", Ma)

	//init rule engine
	knowledgeContext := base.NewKnowledgeContext()
	ruleBuilder := builder.NewRuleBuilder(knowledgeContext, dataContext)

	//读取规则
	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(ma_rule)
	end1 := time.Now().UnixNano()

	logrus.Infof("rules num:%d, load rules cost time:%d ns", len(knowledgeContext.RuleEntities), end1-start1 )


	if err != nil{
		logrus.Errorf("err:%s ", err)
	}else{
		eng := engine.NewGengine()
		start := time.Now().UnixNano()
		// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
		err := eng.Execute(ruleBuilder, true)

		println("----------", Ma.Mx["hello"])
		println("----------", Ma.Mx["your"])
		println("++++++++++", Ma.Ax[2])
		println("==========", Ma.Sx[2])

		end := time.Now().UnixNano()
		if err != nil{
			logrus.Errorf("execute rule error: %v", err)
		}
		logrus.Infof("execute rule cost %d ns",end-start)
	}


}
