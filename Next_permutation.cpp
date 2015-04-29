#include<iostream>
#include<vector>
#include<algorithm>

using namespace std;

class Solution {
	public:
		void nextPermutation(vector<int> &num) {
			for(int i = num.size()-1; i>= 0; i--)
			{
				for(int j = i -1; j >= 0; j--)
				{
					if(num[i] > num[j])
					{
						swap(num[i], num[j]);
						sort(num.begin()+j+1, num.end());
						return ;
					}
				}
			}

			sort(num.begin(), num.end());
		}
};

int main(void)
{
	Solution s;
	int a[7] = {4, 2, 0, 2, 3, 2, 0};
	vector<int> num(a, a+7);

	s.nextPermutation(num);

	return 0;
}
