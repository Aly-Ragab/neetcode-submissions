import (
	"slices"
)

func hasDuplicate(nums []int) bool {
    unique := make([]int,0)
    for _,num := range nums {
        if !slices.Contains(unique, num) {
            unique = append(unique, num)
        }
    }

    return len(unique) != len(nums)
}
