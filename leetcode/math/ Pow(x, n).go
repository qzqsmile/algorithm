func myPow(x float64, n int) float64 {
    if n == 1{
        return x
    }
    if n == 0{
        return 1
    }
    pn := abs(n)
    t := myPow(x, pn/2)
    var pr float64
    if pn % 2 == 1{
        pr = t*t*x
    }else{
        pr = t*t
    }
    if n < 0{
        return 1 / pr
    }
    return pr
}

func abs(x int) int{
    if x < 0{
        return -x
    }
    return x
}