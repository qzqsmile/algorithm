class Solution {
	public:
		vector<string> letterCombinations(string digits) {
			map<char, string> m;
			m['2'] = "abc";
			m['3'] = "def";
			m['4'] = "ghi";
			m['5'] = "jkl";
			m['6'] = "mno";
			m['7'] = "pqrs";
			m['8'] = "tuv";
			m['9'] = "wxyz";
			vector<string> res;
			if(!digits.size()) return res;
			string s;
			letter(res, s, digits, m, 0);
			return res;
		}

		void letter(vector<string> &res, string &s,const string& s1, map<char, string> &m, int l)
		{
			if(l == s1.size())
				res.push_back(s);
			for(int i = 0; i < m[s1[l]].size(); i++)
			{
				s.push_back(m[s1[l]][i]);
				letter(res, s, s1,m, l+1);
				s.pop_back();
			}
		}
};
