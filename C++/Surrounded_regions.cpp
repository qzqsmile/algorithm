class Solution {
public:
    void solve(vector<vector<char>>& board) {
        if(!board.size()) return ;
        
        for(int i = 0; i < board.size(); i++)
        {
            for(int j = 0; j < board[0].size(); j++)
            {
                if((i == 0 || j == 0 || (i == board.size()-1) || (j == board[0].size()-1)) && (board[i][j] == 'O'))
                {
                    dfs(board, i, j);
                }
            }
        }
        
        for(int i = 0; i < board.size(); i++)
        {
            for(int j = 0; j < board[0].size(); j++)
            {
                if(board[i][j] == 'O')
                    board[i][j] = 'X';
                if(board[i][j] == '-')
                    board[i][j] = 'O';
            }
        }
    }
    
    void dfs(vector<vector<char>> & board, int row, int col)
    {
        board[row][col] = '-';
        if((row > 1) && (board[row-1][col] == 'O'))
            dfs(board, row-1, col);
        if((row < (board.size()-1)) && (board[row+1][col] == 'O'))
            dfs(board, row+1, col);
        if((col > 1) && (board[row][col-1] == 'O'))
            dfs(board, row, col-1);
        if((col < (board[0].size()-1)) && (board[row][col+1] == 'O'))
            dfs(board, row, col+1);
    }
};