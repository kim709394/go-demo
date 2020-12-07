package hello

/*
@Author kim
@Description
@date 2020-11-11 11:23
*/

//插入排序算法.sort为true则从小到大，为false为从大到小
func InsertSort(slice []int, sort bool) {

	if slice == nil || len(slice) == 0 {
		return
	}
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if sort {
				if slice[j] < slice[j-1] {
					slice[j-1], slice[j] = slice[j], slice[j-1]
				} else {
					break
				}
			} else {
				if slice[j] > slice[j-1] {
					slice[j-1], slice[j] = slice[j], slice[j-1]
				} else {
					break
				}
			}
		}
	}
}

//希尔排序算法
func ShellSort(slice []int, sort bool) {
	len := len(slice)
	if slice == nil || len == 0 {
		return
	}
	for incr := len / 2; incr > 0; incr /= 2 {
		insertSortByIncr(slice, sort, incr)
	}

}

//根据下标增量对切片进行分组插入排序
func insertSortByIncr(slice []int, sort bool, incr int) {
	for i := 0; i < len(slice); i += incr {
		for j := i; j > 0; j -= incr {
			if sort {
				if slice[j] < slice[j-incr] {
					slice[j-incr], slice[j] = slice[j], slice[j-incr]
				} else {
					break
				}
			} else {
				if slice[j] > slice[j-incr] {
					slice[j-incr], slice[j] = slice[j], slice[j-incr]
				} else {
					break
				}
			}
		}
	}

}

//快速排序,递归方式，sort为true则从小到大，反之从大到小
func QuickSort(slice []int, sort bool) {
	len := len(slice)
	if len == 0 || slice == nil {
		return
	}

	quickSort(slice, sort, 0, len-1)
}

func quickSort(slice []int, sort bool, low int, high int) {
	if low == high {
		return
	}
	var index int
	if sort {
		index = getIndexAsc(slice, low, high)
	} else {
		index = getIndexDesc(slice, low, high)
	}
	if index >= low && index <= high {
		if index > low {
			quickSort(slice, sort, low, index-1)
		}
		if index < high {
			quickSort(slice, sort, index+1, high)
		}
	}

}

//返回基准数排序后的下标
func getIndexAsc(slice []int, low int, high int) int {
	//以第一个数作为基准数
	tmp := slice[low]
	//判断从前往后比较还是从后往前比较
	flag := true
	for {
		if low == high {
			slice[low] = tmp
			return low
		}
		if flag {
			if slice[high] < tmp {
				slice[low] = slice[high]
				low++
				flag = false
			} else {
				high--
			}
		} else {
			if slice[low] < tmp {
				low++
			} else {
				slice[high] = slice[low]
				high--
				flag = true
			}
		}

	}
}
func getIndexDesc(slice []int, low int, high int) int {
	//以第一个数作为基准数
	tmp := slice[low]
	//判断从前往后比较还是从后往前比较
	flag := true
	for {
		if low == high {
			slice[low] = tmp
			return low
		}
		if flag {
			if slice[high] > tmp {
				slice[low] = slice[high]
				low++
				flag = false
			} else {
				high--
			}
		} else {
			if slice[low] > tmp {
				low++
			} else {
				slice[high] = slice[low]
				high--
				flag = true
			}
		}
	}
}
