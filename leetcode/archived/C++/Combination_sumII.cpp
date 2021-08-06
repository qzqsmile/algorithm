class Solution {
	public:
		vector<vector<int> > combinationSum2(vector<int> &num, int target) {
			vector<vector<int> >result;
			vector<int> output;
			sort(num.begin(), num.end());
			search(result, output, target, num.begin(), num.end());
			sort(result.begin(),result.end());

			result.erase(unique(result.begin(), result.end()), result.end());
			return result;
		}   

		void search(vector<vector<int> >&result, vector<int>output, int target, vector<int>::iterator begin, vector<int>::iterator end)
		{   
			if(0 == target)
			{   
				result.push_back(output);
				return;
			}   

			if((*begin <= target) && (begin < end)) 
			{   
				output.push_back(*begin);
				search(result, output, target-*begin, begin+1, end);
				output.pop_back();
			}   

			if(begin+1 < end)
				search(result, output, target, begin+1, end);
		}   
};
