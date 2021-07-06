package main

import (
	//"container/list"
	"fmt"
)

const (
	_           = iota
	sanpai      = iota
	duizi       = iota
	liangdui    = iota
	santiao     = iota
	shunzi      = iota
	tonghua     = iota
	hulu        = iota
	sitiao      = iota
	tonghuashun = iota
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

//检查是否为四条，若是，则返回四条和踢脚的值
//申请15个大小是因为牌的点数从2-14，用tmp[14]的值表示14的个数。
func checksitiao(zp []poker) (int, []int, bool) {
	tmp := [15]int{}
	four, gangzi, tijiao := 0, 0, 0
	for _, pai := range zp {
		tmp[pai.daxiao]++
	}
	for i := len(tmp) - 1; i > -1; i-- {
		if tmp[i] == 4 {
			four++
			gangzi = i
		}
	}
	if four == 1 {
		for i := len(tmp) - 1; i > -1; i-- {
			if tmp[i] != 4 && tmp[i] > 0 {
				tijiao = i
			}
		}
		fmt.Println("检查结果是四条，四条是",gangzi,"踢脚是",tijiao)
		return sitiao, append([]int{}, gangzi, tijiao), true
	}
	return 0, []int{}, false
}

//检查是否是葫芦，返回三条的值和对子的值
func checkhulu(zp []poker) (int, []int, bool) {
	tmp := [15]int{}
	tmp3 := []int{}
	tmp2 := []int{}
	tmpdaxiao := []int{}
	three, two := 0, 0 //three two表示三体和对子的个数，
	for _, pai := range zp {
		tmp[pai.daxiao]++
		tmpdaxiao = append(tmpdaxiao, pai.daxiao)
	}
	for i := len(tmp) - 1; i > -1; i-- {
		if tmp[i] == 3 {
			three++
			tmp3 = append(tmp3, i)
		}
		if tmp[i] == 2 {
			two++
			tmp2 = append(tmp2, i)
		}
	}
	if three == 2 {
		if tmp3[0] > tmp3[1] {
			fmt.Println("检查结果是葫芦,三条是", tmp3[0], "对子是", tmp3[1])
			return hulu, tmp3[0:2], true
		} else {
			fmt.Println("检查结果是葫芦,三条是", tmp3[1], "对子是", tmp3[0])
			return hulu, append([]int{}, tmp3[1], tmp3[0]), true
		}

	} else if three == 1 && two == 2 {
		if tmp2[0] > tmp[1] {
			fmt.Println("检查结果是葫芦,三条是", tmp3[0], "对子是", tmp2[0])
			return hulu, []int{tmp3[0], tmp2[0]}, true
		} else {
			fmt.Println("检查结果是葫芦,三条是", tmp3[0], "对子是", tmp2[1])
			return hulu, []int{tmp3[0], tmp2[1]}, true
		}
	} else if three == 1 && two == 1 {
		fmt.Println("检查结果是葫芦,三条是", tmp3[0], "对子是", tmp2[0])
		return hulu, []int{tmp3[0], tmp2[0]}, true
	} else if three == 1 && two == 0 {
		top1 := Mymax(tmpdaxiao, tmp3[0], tmp3[0], tmp3[0])
		top2 := Mymax(tmpdaxiao, tmp3[0], tmp3[0], tmp3[0], top1)
		fmt.Println("检查结果是三条", tmp3[0], "踢脚是", top1, top2)
		return santiao, []int{tmp3[0], top1, top2}, true
	} else {
		return 0, []int{}, false
	}

}

//func fmtbynum(zp *[]poker) {
//	tmp := [13]int{}
//	four, three, two := 0, 0, 0
//	for _, pai := range *zp {
//		tmp[pai.daxiao]++
//	}
//	for i := len(tmp) - 1; i > -1; i-- {
//		if tmp[i] == 4 {
//			four++
//		}
//		if tmp[i] == 3 {
//			three++
//		}
//		if tmp[i] == 2 {
//			two++
//		}
//		if four == 1 {
//			fmt.Println("四条")
//		}
//		if three >= 1 && three+two >= 2 {
//			fmt.Println("葫芦")
//		}
//	}
//	//for index,i:=range tmp{
//
//	fmt.Println("123123")
//	fmt.Println(tmp)
//
//}

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
//检查顺子时，需要将ace作为1处理
func shunzilist(tmplist []int) (int, bool) {
	if len(tmplist) < 5 {
		return 0, false
	}
	for _, i := range tmplist {
		if i == 14 {
			tmplist = append(tmplist, 1)
		}
	}

	//冒泡排序，按从大到小的顺序排列，方便后面顺子的判断
	for _, _ = range tmplist {
		for i, _ := range tmplist {
			if i+1 < len(tmplist) && tmplist[i] < tmplist[i+1] {
				tmplist[i], tmplist[i+1] = tmplist[i+1], tmplist[i]
			}
		}

	}
	//从前往后遍历，遇到的第一个顺子即为最大的顺子，可以直接返回。
	for i := 0; i <= (len(tmplist) - 5); i++ {
		if tmplist[i] == (tmplist[i+1]+1) && tmplist[i+1]+1 == (tmplist[i+2]+2) && tmplist[i+2]+2 == (tmplist[i+3]+3) && tmplist[i+3]+3 == (tmplist[i+4]+4) {
			{
				return tmplist[i]*5 - 10, true
			}
		}
	}
	//遍历完，没有顺子则返回false
	return 0, false
}

//先检查是否是同花，如果是同花，则判断同花中是否包含顺子。返回同花顺的值。
func checktonghuashun(zongpai []poker) (int, []int, bool) {

	_, b, c := istonghua(zongpai)

	if c == true {
		sum, ok := shunzilist(b)
		if ok {
			fmt.Println("检查结果是同花顺，总值是", sum)
			return tonghuashun, []int{sum}, true
		} else {
			return 0, []int{}, false
		}
	}
	return 0, []int{}, false
}

//先判断是否同花，若是，再返回最大五张。
func checkmaxtonghua(zongpai []poker) (int, []int, bool) {
	_, tmplist, c := istonghua(zongpai)
	if c == true {
		for _, _ = range tmplist {
			for i, _ := range tmplist {
				if i+1 < len(tmplist) && tmplist[i] < tmplist[i+1] {
					tmplist[i], tmplist[i+1] = tmplist[i+1], tmplist[i]
				}
			}
		}
		fmt.Println("检查结果是同花", tmplist[0:5])
		return tonghua, tmplist[0:4], true
	} else {
		return 0, []int{}, c
	}
}

//只是判断是否同花，如果满足是同花，会把同花的牌丢给其他函数调用。
//判断传入的7张牌是否是同花，如是，则返回包含所有同花的一个slice和true，该函数会被tonghuashun和maxtonghua函数调用
func istonghua(zongpai []poker) (int, []int, bool) {
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
		return tonghua, tmp[heitao], true
	}
	if len(tmp[hongtao]) >= 5 {
		return tonghua, tmp[hongtao], true
	}
	if len(tmp[meihua]) >= 5 {
		return tonghua, tmp[meihua], true
	}
	if len(tmp[fangkuai]) >= 5 {
		return tonghua, tmp[fangkuai], true
	}
	return 0, []int{}, false
}

