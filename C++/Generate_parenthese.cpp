class Solution {
	public:
		vector<string> generateParenthesis(int n) {
			vector<string> res;
			generate(res,"",n, n);
			return res;
		}

		void generate(vector<string>&res, string s, int left, int right)
		{
			if((left == 0) && (right == 0)) res.push_back(s);
			if(left > 0) generate(res, s+"(", left-1, right);
			if(right > 0 && (right > left)) generate(res, s+")", left, right-1);
		}
};
