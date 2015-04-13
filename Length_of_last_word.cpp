class Solution {
	public:
		int lengthOfLastWord(const char *s) {
			const char *begin = s, *last_word = NULL, *last_space= NULL;

			while(*s != '\0')
				s++;
			while(s >= begin)
			{
				if(isalpha(*s))
				{
					last_word = s;
					break;
				}
				s--;
			}
			while(s >= begin)
			{
				if(isspace(*s))
				{
					last_space = s;
					break;
				}
				s--;
			}

			if(last_word == NULL)
				return 0;
			if(last_space == NULL)
				return last_word - begin + 1;
			return last_word - last_space;
		}
};
