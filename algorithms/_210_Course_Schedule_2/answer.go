package _210_Course_Schedule_2

func findOrder(numCourses int, prerequisites [][]int) []int {
	return findOrderBfs(numCourses, prerequisites)
}

func findOrderBfs(numCourses int, prerequisites [][]int) []int {
	var (
		matrix   [][]int
		indegree []int
		queue    []int
		ret      []int
	)
	// 初始化 邻接矩阵、入度列表
	for i := 0; i < numCourses; i++ {
		tmp := []int{}
		for j := 0; j < numCourses; j++ {
			tmp = append(tmp, 0)
		}
		matrix = append(matrix, tmp)
		indegree = append(indegree, 0)
	}

	// 构造邻接矩阵、入度列表
	for _, p := range prerequisites {
		ready, pre := p[0], p[1]
		indegree[ready]++ // 入度+1
		matrix[pre][ready] = 1
	}
	// 筛选所有入度为空的节点
	for i, degree := range indegree {
		if degree == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		// 出队列
		course := queue[0]
		queue = queue[1:]
		// 放入结果集
		ret = append(ret, course)
		// 在邻接矩阵找所有该节点的后驱节点
		for i := 0; i < numCourses; i++ {
			if matrix[course][i] != 0 {
				indegree[i]--         // 其入度-1
				if indegree[i] == 0 { // 如果其入度减到0，则加入队列
					queue = append(queue, i)
				}
			}
		}
	}
	if len(ret) != numCourses {
		return []int{}
	}
	return ret
}
