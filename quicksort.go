package main 

import "fmt"
func partition(nums []int, low int, high int) int {
    pivot := nums[low]
    for low < high {
        for low < high && nums[high] >= pivot {
            high --
            fmt.Println("1")
        }
        nums[low] = nums[high]
        for low < high && nums[low] <= pivot {
            low ++
            fmt.Println("2")
        }
        nums[high] = nums[low]
    }
    nums[low] = pivot
    return low
}

func quickSort(nums []int, low int, high int) {
    if low < high {
        pivot_pos := partition(nums, low, high)
        quickSort(nums, 0, pivot_pos - 1)
        quickSort(nums, pivot_pos + 1, high)
    } 
}

func sortArray(nums []int) []int {
    quickSort(nums, 0, len(nums) - 1)
    return nums
}


func main() {
    nums := []int{-74,48,-20,2,10,-84,-5,-9,11,-24,-91,2,-71,64,63,80,28,-30,-58,-11,-44,-87,-22,54,-74,-10,-55,-28,-46,29,10,50,-72,34,26,25,8,51,13,30,35,-8,50,65,-6,16,-2,21,-78,35,-13,14,23,-3,26,-90,86,25,-56,91,-13,92,-25,37,57,-20,-69,98,95,45,47,29,86,-28,73,-44,-46,65,-84,-96,-24,-12,72,-68,93,57,92,52,-45,-2,85,-63,56,55,12,-85,77,-39}
    sortArray(nums)
    fmt.Println(nums)
} 
