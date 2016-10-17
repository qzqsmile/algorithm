class Solution {
	public:
		double pow(double x, int n) {
			if(0 == n)
				return 1;
			if (n < 0){
				x = 1 / x;
				n = -n;
			}
			if(n == 2)
				return x * x;
			double pow1 = pow(x, n /2);
			if(n % 2)
				return x * pow1 * pow1;
			else
				return pow1 * pow1;
		}
};
