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
