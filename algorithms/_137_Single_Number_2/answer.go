package _137_Single_Number_2

const Times = 3

func singleNumber(nums []int) int {
	mark := []bool{}
	for i := 0; i < len(nums); i++ {
		mark = append(mark, false)
	}

	for i, n := range nums {
		if mark[i] == false {
			if i == len(nums)-1 {
				return n
			} else {
				tmp := nums[i+1:]
				times := 0
				for j, m := range tmp {
					if m == n {
						mark[i+j+1] = true
						times++
						if times == Times-1 {
							break
						}
					}
				}
				if times != Times-1 {
					return n
				}
			}
		} else {
			continue
		}
	}
	return 0
}