func checkshunzi(zp []poker) (int, []int, bool) {
	tmp := make(map[int]int)
	tmplist := make([]int, 0)
	//将手牌信息转化为字典格式，key是牌的值，value是这个牌值的数量
	for _, pai := range zp {
		_, ok := tmp[pai.daxiao]
		if ok {
			tmp[pai.daxiao]++
			//fmt.Println(tmp)
		} else {
			tmp[pai.daxiao] = 1
		}
		//tmp[i] = pai.daxiao
	}
	//将牌的所有unique值放入slice
	i := 0
	for k, _ := range tmp {
		tmplist = append(tmplist, k)
		i++
	}
	a, b := shunzilist(tmplist)
	if b==true{fmt.Println("检查结果是顺子，总值是",a)
		return shunzi, []int{a}, b}else{
		return 0, []int{a}, b
	}


}

//Mymax函数输入一个int类型的slice，和一个待删除的int类型的不定个数参数，返回剩余int中最大的那货
func Mymax(biglist []int, removelist ...int) int {
	for _, j := range removelist {
		for k := 0; k < len(biglist); k++ {
			if j == biglist[k] {
				biglist[k] = 0
				break
			}
		}
	}
	for i := 0; i < len(biglist)-1; i++ {
		for j := 0; j < len(biglist)-1; j++ {
			if biglist[j] < biglist[j+1] {
				biglist[j], biglist[j+1] = biglist[j+1], biglist[j]
			}
		}
	}
	return biglist[0]
}

