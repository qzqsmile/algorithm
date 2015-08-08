class Solution {
	public:
		vector<string> anagrams(vector<string> &strs) {
			vector<string> res;
			map<string ,int> anagrams;
			
			//这里关键是在int中存入第一个出现的下标，而不是次数。
			//以后出现可以利用find来判断，不需要记录次数
			for(int i = 0; i != strs.size(); i++)
			{
				string s = strs[i];
				sort(s.begin(), s.end());
				auto l = anagrams.find(s);
				if(l == anagrams.end()){
					anagrams[s] = i;
				}else{
					res.push_back(strs[i]);
					if(l->second != -1)
					{
						res.push_back(strs[l->second]);
						l->second = -1;
					}
				}
			}

			return res;
		}
};
