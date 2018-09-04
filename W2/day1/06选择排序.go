package main

import "fmt"

func main() {
	var nums []int = []int{10, 5, 12, 3, 14, 2, 8, 6, 11, 9, 0, 7, 4, 1, 13}
	fmt.Println("排序前：", nums)
	sortArray(nums)
	fmt.Println("排序后：", nums)
}

func sortArray(nums []int) {
	//逐个查看0-8号宝座
	for i := 0; i < len(nums)-1; i++ {
		//锁定i号宝座
		for j := i + 1; j < len(nums); j++ {

			//发现更有资格做0号宝座的，双方互换位置
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

}
