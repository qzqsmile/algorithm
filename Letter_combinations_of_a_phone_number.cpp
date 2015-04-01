class Solution {
	public:
		vector<string> letterCombinations(string digits) {
			vector<string> res;
			if(!digits.size())
				return res;
			map<char, string> letter;
			letter['2'] = "abc";
			letter['3'] = "def";
			letter['4'] = "ghi";
			letter['5'] = "jkl";
			letter['6'] = "mno";
			letter['7'] = "pqrs";
			letter['8'] = "tuv";
			letter['9'] = "wxyz";

			string s;
			makecombinations(res, s, letter, digits, 0);
			return res;
		}

		void makecombinations(vector<string> &res, string& s, map<char, string>& letter, string& digits, int next_num)
		{
			if(s.size() == digits.size())
			{
				res.push_back(s);
			}
			else
			{
				if(next_num < digits.size())
				{
					for(int i = 0; i < letter[digits[next_num]].size(); i++)
					{
						s.push_back(letter[digits[next_num]][i]);
						makecombinations(res, s, letter, digits, next_num+1);
						s.pop_back();
					}
				}
			}
		}
};
