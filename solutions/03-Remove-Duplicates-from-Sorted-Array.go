package solutions

func removeDuplicates(nums []int) int {
    cnt := 1
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] != nums[i+1] {
            nums[cnt] = nums[i+1]
            cnt++
        }
    }
    return cnt
}