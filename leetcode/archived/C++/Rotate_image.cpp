class Solution {
	public:
		/*clockwise rotate
		 * first reverse up to down, then swap the symmetry 
		 * 1 2 3     7 8 9     7 4 1
		 * 4 5 6  => 4 5 6  => 8 5 2
		 * 7 8 9     1 2 3     9 6 3
		 *
		 *这题的关键在于理解顺时针旋转可以通过那些变换来等效得到
		 *上下交换可以看作 90->90->左右交换
		 *而斜对角先交换可以看作 左右交换->逆时针旋转90
		 *因此可以等效为这两个操作
		*/
		void rotate(vector<vector<int> > &matrix) {
			 reverse(matrix.begin(), matrix.end());
			 for (int i = 0; i < matrix.size(); ++i) {
				 for (int j = i + 1; j < matrix[i].size(); ++j)
					 swap(matrix[i][j], matrix[j][i]);
			 }
		 }
};
