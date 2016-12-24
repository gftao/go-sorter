package qsort

func QuickSort(values []int) {
	l := len(values)
	if l <= 1 {
		return
	}
	first := 0
	end := l - 1
	key := values[end /2]
	for {
		for values[first] < key {
			first++
		}
		for values[end] > key {
			end--
		}

		if first >= end {
			break
		}
		values[end], values[first] = values[first], values[end]
		first++
		end-- //每次较换后向后取值，否则（切片中有相同值时）会出现死循环
	}

	QuickSort(values[:first])
	QuickSort(values[end +1:])
}

func QuickSort1(values []int) {
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
	QuickSort(values[:head])
	QuickSort(values[head+1:])
}
