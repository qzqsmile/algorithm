class Solution {
	public:
		vector<vector<int> > threeSum(vector<int> &num) {
			vector<vector<int> > res;
			if(num.size() < 3)
				return res;

			sort(num.begin(), num.end());

			for(int i = 0; i < (num.size()-2); i++)
			{
				int l, h;
				int target = -num[i];
				l = i + 1;
				h = num.size()-1;

				while(l < h)
				{
					int sum = num[l] + num[h];
					if(sum < target)
						l++;
					else if(sum > target)
						h--;
					else
					{
						vector<int> triplet(3,0);
						triplet[0] = num[i];
						triplet[1] = num[l];
						triplet[2] = num[h];
						res.push_back(triplet);

						while((l < h) && (num[l] == triplet[1])) l++;

						while((l < h) && (num[h] == triplet[2])) h--;
					}
				}

				while ((i + 1 < num.size()) && (num[i + 1] == num[i])) 
					i++;

			}

			return res;
		}
};
