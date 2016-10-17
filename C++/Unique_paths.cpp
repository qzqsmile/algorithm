class Solution {
	public:
		int uniquePaths(int m, int n) {
			int num = 0;
			int **df = new int *[m];
			for(int i = 0; i < m; i++)
				df[i] = new int[n];

			df[0][0] = 1;
			for(int i = 0; i < m; i++)
			{
				for(int j = 0; j < n; j++)
				{
					if((i >= 1) && (j >= 1))
					{
						df[i][j] = df[i-1][j] + df[i][j-1];    
					}else if(i == 0)
					{
						df[i][j] = 1;
					}else if(j == 0){
						df[i][j] = 1;
					}
				}
			}

			return df[m-1][n-1];
		}
};
