class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        if(!matrix.size())  return vector<int>();
       int row  = matrix.size(), col = matrix[0].size();
       vector<int>res;
       
       game(res, matrix, 0, row-1, col-1);
       return res;
    }
    
//这里关键是要把长度和宽度为一的边特殊考虑
    void game(vector<int>& res, vector<vector<int>> &matrix, int begin, int row, int col)
    {
        if((begin > row) || (begin > col))
            return;
       
        if(begin == col)
        {
            for(int i = begin; i <= row; i++)
            {
                res.push_back(matrix[i][col]);
            }
            return ;
        }
        if(begin == row)
        {
            for(int i = begin; i <= col; i++)
            {
                res.push_back(matrix[row][i]);
            }
            return;
        }
        //up left to right
        for(int i = begin; i < col; i++)
        {
            res.push_back(matrix[begin][i]);
        }
        
        //right up to down
        for(int i = begin; i < row; i++)
        {
            res.push_back(matrix[i][col]);
        }
        
        //down right to left
        for(int i = col; i > begin; i--)
        {
            res.push_back(matrix[row][i]);
        }
        
        //left down to up
        for(int i = row; i > begin; i--)
        {
            res.push_back(matrix[i][begin]);
        }
        
        game(res, matrix,begin+1,row-1, col-1);
    }
};