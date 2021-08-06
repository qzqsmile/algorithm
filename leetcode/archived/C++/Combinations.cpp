class Solution {
	public:
		vector<vector<int> > combine(int n, int k) {
			vector<vector<int> >res;
			vector<int> v;

			makecombine(res,v,1,n,k);

			return res;
		}

		void makecombine(vector<vector<int> >&res, vector<int> &v,int pre, int n, int k)
		{
			if(v.size() == k)
			{
				res.push_back(v);
			}
			else
			{
				if(pre <= n)
				{
					v.push_back(pre);
					makecombine(res,v,pre+1,n,k);
					v.pop_back();
					makecombine(res,v,pre+1,n,k);
				}
			}
		}
};
