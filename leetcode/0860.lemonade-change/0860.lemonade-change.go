package problem0860

// 860. 柠檬水找零
// 在柠檬水摊上，每一杯柠檬水的售价为 5 美元。顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。

// 每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。

// 注意，一开始你手头没有任何零钱。

// 给你一个整数数组 bills ，其中 bills[i] 是第 i 位顾客付的账。如果你能给每位顾客正确找零，返回 true ，否则返回 false 。

func lemonadeChange(bills []int) bool {
	mem := map[int]int{5: 0, 10: 0}
	var isok func(bill int) bool
	isok = func(bill int) bool {
		if bill == 5 {
			mem[5]++
			return true
		}
		if bill == 10 {
			if mem[5] < 1 {
				return false
			}
			mem[5]--
			mem[10]++
			return true
		}
		if bill == 20 {
			if mem[10] >= 1 && mem[5] >= 1 {
				mem[10]--
				mem[5]--
				return true
			}
			if mem[5] >= 3 {
				mem[5] -= 3
				return true
			}
			return false
		}
		return false
	}
	for _, b := range bills {
		if !isok(b) {
			return false
		}
	}
	return true
}
