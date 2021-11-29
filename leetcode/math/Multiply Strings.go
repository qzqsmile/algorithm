func multiply(num1 string, num2 string) string {
    res := ""
    shift := 0
    for i := len(num2)-1; i >= 0; i--{
        tmp := multi(num1, string(num2[i]))
        tmp = shiftTmp(tmp, shift)
        fmt.Println("before:", res, tmp)
        res = add(tmp, res)
        shift++
    }
    if len(res) > 0 && res[0] == '0'{
        return "0"
    }
    return res
}
func shiftTmp(str string, count int) string{
    res := str 
    for ;count > 0; count--{
        res += "0"
    }
    return res 
}

func multi(nums1 string, nums2 string) string{
    res := ""
    sum, plus := 0, 0
    
    for i := len(nums1)-1; i >= 0; i--{
        sum = int(nums2[0]-'0') * int(nums1[i]-'0')+plus 
        res = strconv.Itoa(sum % 10) + res
        plus = sum / 10
    }
    if plus > 0{
        res = strconv.Itoa(plus) + res
    }
    return res
}

func add(nums1 string, nums2 string) string{
    res := ""
    i, j := len(nums1)-1, len(nums2)-1
    sum := 0
    plus := 0
    for ;i >= 0 && j >= 0;{
        sum = int(nums1[i]-'0')+int(nums2[j]-'0')+plus
        res = strconv.Itoa(sum % 10) + res
        plus = sum / 10
        i--
        j--
    }
    for;i >= 0; i--{
        sum = int(nums1[i] - '0')+plus
        res = strconv.Itoa(sum % 10) + res
        plus = sum / 10
    }

    for;j >= 0; j--{
        sum = int(nums2[j] - '0')+plus
        res = strconv.Itoa(sum % 10) + res
        plus = sum / 10
    }
    if plus > 0{
        res = strconv.Itoa(plus) + res
    }
    return res
}