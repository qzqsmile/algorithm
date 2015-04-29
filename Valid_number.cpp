class Solution {
	public:
		bool isNumber(string s) {
			//remove space
			int begin = 0, end = s.size();
			int locate_e = -1;

			while((s[begin] == ' ' || s[begin] == '-' || (s[begin] == '+')) && (begin < end)) 
			{
				if((begin > 0) && (s[begin] == ' ') && (s[begin-1] != ' ')) return false;
				begin++;
			}
			while((s[end-1] == ' ') && (begin < end)) end--;

			//locate e index
			for(int i = begin; i < end; i++)
			{
				if(s[i] == 'e')
				{
					locate_e = i;
					break;
				}
			}

			if(locate_e == -1)
			{
				return numwithnoe(s, begin, end);
			}
			else
			{

				bool s1 = numwithnoe(s, begin, locate_e);

				locate_e++;
				if(s[locate_e] == '+' || s[locate_e] == '-')
					locate_e++;

				bool s2 = isnum(s, locate_e, end);

				return s2 && s1;
			}
		}

		bool numwithnoe(const string &s, int begin, int end)
		{
			if (begin >= end) 
				return false;
			//locate '.'
			int locate_decimal = -1;

			for (int i = begin; i < end; i++)
			{
				if(s[i] == '.')
				{
					locate_decimal = i;
					break;
				}
			}

			if(locate_decimal == -1)
				return isnum(s, begin, end);
			else
			{
				for(int i = begin; i < locate_decimal; i++)
				{
					if(!isdigit(s[i]))
						return false;
				}

				for(int i = locate_decimal+1; i < end; i++)
				{
					if(!isdigit(s[i]))
						return false;
				}

				if(end - begin == 1)
					return false;
				return true;
			}
		}

		bool isnum(const string &s, int begin, int end)
		{
			if(begin >= end)
				return false;
			for(int i = begin; i < end; i++)
			{
				if(!isdigit(s[i]))
					return false;
			}

			return true;
		}
};
