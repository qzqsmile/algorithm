class Solution {
public:
    int titleToNumber(string s) {
        int col_num = 0;
        
        col_num = cal_col_num(s, 0, s.size()-1);
        
        return col_num;
    }
    
    int cal_col_num(string s, int pre, int end){
        if(pre == end)
            return s[pre]-'A'+1;
        else
            return (s[end]-'A'+1)+ 26*cal_col_num(s, pre, end-1);
    }
};