func numIdenticalPairs(nums []int) int {
    numCounts := make(map[int]int)
    var pair int
    for _, num := range nums {
        numCounts[num]++
    }
    for _, count := range numCounts {
        pair += (count * (count - 1)) / 2
    }
    return pair
}
