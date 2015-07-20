class Solution {
public:
    int climbStairs(int n) {
       int low1, low2;
       int res;
       
       if(n == 0) return 0;
       if(n == 1) return 1;
       low1 = 1;
       low2 = 1;
       
       for(int i = 1; i < n; i++)
       {
           res = low1 + low2;
           low1 = low2; 
           low2 = res;
       }
       
       return res;
    }
};