package slidingwindows

func lengthOfLongestSubstring(s string) int {
    left, right := 0, 0 
    maxLength := 0
    windows := make(map[uint8]int)

    for;right < len(s);{
        windows[s[right]]++
        right++
        for;windows[s[right-1]] == 2;{
            windows[s[left]]--
            left++
        }
        if right-left > maxLength{
            maxLength = right-left
        }
    }
    return maxLength
}