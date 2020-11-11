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
