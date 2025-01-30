package functions 

func SumB(nums ...int) int {
    var out int
    for _, n := range nums {
        out += n
    }
    return out
}
