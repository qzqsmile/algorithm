#include<iostream>
#include<string>
#include<unordered_map>
#include<cmath>

using namespace std;

class Solution {
	public:
		string fractionToDecimal(int numerator, int denominator) {
			string res;

			if(numerator == 0) return "0";
			if(numerator < 0 ^ denominator< 0) res += '-';

			long long n = abs((long long)numerator);
			long long d = abs((long long)denominator);

			res += to_string(n / d);

			if(n % d == 0)
				return res;
			res += '.';

			unordered_map<int , int> map;
			for(long long   r = n % d; r; r %= d)
			{
				if(map.count(r) > 0)
				{
					res.insert(map[r],1,'(');
					res += ')';
					break;
				}
				map[r] = res.size();

				r *= 10;            
				res += to_string(r / d);
			}

			return res;
		}
};

int main(void)
{
	Solution s;
	string s1;

	s1 = s.fractionToDecimal(3, 7);

	cout << s1 <<endl;

	return 0;
}
