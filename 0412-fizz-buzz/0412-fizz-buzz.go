import("strconv")

func fizzBuzz(n int) []string {
    var res []string
    for i := 1; i <= n; i++{
        if i % 3 == 0 && i % 5 == 0 {
            res = append(res, "FizzBuzz")
        } else if i % 3 == 0 {
            res = append(res, "Fizz")
        } else if i % 5 == 0 {
            res = append(res, "Buzz")
        } else {
            s := strconv.Itoa(i)
            res = append(res, s)
        }
    }
    return res
    
}