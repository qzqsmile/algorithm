class Solution {
	public:
		int findMin(vector<int> &num) {
			int l, h, mid;
			l = 0;
			h = num.size()-1;

			while(l < h)
			{
				mid = (l + h) / 2;
				if(num[mid] < num[h])
					h = mid;
				else if(num[mid] > num[h])
					l = mid + 1;
				else
					h--;
			}

			return num[l];
		}
};
