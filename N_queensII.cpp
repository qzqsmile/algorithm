class Solution {
	public:
		int limit, count;


		void totalNQueens(vector<int> &arr, int r)
		{
			if(r == limit)
				count++;
			else
			{
				for(int col = 0; col < limit; col++)
				{
					arr[r] = col;
					if(isSafe(arr, r, col))
						totalNQueens(arr, r+1);
				}
			}
		}

		bool isSafe(vector<int> &arr, int r, int c)
		{

			for(int row = r-1, ldia = c-1, rdia = c + 1; row >= 0; row--, ldia--, rdia++)
			{
				int col = arr[row];
				if(col == c || col == ldia || col == rdia)
					return false;
			}
			return true;
		}

		int totalNQueens(int n) {
			vector<int>arr(n, 0);
			limit = n;
			count = 0;
			totalNQueens(arr, 0);

			return count;
		}
};
