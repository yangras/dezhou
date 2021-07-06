package main

import (
	"fmt"
	"testing"
)

func Test_tonghuashun(t *testing.T){
	dipai := []poker{{heitao, 14}, {heitao, 2}, {heitao, 3}, {heitao, 12}, {heitao, 13}}
	shoupai1 :=[]poker{{heitao, 4}, {heitao, 5}}
	shoupai2 :=[]poker{{heitao,10},{heitao,11}}
	shoupai3 :=[]poker{{heitao,8},{heitao,9}}

	zongpai1 := append(dipai, shoupai1...)
	zongpai2 := append(dipai, shoupai2...)
	zongpai3 := append(dipai, shoupai3...)
	fmt.Println("同花顺的检查结果如下")
	fmt.Println(defeat(zongpai1,zongpai2))
	fmt.Println(defeat(zongpai1,zongpai3))
}


func Test_sitiao(t *testing.T){
	dipai := []poker{{heitao, 14}, {hongtao, 14}, {heitao, 14}, {heitao, 10}, {heitao, 13}}
	shoupai1 :=[]poker{{heitao, 14}, {heitao, 13}}
	shoupai2 :=[]poker{{heitao,14},{heitao,11}}
	shoupai3 :=[]poker{{heitao,13},{heitao,13},{heitao,13}}

	zongpai1 := append(dipai, shoupai1...)
	zongpai2 := append(dipai, shoupai2...)
	zongpai3 := append(dipai, shoupai3...)
	fmt.Println("四条的检查结果如下")
	fmt.Println(defeat(zongpai1,zongpai2))
	fmt.Println(defeat(zongpai1,zongpai3))
}

func Test_hulu(t *testing.T){
	dipai := []poker{{heitao, 14}, {hongtao, 14}, {meihua, 14}, {fangkuai, 10}, {heitao, 13}}
	shoupai1 :=[]poker{{heitao, 13}, {heitao, 13}}
	shoupai2 :=[]poker{{heitao,10},{heitao,10}}
	shoupai3 :=[]poker{{heitao,10},{heitao,13}}

	zongpai1 := append(dipai, shoupai1...)
	zongpai2 := append(dipai, shoupai2...)
	zongpai3 := append(dipai, shoupai3...)
	fmt.Println("葫芦的检查结果如下")
	fmt.Println(defeat(zongpai1,zongpai2))
	fmt.Println(defeat(zongpai1,zongpai3))
}
func Test_tonghua(t *testing.T){
	dipai := []poker{{heitao, 5}, {heitao, 6}, {heitao, 7}, {heitao, 9}, {heitao, 10}}
	shoupai1 :=[]poker{{heitao, 3}, {heitao, 4}}
	shoupai2 :=[]poker{{heitao,13},{heitao,11}}
	shoupai3 :=[]poker{{heitao,8},{heitao,2}}

	zongpai1 := append(dipai, shoupai1...)
	zongpai2 := append(dipai, shoupai2...)
	zongpai3 := append(dipai, shoupai3...)
	fmt.Println("同花的检查结果如下")
	fmt.Println(defeat(zongpai1,zongpai2))
	fmt.Println(defeat(zongpai1,zongpai3))
}
func Test_shunzi(t *testing.T){
	dipai := []poker{{heitao, 14}, {hongtao, 13}, {meihua, 12}, {fangkuai, 2}, {heitao, 3}}
	shoupai1 :=[]poker{{heitao, 4}, {heitao, 5}}
	shoupai2 :=[]poker{{heitao,10},{heitao,11}}
	shoupai3 :=[]poker{{heitao,8},{heitao,7}}

	zongpai1 := append(dipai, shoupai1...)
	zongpai2 := append(dipai, shoupai2...)
	zongpai3 := append(dipai, shoupai3...)
	fmt.Println("\n\n顺子的检查结果如下")
	fmt.Println(defeat(zongpai1,zongpai2))
	fmt.Println(defeat(zongpai1,zongpai3))
}

func Test_santiao(t *testing.T){
	dipai := []poker{{heitao, 14}, {hongtao, 14}, {meihua, 10}, {fangkuai, 6}, {heitao, 8}}
	shoupai1 :=[]poker{{heitao, 14}, {heitao, 7}}
	shoupai2 :=[]poker{{heitao,14},{heitao,5}}
	shoupai3 :=[]poker{{heitao,9},{heitao,7}}

	zongpai1 := append(dipai, shoupai1...)
	zongpai2 := append(dipai, shoupai2...)
	zongpai3 := append(dipai, shoupai3...)
	fmt.Println("\n\n三条的检查结果如下")
	fmt.Println(defeat(zongpai1,zongpai2))
	fmt.Println(defeat(zongpai1,zongpai3))
}
func Test_duizi(t *testing.T){
	dipai := []poker{{heitao, 14}, {hongtao, 13}, {meihua, 10}, {fangkuai, 6}, {heitao, 8}}
	shoupai1 :=[]poker{{heitao, 14}, {heitao, 7}}
	shoupai2 :=[]poker{{heitao,14},{heitao,5}}
	shoupai3 :=[]poker{{heitao,6},{heitao,13}}

	zongpai1 := append(dipai, shoupai1...)
	zongpai2 := append(dipai, shoupai2...)
	zongpai3 := append(dipai, shoupai3...)
	fmt.Println("\n\n对子的检查结果如下")
	fmt.Println(defeat(zongpai1,zongpai2))
	fmt.Println(defeat(zongpai1,zongpai3))
}

func Test_sanpai(t *testing.T){
	dipai := []poker{{heitao, 14}, {hongtao, 13}, {meihua, 10}, {fangkuai, 6}, {heitao, 8}}
	shoupai1 :=[]poker{{heitao, 2}, {heitao, 3}}
	shoupai2 :=[]poker{{heitao,4},{heitao,5}}
	shoupai3 :=[]poker{{heitao,7},{heitao,11}}

	zongpai1 := append(dipai, shoupai1...)
	zongpai2 := append(dipai, shoupai2...)
	zongpai3 := append(dipai, shoupai3...)
	fmt.Println("\n\n散排的检查结果如下")
	fmt.Println(defeat(zongpai1,zongpai2))
	fmt.Println(defeat(zongpai1,zongpai3))
}