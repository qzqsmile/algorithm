class Solution {
public:
    vector<int> findSubstring(string s, vector<string>& words) {
        vector<int> res;
        int n = s.size(), cnt = words.size();
        if (n <= 0 || cnt <= 0) return res;
        
        int len = words[0].size();
        unordered_map<string,int>dict;
        for(int i = 0; i < words.size(); i++)   dict[words[i]]++;
        int scope = n-cnt*len+1;
        
        for(int i = 0; i < scope; i++)
        {
            unordered_map<string, int>dict1;
            int scope1 = cnt*len + i;
            for(int j = i; j < scope1; j += len)
            {
                string s1 = s.substr(j, len);
                if(dict.count(s1))
                    dict1[s1]++;
                else
                    break;
            }
            if(dict1 == dict)
                res.push_back(i);
            dict1.clear();
        }
        
        return res;
    }
};