package main

import (
	"errors"
	"fmt"
)

type chouma struct {
	Cchouma int//当前筹码
	Cyingli int//当前盈利
	jiechou int//借了几手
}

func (cm *chouma) change(i int ) (err error)   {
	if cm.Cchouma < i{ err=errors.New("当前筹码不足");return err}
	cm.Cchouma=cm.Cchouma+i
	cm.jiechou=cm.jiechou
	cm.Cyingli=cm.Cyingli+i
	err=nil;return err
}

func (cm *chouma) jieyishou( ) ()   {
	cm.Cchouma=cm.Cchouma+1000
	cm.jiechou=cm.jiechou+1000
}
func  initchouma()  ( chouma){
	return chouma{Cchouma: 1000,
					Cyingli: 0,
					jiechou: 1000,
	}

}
func main()  {
a:=initchouma()
a.change(100)
a.jieyishou()
fmt.Println(a)

}
