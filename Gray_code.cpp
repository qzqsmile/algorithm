class Solution {
	public:
		vector<int> grayCode(int n) {
			vector<int> v(1,0);
			for (int i=1;i<(1<<n);i++) v.push_back(v[i-1]^(i&-i));
			return v;
		}
};
