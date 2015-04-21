class Solution 
{
	void getRes(const string &s, vector<int> *dp, int r, string ns, vector<string> &res)
	{
		for (int l : dp[r])
			if (l)
				getRes(s, dp, l, " " + s.substr(l, r - l) + ns, res);
			else
				res.push_back(s.substr(0, r) + ns);
	}
	public:
	vector<string> wordBreak(string s, unordered_set<string> &dict) 
	{
		int l(s.length());
		vector<int> dp[l + 1];
		vector<string> res;
		dp[0].push_back(-1);
		for (int i = 1; i <= l; ++ i)
			for (int j = 0; j < i; ++ j)
				if (dict.count(s.substr(j, i - j)) && !dp[j].empty())
					dp[i].push_back(j);

		getRes(s, dp, l, "", res);
		return res;
	}
};
