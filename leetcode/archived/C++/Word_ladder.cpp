#include<iostream>
#include<vector>
#include<string>
#include<unordered_set>

using namespace std;

class Solution {
	public:
		int ladderLength(string start, string end, unordered_set<string> &dict) {
			int short_path = dict.size(), flag = 0;

			for(auto t = dict.begin(); t != dict.end(); t++)
			{
				if(isOneLetter(end, *t))
				{
					int store = ladderLength(start, *t, dict);
					if((short_path > (store + 1)) && ( store != 0))
					{
						short_path = store + 1;
						flag = 1;
					}
				}
			}

			if(flag)
				return short_path;
			else 
				return 0;
		}

		bool isOneLetter(string a, string b)
		{
			int count = 0;
			for(int i = 0; i < a.size(); i++)
			{
				if(a[i] != b[i])
					count++;
			}
			if(count == 1)
				return true;
			return false;
		}
};

int main(void)
{
	unordered_set<string> dict{"a","b","c"};
	Solution s;
	cout << s.ladderLength("a","c",dict) << endl;
	return 0;
}

