class Solution {
	public:
		int atoi(string str) {
			long long num = 0;
			long long result = 0;
			int s  = 0, flag = 1;

			if(!str.size())
				return 0;
			while(str[s] == ' ')
				s++;
			if(str[s] == '-')
				flag = -1;
			if((str[s] == '-') || (str[s] == '+')) 
			{
				s++;
				if(str[s] == '+' || str[s] == '-')
					return 0;
			}

			while(isdigit(str[s]))
			{
				num = 10 * num + str[s] - '0';
				s++;
				result = num * flag;

				if(result > INT_MAX)
					return INT_MAX;
				if(result < INT_MIN)
					return INT_MIN;
			}

			return result;
		}
};
