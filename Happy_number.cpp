class Solution {
public:
    bool isHappy(int n) {
        unordered_set<int> set;
        
        while(n != 1)
        {
            if(set.find(n) != set.end())
                return false;
            set.insert(n);
            int c = n;
            n = 0;
            while(c)
            {
                n += pow((c%10),2);
                c /= 10;
            }
        }
        
        return true;
    }
};