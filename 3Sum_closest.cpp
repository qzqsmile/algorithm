class Solution {
	public:
		int threeSumClosest(vector<int> &num, int target) {
			sort(num.begin(), num.end());
			int closest_sum = 0;
			closest_sum = num[0] + num[1] + num[2];
			for(int i = 0; i != num.size()-2; i++)
			{
				int j, k;
				j = i + 1;
				k = num.size()-1;

				while(j < k)
				{
					int sum = num[i] + num[j] + num[k];
					if(sum == target) return target;
					if(sum > target) 
					{
						if((sum-target) < fabs(closest_sum-target))
							closest_sum = sum;
						k--;
					}
					else
					{
						if((target-sum) < fabs(closest_sum-target))
							closest_sum = sum;
						j++;
					}
				}
			}

			return closest_sum;

		}
};
