#include<vector>
#include<iostream>

using namespace std;

class Solution {
	public:
		int minPathSum(vector<vector<int> > &grid) {
			if(grid.size() == 0)
				return 0;
			vector<vector<int> >dp;
			vector<int> row;

			//initilize row 1
			for(int i = 0; i != grid[0].size(); i++)
			{
				if(i == 0)
					row.push_back(grid[0][0]);
				else
					row.push_back(row[i-1]+grid[0][i]);
			}
		dp.push_back(row);
			if(grid.size() < 1)
				return row[row.size()-1];
			row.clear();

			for(int i = 1; i != grid.size(); i++)
			{
				for(int j = 0; j != grid[0].size(); j++)
				{
					if(j == 0)
						row.push_back(dp[i-1][j]+grid[i][j]);
					else
						row.push_back(min(row[j-1],dp[i-1][j])+grid[i][j]);
				}
				dp.push_back(row);
				row.clear();
			}

			return dp[grid.size()-1][grid[0].size()-1];
		}
};

int main(void)
{
	vector<vector<int> >grid;
	vector<int> row1;
	vector<int> row2;
	vector<int> row3;
	Solution s;
	
	row1.push_back(1);
	row1.push_back(3);
	row1.push_back(1);
	row2.push_back(1);
	row2.push_back(5);
	row2.push_back(1);
	row3.push_back(4);
	row3.push_back(2);
	row3.push_back(1);
	grid.push_back(row1);
	grid.push_back(row2);
	grid.push_back(row3);
	
	cout <<	s.minPathSum(grid) << endl;

	return 0;
}
