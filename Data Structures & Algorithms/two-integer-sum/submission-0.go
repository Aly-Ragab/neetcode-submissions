func twoSum(nums []int, target int) []int {
    visited := make(map[int]int)

    for i, value := range nums {

        numberToAdd := target - value
        
        if index, existed := visited[numberToAdd]; existed {
            return []int{index, i}
        } 

        visited[value] = i

    }

    return []int{} 
}
