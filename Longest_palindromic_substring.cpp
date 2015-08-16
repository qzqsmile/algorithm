/*
这个算法要注意以下几点
1.以某个字符为中心向两边扩展从而确定是否是回文
2.确定时要注意奇偶数的问题，因为偶数最中心两点必相等，所以可以通过判断是否相等去除这种情形
3.某一串连续的数字仅仅在刚最开始的字符可能存在最大的回文 如 *aaaaab*，仅仅需要判断第一个a，原因在于其中i仅仅是帮助
 寻找begin 和 end点进行比较，出去第一a之外在点b一定会停下来。这也是下边代码中i = end的原因
*/
class Solution {
public:
    string longestPalindrome(string s) {
        if (s.empty()) return "";
        if (s.size() == 1) return s;
        int max_len = 0, max_begin = 0;
        
        for(int i = 0; i < s.size();)
        {
            if (s.size() - i <= max_len / 2) break;
            int begin,end;
            end = i;
            while((end < (s.size()-1)) && (s[end] == s[end+1])) end++;
            end++;
            begin = i -1;
            i = end;
            while((end < s.size()) && (begin >= 0) && (s[begin] == s[end]))
            {
                begin--;
                end++;
            }
            int new_len = end-begin-1;
            if(new_len > max_len)
            {
                max_len = new_len;
                max_begin = begin+1;
            }
        }
        
        return s.substr(max_begin, max_len);
    }
};