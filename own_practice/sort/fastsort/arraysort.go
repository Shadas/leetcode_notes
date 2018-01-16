package fastsort

//
func ArrayFastSort1(values []int) []int {
	arrayFastSort1(values)
	return values
}

func arrayFastSort1(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	arrayFastSort1(values[:head])
	arrayFastSort1(values[head+1:])
}

//
func ArrayFastSort2(values []int) []int {
	arrayFastSort2(values, 0, len(values)-1)
	return values
}

func arrayFastSort2(values []int, left, right int) {
	tmp, p := values[left], left
	i, j := left, right
	for i < j {
		if j >= p && values[j] >= tmp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= tmp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = tmp
	if p-left > 1 {
		arrayFastSort2(values, left, p-1)
	}
	if right-p > 1 {
		arrayFastSort2(values, p+1, right)
	}
}
