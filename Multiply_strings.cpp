#include<string>
#include<iostream>

using namespace std;

class Solution {
	public:
		string multiply(string num1, string num2) {
			string result(num1.size()+num2.size(),'0');

			for(int i = 0; i < num1.size(); i++)
			{
				int  carry = 0;
				for(int j = 0; j < num2.size(); j++)
				{
					int temp = (result[i+j]-'0') + (num1[i] - '0') * (num2[j]-'0') + carry;
					result[i+j] = temp %  10 + '0';
					carry = temp / 10;
				}
			}

			size_t startpos = result.find_first_not_of("0");
			if (string::npos != startpos) {
				return result.substr(startpos);
			}

			return "0";
		}
};

int main(void)
{
	Solution s;
	
	cout << s.multiply("1", "1")<< endl;

	return 0;
}
