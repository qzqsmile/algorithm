#include<string>
#include<limits.h>
#include<cmath>
#include<iostream>

using namespace std;

class Solution {
	public:
		int divide(int dividend, int divisor) {

			if(divisor == 1)
				return dividend;
			if((dividend == INT_MIN) && (divisor == -1))
				return INT_MAX;
			long long d = abs(dividend);
			long long s = abs(divisor);
			long long ans = 0;

			/*while(n>=d){
			  int a=d;
			  int m=1;
			  while((a<<1) < n){a<<=1;m<<=1;}
			  ans+=m;
			  n-=a;
			  }*/
			while(d >= s)
			{
				long long a = s;
				long long m = 1;
				while((a << 1) < d){ a <<= 1; m <<= 1;}
				ans += m;
				d -= a;
			}

			 if(((dividend >= 0)&&(divisor <=0)) || ((dividend <= 0) && (divisor >= 0)))
				 return -ans;

			return ans;
		}
};

int main(void)
{
	Solution s;
	cout << s.divide(-1010369383,-2147483648) << endl;

	return 0;
}
