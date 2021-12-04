func rand10() int {
    n := rand7()
    if n == 7{
        return rand10()
    }else if n % 2 == 0{
        nc := rand7()
        if nc > 5{
            return rand10()
        }
        return nc
    }else{
        nc := rand7()
        if nc > 5{
            return rand10()
        }
        return 5+nc
    }
}