//检查是否为两对
func checkduizi2(zp []poker) (int, []int, bool) {
	tmp := [15]int{}
	tmpdaxiao := []int{}
	two := 0 //three two表示三体和对子的个数，
	tmp2 := []int{}
	for _, pai := range zp {
		tmp[pai.daxiao]++
		tmpdaxiao = append(tmpdaxiao, pai.daxiao)
	}
	for i := len(tmp) - 1; i > -1; i-- {
		if tmp[i] == 2 {
			two++
			tmp2 = append(tmp2, i)
		}
		if two >= 2 {

			tijiao := Mymax(tmpdaxiao[:], tmp2[0], tmp2[0], tmp2[1], tmp2[1])
			fmt.Println("是有两对", tmp2[0], tmp2[1], "踢脚是", tijiao)
			return liangdui, []int{tmp2[0], tmp2[1], tijiao}, true
		}

	}
	return 0, []int{}, false
}

//检查是否为一对
func checkduizi1(zp []poker) (int, []int, bool) {
	tmp := [15]int{}
	tmpdaxiao := []int{}

	two := 0 //three two表示三体和对子的个数，
	tmp2 := []int{}
	for _, pai := range zp {
		//tmp数组报保存得信息是有多少个1，多少个2...多少个13
		//tmpdaxiaoslice保存得是具体得7个值得列表
		tmp[pai.daxiao]++
		tmpdaxiao = append(tmpdaxiao, pai.daxiao)
	}
	for i := len(tmp) - 1; i > -1; i-- {
		if tmp[i] == 2 {
			two++
			tmp2 = append(tmp2, i)
		}
	}
	if two >= 2 {
		return 0, []int{}, false
	} else if two == 1 {
		top1 := Mymax(tmpdaxiao[:], tmp2[0], tmp2[0])

		top2 := Mymax(tmpdaxiao[:], tmp2[0], top1)

		top3 := Mymax(tmpdaxiao[:], tmp2[0], top1, top2)
		fmt.Println("有一对", tmp2[0], "踢脚是", top1, top2, top3)
		return duizi, []int{tmp2[0], top1, top2, top3}, true
	} else {
		return 0, []int{}, false
	}

}

func checksanpai(zp []poker) (int, []int, bool) {
	tmpdaxiao := []int{}
	for _, pai := range zp {
		tmpdaxiao = append(tmpdaxiao, pai.daxiao)
	}
	for i := 0; i < len(tmpdaxiao)-1; i++ {
		for j := 0; j < len(tmpdaxiao)-1; j++ {
			if tmpdaxiao[j] < tmpdaxiao[j+1] {
				tmpdaxiao[j], tmpdaxiao[j+1] = tmpdaxiao[j+1], tmpdaxiao[j]
			}
		}
	}
	fmt.Println("检查结果是散牌", tmpdaxiao[0:5])
	return sanpai, tmpdaxiao[0:5], true
}

