package main

import (
	//"container/list"
	"fmt"
)

const (
	heitao   = 1
	hongtao  = 2
	meihua   = 3
	fangkuai = 4
)

//import ""

type poker struct {
	huase  int //
	daxiao int
}

//type dipai struct {
//	a poker
//	b poker
//	c poker
//	d poker
//	e poker
//}
//type shoupai struct {
//	f poker
//	g poker
//}
//type zongpai struct {
//	dipai
//	shoupai
//}

//func (zp zongpai) tonghuashun() bool {
//	return true
//}

//遍历7张牌，看是否有同花。后续加入都是同花之后的大小判断。
//func tonghua(shoupai []poker) bool {
//	heitao, hongtao, meihua, fangkuai := 0, 0, 0, 0
//	for _, pai := range shoupai {
//		switch pai.huase {
//		case 1:
//			heitao++
//		case 2:
//			hongtao++
//		case 3:
//			meihua++
//		case 4:
//			fangkuai++
//
//		}
//
//	}
//	if heitao >= 5 || hongtao >= 5 || meihua >= 5 || fangkuai >= 5 {
//		fmt.Println("true")
//		return true
//	}
//	fmt.Println("flase")
//	return false
//}
//检查输入的slice是否是顺子，返回最大顺子五张牌的相加值和true
func shunzilist(tmplist []int) (int, bool) {
	if len(tmplist) < 5 {
		return 0, false
	}
	//冒泡排序，按从大到小的顺序排列，方便后面顺子的判断
	for _, _ = range tmplist {
		for i, _ := range tmplist {
			if i+1 < len(tmplist) && tmplist[i] < tmplist[i+1] {
				tmplist[i], tmplist[i+1] = tmplist[i+1], tmplist[i]
			}
		}
	}
	fmt.Println(tmplist)
	//从前往后遍历，遇到的第一个顺子即为最大的顺子，可以直接返回。
	for i := 0; i <= (len(tmplist) - 5); i++ {
		if tmplist[i] == (tmplist[i+1]+1) && tmplist[i+1]+1 == (tmplist[i+2]+2) && tmplist[i+2]+2 == (tmplist[i+3]+3) && tmplist[i+3]+3 == (tmplist[i+4]+4) {
			fmt.Println("shunzi")
			fmt.Println(tmplist[i]*5 - 10)
			{
				return tmplist[i]*5 - 10, true
			}
		}
	}
	//遍历完，没有顺子则返回false
	return 0, false
}
//先检查是否是同花，如果是同花，则判断同花中是否包含顺子。返回同花顺的值。
func tonghuashun(zongpai []poker) (int, bool) {
	a, b := tonghua(zongpai)
	if b == true {
		sum, ok := shunzilist(a)
		if ok {
			fmt.Println(sum, "tonghuashun")
			return sum, true
		} else {
			return 0, false
		}
	}
	return 0, false
}
//先判断是否同花，若是，再返回最大五张。
func maxtonghua(zongpai []poker)([]int,bool){
	tmplist,b:=tonghua(zongpai)
	if b==true{
		for _, _ = range tmplist {
			for i, _ := range tmplist {
				if i+1 < len(tmplist) && tmplist[i] < tmplist[i+1] {
					tmplist[i], tmplist[i+1] = tmplist[i+1], tmplist[i]
				}
			}
		}
		fmt.Println("maxtonghua",tmplist[0:5])
		return tmplist[0:4],true
	}else {
		return tmplist,b
	}
}



//判断传入的7张牌是否是同花，如是，则返回包含所有同花的一个slice和true，该函数会被tonghuashun和maxtonghua函数调用
func tonghua(zongpai []poker) ([]int, bool) {
	tmp := make(map[int][]int)
	for _, pai := range zongpai {
		switch pai.huase {
		case heitao:
			tmp[heitao] = append(tmp[heitao], pai.daxiao)
		case hongtao:
			tmp[hongtao] = append(tmp[hongtao], pai.daxiao)
		case meihua:
			tmp[meihua] = append(tmp[meihua], pai.daxiao)
		case fangkuai:
			tmp[fangkuai] = append(tmp[fangkuai], pai.daxiao)
		}
	}
	if len(tmp[heitao]) >= 5 {
		return tmp[heitao], true
	}
	if len(tmp[hongtao]) >= 5 {
		return tmp[hongtao], true
	}
	if len(tmp[meihua]) >= 5 {
		return tmp[meihua], true
	}
	if len(tmp[fangkuai]) >= 5 {
		return tmp[fangkuai], true
	}
	fmt.Println(tmp)
	return tmp[heitao], false
}



func shunzi(zp []poker) (int, bool) {
	tmp := make(map[int]int)
	tmplist := make([]int, 0)
	//将手牌信息转化为字典格式，key是牌的值，value是这个牌值的数量
	for _, pai := range zp {
		_, ok := tmp[pai.daxiao]
		if ok {
			tmp[pai.daxiao]++
			fmt.Println(tmp)
		} else {
			tmp[pai.daxiao] = 1
		}
		//tmp[i] = pai.daxiao
	}
	//将牌值单独取出放入一个slice切片中，如果列表数量小于5说明则不可能有顺子，直接返回false
	i := 0
	for k, _ := range tmp {
		tmplist = append(tmplist, k)
		i++
	}
	a, b := shunzilist(tmplist)
	return a, b
}

func main() {
	dipai1 := []poker{{heitao, 6}, {heitao, 7}, {hongtao, 9}, {heitao, 9}, {heitao, 12}}
	shoupai1 := []poker{{heitao, 10}, {heitao, 8}}
	zongpai1 := append(dipai1, shoupai1...)
	tonghuashun(zongpai1)
	maxtonghua(zongpai1)
	//a, b := tonghua(zongpai1)
	//fmt.Println(a, b)
	shunzi(zongpai1)
}
