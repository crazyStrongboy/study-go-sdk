package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	gradient := make(map[int]int)
	// 每个课程初始化梯度都是0
	for i := 0; i < numCourses; i++ {
		gradient[i] = 0
	}

	priorityCourse := make(map[int][]int)
	// 遍历 prerequisites 给每个课程生成对应的梯度
	for i := 0; i < len(prerequisites); i++ {
		p := prerequisites[i]
		cur := p[1] // 依赖课程优先
		next := p[0]
		gradient[next]++ // 增加梯度
		priorityCourse[cur] = append(priorityCourse[cur], next)
	}

	var queue []int
	// 取梯度是0的加到队列，需要优先学习
	for k, v := range gradient {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	// 遍历队列直至所有课程梯度都缩减到最低
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if priorityCourse[top] == nil {
			continue
		}
		for _, next := range priorityCourse[top] {
			gradient[next]--
			if gradient[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// 所有课程梯度都是0，则表示可以全部学完
	for _, v := range gradient {
		if v != 0 {
			return false
		}
	}

	return true
}
