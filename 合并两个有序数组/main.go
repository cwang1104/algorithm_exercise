package main

/*
	给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
	请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
*/

func main() {
	merge([]int{1, 2, 3, 0, 0}, 3, []int{1, 2}, 2)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	var sortedNum []int = make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		//已遍历完nums1
		if p1 == m {
			sortedNum = append(sortedNum, nums2[p2:]...)
			break
		}
		//已遍历完nums2
		if p2 == n {
			sortedNum = append(sortedNum, nums1[p1:]...)
			break
		}

		//哪个小一点则往前移动一位
		if nums1[p1] < nums2[p2] {
			sortedNum = append(sortedNum, nums1[p1])
			p1++
		} else {
			sortedNum = append(sortedNum, nums2[p2])
			p2++
		}
	}
	copy(nums1, sortedNum)
}
