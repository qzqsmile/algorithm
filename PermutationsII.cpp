/*
����Ҫע��2��
1.�ȶ������������Լ��ٸ��Ӷ�
2.Ҫ�ж��Ƿ���ȴӶ������ظ�
*/
class Solution {
public:
    void build(vector<vector<int> >&res, vector<int> nums, int k)
    {
        if(k == nums.size())
        {
            res.push_back(nums);
            return ;
        }
        for(int i = k; i < nums.size(); i++)
        {
            if((i != k) && (nums[i] == nums[k]))
                continue;
            swap(nums[k], nums[i]);
            build(res, nums, k+1);
            //swap(nums[i], nums[k]);
            //res.push_back(nums);
        }
    }
    vector<vector<int>> permuteUnique(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        vector<vector<int> >res;
        if(!nums.size()) return res;
        build(res, nums, 0);
        return res;
    }
};