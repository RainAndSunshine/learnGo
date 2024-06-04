package _2share_gold_test

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	//保存用户名——切片
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	//保存用户的金币数量——map
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	//遍历用户
	for _, name := range users {
		//遍历用户名的每一位字母
		for _, v := range name {
			//使用switch实现
			switch v {
			case 'e', 'E':
				if _, ok := distribution[name]; !ok {
					distribution[name] = 1
				} else {
					distribution[name]++
				}
				coins--
				break
			case 'i', 'I':
				if _, ok := distribution[name]; !ok {
					distribution[name] = 2
				} else {
					distribution[name] += 2
				}
				coins -= 2
				break
			case 'o', 'O':
				if _, ok := distribution[name]; !ok {
					distribution[name] = 3
				} else {
					distribution[name] += 3
				}
				coins -= 3
				break
			case 'u', 'U':
				if _, ok := distribution[name]; !ok {
					distribution[name] = 4
				} else {
					distribution[name] += 4
				}
				coins -= 4
				break
			}
			//使用if实现
			//if v == 'e' || v == 'E' {
			//	if _, ok := distribution[name]; !ok {
			//		distribution[name] = 1
			//	} else {
			//		distribution[name]++
			//	}
			//	coins--
			//}
			//if v == 'i' || v == 'I' {
			//	if _, ok := distribution[name]; !ok {
			//		distribution[name] = 2
			//	} else {
			//		distribution[name] += 2
			//	}
			//	coins -= 2
			//}
			//if v == 'o' || v == 'O' {
			//	if _, ok := distribution[name]; !ok {
			//		distribution[name] = 3
			//	} else {
			//		distribution[name] += 3
			//	}
			//	coins -= 3
			//}
			//if v == 'u' || v == 'U' {
			//	if _, ok := distribution[name]; !ok {
			//		distribution[name] = 4
			//	} else {
			//		distribution[name] += 4
			//	}
			//	coins -= 4
			//}
		}
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	//打印每一位用户的金币数目
	for _, name := range users {
		fmt.Println(name, "：", distribution[name])
	}
}
