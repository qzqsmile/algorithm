class Solution {
	public:
		bool isPalindrome(string s) {
			if(!s.size())
				return true;
			int pre = 0, end = s.size()-1;

			while(pre <= end)
			{
				if((isalpha(s[pre])||(isdigit(s[pre]))) && ((isalpha(s[end])) || (isdigit(s[end]))))
				{
					if(!isEqual(s[pre], s[end]))
						return false;
					pre++;
					end--;
				}
				else
				{
					if(isalpha(s[pre])||isdigit(s[pre]))
						end--;
					else
						pre++;
				}
			}
		}

		bool isEqual(char a, char b)
		{
			if((a == b) || (a - b == 'A' - 'a') || (b - a == 'A' - 'a'))
				return true;
			return false;
		}
};
