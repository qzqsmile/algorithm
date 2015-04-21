class Solution {
	public:
		string expand(string&s, int l, int h)
		{
			while((l >= 0) && (h < s.size()) && (s[l] == s[h]))
			{
				l--; h++;
			}

			return s.substr(l+1, h-1-l);
		}

		string longestPalindrome(string s) {
			string res = s.substr(0,1);

			for(int i = 0; i < s.size(); i++)
			{
				string p;
				p = expand(s, i, i);

				if(p.size() > res.size())
					res = p;
				p = expand(s, i, i+1);

				if(p.size() > res.size())
					res = p;
			}

			return res;
		}
};