func checkzhonglei(zp1 []poker) (int, []int) {
	_, b, c := checktonghuashun(zp1)
	if c != true {
		_, b, d := checksitiao(zp1)
		if d != true {
			_, b, d := checkhulu(zp1)
			if d != true {
				_, b, c := checkmaxtonghua(zp1)
				if c != true {
					_, b, c := checkshunzi(zp1)
					if c != true {
						_, b, e := checkduizi2(zp1)
						if e != true {
							_, b, f := checkduizi1(zp1)
							if f != true {
								_, b, g := checksanpai(zp1)
								if g == true {
									return sanpai, b
								}
							}
							return duizi, b
						}
						return liangdui, b
					}
					return shunzi, b
				}
				return tonghua, b
			}
			return hulu, b
		}
		return sitiao, b
	}
	return tonghuashun, b
}

//当牌型相同，比如都为葫芦，都为顺子时，则比较返回的列表的大小。
//l1>l2 return 1,l1=l2 return 2 l1>l3 return 3
func bidaxiao(l1 []int, l2 []int) int {
	for i := 0; i < len(l1); i++ {
		if l1[i] > l2[i] {
			return 1
		} else if l1[i] < l2[i] {
			return 3
		}
	}
	return 2
}

//zp1>zp2返回1，zp1=zp2返回2，zp1>zp2返回3,牌检查失败返回4
func defeat(zp1 []poker, zp2 []poker) int {
	a := precheckpoker(zp1)
	if a == false {
		return 4
	}
	a = precheckpoker(zp2)
	if a == false {
		return 4
	}
	a1, b1 := checkzhonglei(zp1)
	a2, b2 := checkzhonglei(zp2)
	if a1 > a2 {
		fmt.Println("第一个大")
		return 1

	} else if a1 < a2 {
		fmt.Println("第二个大")
		return 3
	} else if a1==a2 {
		for i := 0; i < len(b1); i++ {
			if b1[i] > b2[i] {
				fmt.Println("第一个大")
				return 1
			} else if b1[i] < b2[i] {
				fmt.Println("第二个大")
				return 3
			}
		}
		fmt.Println("一样大")
		return 2

	}
	return 100
}

//检查poker牌的合法性，花色取值1，2，3，4。牌点取值2-14.ace按14点输入。
//后续加入总牌数限制=7，牌张不可重复等
func precheckpoker(zp []poker) bool {
	for _, b := range zp {
		if b.huase < 1 && b.huase > 4 {
			fmt.Println("花色只能是黑红梅方")
			return false
		}
		if b.daxiao < 2 && b.daxiao > 14 {
			fmt.Println("牌点数在2-14中间，ace当14")
			return false
		}
	}
	return true
}
func main() {
	tonghuadipai := []poker{{heitao, 3}, {heitao, 6}, {heitao, 8}, {heitao, 5}, {heitao, 10}}
	dipai := []poker{{heitao, 2}, {heitao, 6}, {hongtao, 8}, {heitao, 5}, {heitao, 10}}
	shoupai1 := []poker{{hongtao, 11}, {hongtao, 9}}
	shoupai2 := []poker{{hongtao, 3}, {hongtao, 13}}
	//shoupai3 := []poker{{heitao,3},{hongtao,13}}
	zongpai1 := append(dipai, shoupai1...)
	zongpai2 := append(dipai, shoupai2...)
	tonghuazongpai1 := append(tonghuadipai, shoupai1...)
	tonghuazongpai2 := append(tonghuadipai, shoupai2...)
	//checktonghuashun(zongpai1)
	//maxtonghua(zongpai1)
	//a, b := tonghua(zongpai1)
	//fmt.Println(a, b)
	//checkshunzi(zongpai1)
	//fmtbynum(&zongpai1)
	//a, b, c := checkhulu(&zongpai1)
	//fmt.Println(a, b, c)
	//四条检测要在三条之前，因为按3+2等于三条，会导致3=4的被误判。
	//3+3取最大的三条
	//fmt.Println(checkduizi2(&zongpai1))
	//fmt.Println(checkduizi1(&zongpai1))
	fmt.Println(defeat(tonghuazongpai1, tonghuazongpai2))
	fmt.Println("11223344")
	fmt.Println(checkzhonglei(zongpai1))
	fmt.Println("11223344")
	fmt.Println(checkzhonglei(zongpai2))

}
