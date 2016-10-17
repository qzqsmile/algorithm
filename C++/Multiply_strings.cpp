#include<string>
#include<iostream>

using namespace std;

class Solution {
	public:
		string multiply(string num1, string num2) {
			string result(num1.size()+num2.size(),'0');

			reverse(num1.begin(), num1.end());
			reverse(num2.begin(), num2.end());

			for(int i = 0; i < num1.size(); i++)
			{
				int  carry = 0;
				for(int j = 0; (j < num2.size() || carry > 0); j++)
				{
					if(j < num2.size())
					{
						int temp = result[i+j]-'0' + (num1[i] - '0') * (num2[j]-'0') + carry;
						result[i+j] = temp %  10 + '0';
						carry = temp / 10;
					}
					else
					{
						int temp = result[i+j] - '0' + carry;
						result[i+j] = temp %  10 + '0';
						carry = temp / 10;
					}

				}
			}

			reverse(result.begin(), result.end());

			size_t startpos = result.find_first_not_of("0");
			if (string::npos != startpos) {
				return result.substr(startpos);
			}

			return "0";
		}
};
