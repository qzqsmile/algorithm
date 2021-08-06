class Solution {
	public:
		int findPeakElement(const vector<int> &num) {
			int l, h, mid;
			l = 0;
			h = num.size()-1;

			while(l < h)
			{
				mid = (l + h) / 2;
				if(num[mid] > num[mid+1])
					h = mid;
				else if(num[mid] < num[mid+1])
					l = mid + 1;
			}       

			return l;
		}
};
