//iterator from left to right

func toHex(num int) string {
    if num == 0{
        return "0"
    }
    ans := ""
    for i := 7; i >= 0; i-- {
        val := num >> (4 * i) & 0xf
        var digit byte
        if val > 0 || len(ans) > 0{
            if val < 10 {
                digit = '0' + byte(val)
            } else {
                digit = 'a' + byte(val-10)
            }
            ans = ans + string(digit) 
        }
    }

    return ans
}
