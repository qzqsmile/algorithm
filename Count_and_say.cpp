class Solution {
public:
    string countAndSay(int n) {
        string res("1");
        
        for(int i = 0; i < (n-1); i++)
        {
            int k = 0;
            string s = res;
            res.clear();
            for(int j = 0; j < s.size(); j=k)
            {
                int count = 0;
                for( k = j; s[k] == s[j] && k < s.size(); k++)
                    count++;
                res.push_back(count+'0');
                res.push_back(s[j]);
            }
        }
        
        return res;
    }
};  
