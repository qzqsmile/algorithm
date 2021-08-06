class Solution {
	public:
		string convertToTitle(int n) {
			string s;

			while(n > 0)
			{
				s += (n % 26) + 'A' - 1;
				n = n / 26;
			}

			for(int i = 0; i != s.size(); i++)
			{
				if(s[i] >= 'A')
					continue;
				else
				{
					if(i != (s.size() -1))
					{
						s[i] = 'Z' + s[i] - 'A' + 1;
						s[i+1]--;
					}
				}
			}

			if(s[s.size()-1] < 'A')
				s.pop_back();
			reverse(s.begin(), s.end());

			return s;
		}
};
