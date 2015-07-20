class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        vector<int> res;
        int flag = 1;
        
        for(int i = digits.size()-1; i >= 0; i--)
        {
            int s = digits[i] + flag;
            res.push_back(s%10);
            flag = s / 10;
        }
        
        if(flag) 
            res.push_back(flag);
        reverse(res.begin(), res.end());
        
        return res;
    }
};