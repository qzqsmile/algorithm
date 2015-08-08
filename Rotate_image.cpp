class Solution {
	public:
		/*clockwise rotate
		 * first reverse up to down, then swap the symmetry 
		 * 1 2 3     7 8 9     7 4 1
		 * 4 5 6  => 4 5 6  => 8 5 2
		 * 7 8 9     1 2 3     9 6 3
		 *
		 *����Ĺؼ��������˳ʱ����ת����ͨ����Щ�任����Ч�õ�
		 *���½������Կ��� 90->90->���ҽ���
		 *��б�Խ��Ƚ������Կ��� ���ҽ���->��ʱ����ת90
		 *��˿��Ե�ЧΪ����������
		*/
		void rotate(vector<vector<int> > &matrix) {
			 reverse(matrix.begin(), matrix.end());
			 for (int i = 0; i < matrix.size(); ++i) {
				 for (int j = i + 1; j < matrix[i].size(); ++j)
					 swap(matrix[i][j], matrix[j][i]);
			 }
		 }
};
