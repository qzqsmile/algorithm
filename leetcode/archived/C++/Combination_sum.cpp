class Solution {
	public:
		vector<vector<int> > combinationSum(vector<int> &candidates, int target) {
			vector<vector<int> >result;
			vector<int> output;
			sort(candidates.begin(), candidates.end());
			search(result, output, target, candidates.begin(), candidates.end());

			return result;
		}

		void search(vector<vector<int> >&result, vector<int>output, int target, vector<int>::iterator begin, vector<int>::iterator end)
		{
			if(0 == target)
			{
				result.push_back(output);
				return;
			}

			if(*begin <= target)
			{
				output.push_back(*begin);
				search(result, output, target-*begin, begin, end);
				output.pop_back();
			}

			if(begin+1 < end)
				search(result, output, target, begin+1, end);
		}
};
