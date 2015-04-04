class Solution {
	public:
		string countAndSay(int n) {
			string res = "1";
			string s="1";

			for( ; n > 1; n--)
			{
				res.clear();
				int count = 0; 
				char s1 = s[0];

				for(int cur = 0; cur < s.size(); cur++)
				{
					if(s1 == s[cur])
						count++;
					else
					{
						res += to_string(count);
						res += s1;
						s1 = s[cur];
						count = 1;
					}
				}
				res += to_string(count);
				res += s1;
				s = res;
			}

			return res;
		}
};
