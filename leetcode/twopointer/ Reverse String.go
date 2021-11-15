func reverseString(s []byte)  {
    b, e := 0, len(s)-1
    for;b < e;{
        s[b], s[e] = s[e], s[b]
        b++
        e--
    }
}