class Solution {
	public:
		void nextPermutation(vector<int> &num) {
			if(num.size() <= 1) return ;
			int i = num.size()-1, j;

			for(j = num.size()-2; j >= 0; j--)
			{
				if(num[j] < num[j+1]) break;
			}
			if(j>=0){
				while(num[i]<=num[j]) i--;
				swap(num[i], num[j]);
			}

			sort(num.begin()+j+1, num.end());
		}
};
