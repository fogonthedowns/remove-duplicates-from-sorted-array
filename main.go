package main

func removeDuplicates(nums []int) int {
	// initialize last
	last := 0

	// loop over len of nums
	for i := 0; i < len(nums); i++ {
		// first value or if this value and ext are not equal
		// write i
		if i == 0 || nums[i] != nums[i-1] {
			nums[last] = nums[i]
			// increment last
			last++
		}
	}
	// fmt.Println(last)
	// fmt.Println(".")
	// fmt.Println(nums)
	return last
}
