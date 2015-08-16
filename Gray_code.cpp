/*
  因为要求的数是2的幂次方，可以从低位到高位对每个位^，这样的数
  每个位不同，注意要从result.size()开始取起，反向取
*/
class Solution {
public:
    vector<int> grayCode(int n) {
        vector<int> result = { 0 };
        int t = 1;
        for(int i = 0; i<n; i++) {
            for(int j = result.size() - 1; j >= 0; j--)
                result.push_back(result[j]^t);
            t <<= 1;
        }
    
        return result;
    }
};