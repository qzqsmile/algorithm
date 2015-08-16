/*
这题有几点要注意的
1 注意回溯结束的条件有两个
2 注意Ip每个域不能有多个0即不能有000.000，每个0一开始就要压入，压出
*/
class Solution {
public:
    vector<string> restoreIpAddresses(string s) {
        vector<string> res;
        string  row;
        build(res, row, s, 0, 1);
        return res;
    }
    
    void build(vector<string>&res, string &row, const string &s, int begin, int scope)
    {
        if(begin >= s.size()) 
            return;
   
        if(scope == 4)
        {
            if(s[begin] == '0' && (begin < (s.size()-1))) return;
            int num = 0; 
            for(int i = begin; i < s.size(); i++)
            {
                num = s[i] - '0' + 10 * num;
                if(num > 255) return;
            }
            //row + s.substr(begin, (s.size()-begin));
            res.push_back( row + s.substr(begin, (s.size()-begin)));
            //row.erase(begin+scope,(s.size()-begin));
            return ;
        }
        
        if(s[begin] == '0')
        {
            row.push_back('0');
            row.push_back('.');
            build(res, row, s, begin+1, scope+1);
            row.erase(begin+scope-1, 2);
            return ;
        }
        
        int num = 0;
        for(int i = begin; i < (begin+3); i++)
        {
            num = 10 * num + s[i] - '0';
            if(num <= 255) 
            {
                row += s.substr(begin, i-begin+1);
                row.push_back('.');
                build(res, row, s, i+1, scope+1);
                row.erase(begin+scope-1, i-begin+2);
            }
        }
    }
};