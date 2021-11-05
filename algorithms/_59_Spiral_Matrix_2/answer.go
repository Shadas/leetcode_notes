package _59_Spiral_Matrix_2

func generateMatrix(n int) [][]int {
	return generateMatrixWithGenSeq(n)
}

func generateMatrixWithGenSeq(n int) [][]int {
	var (
		ret, record [][]int
	)
	// 初始化
	ret = make([][]int, n)
	record = make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
		record[i] = make([]int, n)
	}
	// 遍历赋值
	count := 0 // 用于记录已经遍历的次数
	// 起始位置
	pos := Pos{0, 0}
	dire := Right
	for count < n*n { // 当遍历次数打到元素个数时结束
		ret[pos.Y][pos.X] = count + 1
		pos, record, dire = newPos(pos, record, dire)
		count++
	}
	return ret
}

// 方向枚举
type Dire int

const (
	Right Dire = iota
	Down
	Left
	Up
)

type Pos struct {
	X, Y int
}

func newPos(pos Pos, record [][]int, dire Dire) (newPos Pos, newRecord [][]int, newDire Dire) {
	// 更新标记
	record[pos.Y][pos.X] = 1
	newRecord = record
	// 更新方向
	newDire = dire
	// 尝试移动
	switch dire {
	case Right:
		pos.X += 1
		if pos.X >= len(record[0]) || record[pos.Y][pos.X] == 1 {
			pos.X -= 1 // 恢复
			pos.Y += 1
			newDire = Down
		}
	case Down:
		pos.Y += 1
		if pos.Y >= len(record) || record[pos.Y][pos.X] == 1 {
			pos.X -= 1
			pos.Y -= 1
			newDire = Left
		}

	case Left:
		pos.X -= 1
		if pos.X < 0 || record[pos.Y][pos.X] == 1 {
			pos.X += 1
			pos.Y -= 1
			newDire = Up
		}

	case Up:
		pos.Y -= 1
		if pos.X < 0 || record[pos.Y][pos.X] == 1 {
			pos.X += 1
			pos.Y += 1
			newDire = Right
		}
	}
	newPos = pos
	return
}
