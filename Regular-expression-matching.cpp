class Solution {
public:
    bool is_match(string &s, string &p , size_t indexs, size_t indexp)
    {
        if(indexp >= p.size()) return (indexs >= s.size());
        
        if('*' == p[indexp+1])
        {
            return (is_match(s, p, indexs, indexp+2)) || ((indexs < s.size()) && (s[indexs] == p[indexp] || '.' == p[indexp]) && is_match(s, p, indexs+1, indexp));
        }
        else
        {
            if(indexs >= s.size()) return false;
            return (s[indexs] == p[indexp] || s[indexs] == '.') ? is_match(s, p, indexs+1, indexp+1) : false;
        }
    }
     bool isMatch(string s, string p) {
        return is_match(s, p, 0, 0);
    }
};