func convertToTitle(columnNumber int) string {
    res := ""
    for;columnNumber > 0;{
        if columnNumber % 26 != 0{
            res = string(columnNumber % 26 +'A'-1) + res
            columnNumber = columnNumber / 26
        }else{
            res = "Z"+res 
            columnNumber = columnNumber - 26
            columnNumber = columnNumber / 26
        }
    }
    return res
}