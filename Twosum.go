package main

import "fmt"

func twoSum(nums []int, target int) []int {
    n := len(nums)

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }

    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Println(result) 

	nums = []int{3,2,4}
    target = 6
    result = twoSum(nums, target)
    fmt.Println(result) 

	nums = []int{3,3}
    target = 6
    result = twoSum(nums, target)
    fmt.Println(result) 
}
