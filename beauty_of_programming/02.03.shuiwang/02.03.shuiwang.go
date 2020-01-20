package problem02_03

//寻找发帖水王
func Find(ids []int) int {
	count := 0
	var condition int
	for _, id := range ids {
		if count == 0 {
			condition = id
			count = 1
		} else {
			if condition == id {
				count += 1
			} else {
				count -= 1
			}
		}
	}
	return condition
}
