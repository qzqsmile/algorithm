/*
 *利用一个counter记录t中的字母数
 *如果s中出现一个就将counter减1
 *如果为0，则统计end-begin，
 *这里要注意，begin刚开始并且指向合理的位置，
 *不停的循环使其指向刚开始的符合的位置
 * */
class Solution {
public:
    string minWindow(string s, string t) {
        vector<int> map(128, 0);
        for(auto c:t) map[c]++;
        int d = INT_MAX, begin = 0, end = 0, head = 0, counter = t.size();
        
        while(end < s.size())
        {
            if(map[s[end++]]-- > 0) counter--;
            while(counter == 0){
                if(end-begin < d) d = end - (head = begin);
                if(map[s[begin++]]++ == 0) counter++;
            }
        }
        
        return d==INT_MAX ? "":s.substr(head, d);
    }
};
