class Solution {
	public:
		void nextPermutation(vector<int>& nums) {
			if(!nums.size()) return ;
			cal_next(nums, nums.size()-1);

		}

		void cal_next(vector<int> & nums, int end)
		{
			if(end < 0)
			{
				sort(nums.begin(), nums.end());
				return ;
			}

			for(int i = nums.size()-1; i >= end; i--)
			{
				if(nums[i] > nums[end])
				{
					swap(nums[i], nums[end]);
					sort(nums.begin()+end+1, nums.end());
					return ;
				}
			}
			cal_next(nums, end-1);
		}
};
