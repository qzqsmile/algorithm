class Solution {
	public:
		int sqrt(int x) {
			if((0 == x) || (1 == x))
				return x;
			int l = 1, r = x, res;
			while(l <= r)
			{
				int m = (l + r) / 2;
				if(m == x / m)
					return m;
				else if(m > x / m){
					r = m - 1;
				}else{
					l = m + 1;
					res = m;
				}
			}

			return res;
		}
};
