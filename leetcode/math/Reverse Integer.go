
func reverse(x int) int {
    if x > (pow(2, 31)-1) || x < (pow(-2, 31)){
        return 0
    }
    res := 0
    negative := false
    if x < 0{
        negative = true
    }

    if negative{
        x = -x
    }

    for;x > 0;{
        res = 10 * res + x % 10
        x /= 10
    }

    if negative{
        res = -res
    }
    if res > (pow(2, 31)-1) || res < (pow(-2, 31)){
        return 0
    }
    return res
}

func pow(x int, count int) int{
    res := 1
    for;count > 0; count--{
        res = res * x
    }
    return res
}
