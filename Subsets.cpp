class Solution {
	public:
		vector<vector<int> > subsets(vector<int> &S) {
			vector<vector<int> >res;
			vector<int> p;

			sort(S.begin(), S.end());
			makearray(res, S, p, 0, S.size());

			return res;
		}

		void makearray(vector<vector<int> >&res, vector<int> &s, vector<int> &p, int pre, int end)
		{
			if(pre >= end){
				res.push_back(p);
				return;
			}
			else{
				p.push_back(s[pre]);
				makearray(res, s, p, pre+1, end);
				p.pop_back();
				makearray(res, s, p, pre+1, end);
			}
		}
};
