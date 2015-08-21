class Solution {
	public:
		vector<vector<int>> permute(vector<int>& nums) {
			vector<vector<int> >res;
			if(!nums.size()) return res;
			vector<int>r;
			r.push_back(nums[0]);
			res.push_back(r);

			build(res, nums, 1);

			return res;
		}

		void build(vector<vector<int> > &res, const vector<int>&nums, int index)
		{
			vector<vector<int> >res_next;
			if(index >= nums.size())
			{
				return ;
			}
			for(int i = 0; i < res.size(); i++)
			{
				//这里注意插入操作不会导致复杂性过高
				for(auto j = 0; j <= res[i].size(); j++)
				{
					vector<int> r = res[i];
					r.insert(r.begin()+j, nums[index]);
					res_next.push_back(r);
				}
			}
			res = res_next;
			build(res, nums, index+1);
		}
};
