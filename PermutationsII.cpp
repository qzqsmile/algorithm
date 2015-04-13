class Solution {
	public:
		int lengthOfLongestSubstring(string s) {

			int m[129] = {0};
			int i, j;
			int cnt = 0, pre = 0;
			int max = 0;
			int c;

			for (i = 0; c = s[i]; i++) {
				if (pre < m[c]) {
					if (max < cnt)
						max = cnt;

					cnt = i-m[c];
					pre = m[c];
				}

				cnt++;
				m[c] = i+1;
			}
			return max > cnt ? max : cnt;
		}
};
class Solution {
	public:
		vector<vector<int> > permute(vector<int> &num) {
			vector<vector<int> > result;

			permuteRecursive(num, 0, result);
			return result;
		}

		// permute num[begin..end]
		// invariant: num[0..begin-1] have been fixed/permuted
		void permuteRecursive(vector<int> &num, int begin, vector<vector<int> > &result)    {
			if (begin >= num.size()) {
				// one permutation instance
				result.push_back(num);
				return;
			}

			for (int i = begin; i < num.size(); i++) {
				swap(num[begin], num[i]);
				permuteRecursive(num, begin + 1, result);
				// reset
				swap(num[begin], num[i]);
			}
		}
};
