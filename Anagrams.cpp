class Solution {
	public:
		vector<string> anagrams(vector<string> &strs) {
			vector<string> res;
			map<string ,int> anagrams;

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
