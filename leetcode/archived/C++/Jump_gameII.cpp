class Solution {
	public:
		int jump(vector<int>& nums) {
			int cur_max = nums[0], max_step = nums[0];
			int n = nums.size();
			int cur_step = 1;

			for(int i = 1; i <= min(max_step, n); i++)
			{
				cur_max = max(cur_max, i+nums[i]);
				if(i == n-1)
				{
					return cur_step;
				}
				if(i == max_step)
				{
					max_step = cur_max;
					cur_step++;
				}
			}

			return 0;
		}
};
