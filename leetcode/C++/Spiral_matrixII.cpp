class Solution {
	public:
		vector<vector<int> > generateMatrix(int n) {
			vector<vector<int> > res( n, vector<int>(n));
			int k = 1, i = 0, j;

			while(k < n * n)
			{
				j = i;
				//top row left to right
				while(j < (n-i-1))
					res[i][j++] = k++;
				//right col top to bottom
				j = i;
				while(j < (n-i-1))
					res[j++][n-1-i] = k++;
				//bottom row right to left
				j = n-i-1;
				while(j > i)
					res[n-i-1][j--] = k++;
				//left col down to top
				j = n-i-1;
				while(j > i)
					res[j--][i] = k++;
				i++;
			}

			if(k == (n * n))
				res[i][i] = k;

			return res;
		}
};
