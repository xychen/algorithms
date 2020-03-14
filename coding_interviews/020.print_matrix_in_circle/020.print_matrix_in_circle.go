package problem020

//import "fmt"


func PrintMatrixClockwisely(numbers []int, columns int, rows int) []int {
	if columns <= 0 || rows <= 0 {
		return nil
	}
	if len(numbers) != columns*rows {
		return nil
	}
	//结果集
	res := make([]int, 0)
	for start := 0; (start*2 <= (columns - 1)) && (start*2 <= (rows - 1)); start += 1 {
		//fmt.Println("--+++-", start)
		tmpRes := PrintMatrixInCircle(numbers, columns, rows, start)
		res = append(res, tmpRes...)
	}
	return res
}

func PrintMatrixInCircle(numbers []int, columns, rows, start int) []int {
	res := make([]int, 0)
	endX := columns - 1 - start
	endY := rows - 1 - start
	//打印行
	for i := start; i <= endX; i++ {
		n := columns*start + i
		res = append(res, numbers[n])
	}
	//打印向下的列
	if start < endY {
		for i := start + 1; i < rows-start; i++ {
			n := columns*i + endX
			res = append(res, numbers[n])
		}
	}

	//从右到左打印行
	if start < endY && start < endX {
		for i := endX - 1; i >= start; i-- {
			n := columns*endY + i
			res = append(res, numbers[n])
		}
	}

	//从下往上打印列
	if start < endX && start < endY - 1 {
		for i := endY - 1; i >= start + 1; i-- {
			n := columns*i + start
			res = append(res, numbers[n])
		}
	}
	return res
}
