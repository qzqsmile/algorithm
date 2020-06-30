class Solution {
	public:
		bool exist(vector<vector<char> > &board, string word) {
			if((!word.size()) || (!board.size()))
				return false;
			for(int i = 0; i < board.size(); i++)
			{
				for(int j = 0; j < board[0].size(); j++)
					if(searchword(board, i, j, word, 0))
						return true;
			}
			return false;
		}

		//dfs
		bool  searchword(vector<vector<char> > &board, int row, int col, string &word, int index)
		{
			if(board[row][col] == word[index])
			{
				if(index == (word.size()-1))
					return true;
				char temp = board[row][col];
				board[row][col] = '.';

				//不要单独判定条件，那样就不是dfs，会TLE   
				if(((row > 0) && (searchword(board, row-1, col, word, index+1)))
						||((row < (board.size()-1))&&(searchword(board, row+1, col, word, index+1)))
						||((col > 0) && (searchword(board, row, col-1, word, index+1)))
						||((col < (board[0].size()-1)) && (searchword(board, row, col+1, word, index+1))))
					return true;

				board[row][col] = temp;
			}

			return false;

		}
};
