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

/*
class Solution {
	public:
		int lengthOfLongestSubstring(string s) {
			int maxlen = 0, left = 0;
			map<char, int> m;

			for(int i = 0 ; i < s.size(); i++)
			{
				if(m.count(s[i]))
				{
					if(m[s[i]] < left)
					{
						m[s[i]] = i;
						continue;
					}
					if((i-left) > maxlen)
						maxlen = i - left;
					left = m[s[i]] + 1;
					m[s[i]] = i;
				}
				else
					m[s[i]] = i;
			}
			if(maxlen < (s.size() - left))
				maxlen = s.size() - left;
			return maxlen;
		}
};*/
