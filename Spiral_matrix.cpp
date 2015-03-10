class Solution {
	public:
		vector<int> spiralOrder(vector<vector<int> > &matrix) {
			vector<int> res;

			if(matrix.size() > 0)
				printmatrix(res, matrix, 0, matrix.size()-1, matrix[0].size()-1);

			return res;
		}

		void printmatrix(vector<int> &res, vector<vector<int> >&matrix, int begin, int m, int n)
		{
			if((begin >= m) || (begin >= n))
			{
				if((begin == m) && (m == n))
					res.push_back(matrix[begin][begin]);
				else if((begin == m) && (n > begin))
				{
					for(int i = begin; i <= n; i++)
						res.push_back(matrix[begin][i]);
				}
				else if((begin == n) && (m > begin))
				{
					for(int i = begin; i <= m; i++)
						res.push_back(matrix[i][begin]);
				}
				return ;
			}
			else
			{
				//row
				for(int i = begin; i < n; i++)
				{
					res.push_back(matrix[begin][i]);
				}

				//col
				for(int i = begin; i < m; i++)
				{
					res.push_back(matrix[i][n]);
				}

				//bottom row

				for(int i = n; i > begin; i--)
				{
					res.push_back(matrix[m][i]);
				}

				//left col
				for(int i = m; i > begin; i--)
				{
					res.push_back(matrix[i][begin]);
				}

				printmatrix(res, matrix, begin+1, m-1, n-1);
			}
		}
};
