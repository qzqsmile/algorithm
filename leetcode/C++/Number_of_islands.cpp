class Solution {
public:
    int numIslands(vector<vector<char>> &grid) {
        int count = 0;
        for(int i = 0; i < grid.size(); i++)
        {
            for(int j = 0; j < grid[0].size(); j++)
            {
                if(grid[i][j] == '1')
                {
                    count++;
                    setzero(grid, j, i);
                }
            }
        }
        return count;
    }
    
    void setzero(vector<vector<char>> &grid, int col, int row)
    {
        grid[row][col] = '0';
        if((row >= 1) && (grid[row-1][col]== '1'))
        {
            setzero(grid, col, row-1);
        }
        if((row < (grid.size()-1)) && (grid[row+1][col] == '1'))
        {
            setzero(grid, col, row+1);
        }
        if((col >= 1) && (grid[row][col-1] == '1'))
        {
            setzero(grid, col-1, row);
        }
        if((col < (grid[0].size()-1))&&(grid[row][col+1] == '1'))
        {
            setzero(grid, col+1, row);
        }
    }
};