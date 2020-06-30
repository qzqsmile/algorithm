#include<vector>
#include<iostream>

using namespace std;

class Solution {
	public:
		int uniquePathsWithObstacles(vector<vector<int> > &obstacleGrid) {
			if(0 == obstacleGrid.size())
				return 0;
			int row, col;
			row = obstacleGrid.size();
			col = obstacleGrid[0].size();

			int **df = new int *[row];
			for(int i = 0; i != row; i++)
				df[i] = new int[col];

			for(int i = 0; i != row; i++)
			{
				for(int j = 0; j != col; j++)
				{
					df[i][j] = 0;
				}
			}

			for(int i = 0; i != col; i++)
			{
				if(obstacleGrid[0][i] != 1)
					df[0][i] = 1;
				else
					break;
			}

			if(row == 1)
				return df[0][col-1];

			for(int i = 1; i != row; i++)
			{
				for(int j = 0; j != col; j++)
				{
					if(obstacleGrid[i][j] == 1)
						continue;
					else
					{
						if(j == 0)
							df[i][j] = df[i-1][j];
						else
							df[i][j] = df[i-1][j] + df[i][j-1];
					}
				}
			}

			return df[row-1][col-1];
		}
};

int main(void)
{
	Solution s;
	vector<int> row;
	row.push_back(1);
	vector<vector<int> > obstacleGrid;

	obstacleGrid.push_back(row);

	cout << s.uniquePathsWithObstacles(obstacleGrid) << endl;

	return 0;
}
