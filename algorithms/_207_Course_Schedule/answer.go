package _207_Course_Schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		matrix   = [][]int{}
		indegree = []int{}
	)
	// 初始化 邻接矩阵二维数组，及入度数组
	for i := 0; i < numCourses; i++ {
		tmp := []int{}
		for j := 0; j < numCourses; j++ {
			tmp = append(tmp, 0)
		}
		matrix = append(matrix, tmp)
		indegree = append(indegree, 0)
	}
	// 生成邻接矩阵及入度列表
	for _, p := range prerequisites {
		ready, pre := p[0], p[1]
		if matrix[pre][ready] == 0 {
			indegree[ready]++
		}
		matrix[pre][ready] = 1
	}

	var (
		count = 0
		queue = []int{}
	)
	// 每一个入度为0的节点，进入一个队列
	for i, d := range indegree {
		if d == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		course := queue[0]
		queue = queue[1:]
		count++
		for i := 0; i < numCourses; i++ {
			if matrix[course][i] != 0 { // 节点i为course的后驱节点
				indegree[i]-- // 其入度减一
				if indegree[i] == 0 {
					queue = append(queue, i)
				}
			}
		}
	}

	return count == numCourses
}
