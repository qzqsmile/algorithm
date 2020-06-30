/**
 * Definition for an interval.
 * struct Interval {
 *     int start;
 *     int end;
 *     Interval() : start(0), end(0) {}
 *     Interval(int s, int e) : start(s), end(e) {}
 * };
 */
class Solution {
public:
    static bool cmp(const Interval &a, const Interval &b)
    {
        return a.start < b.start;
    }
    vector<Interval> merge(vector<Interval>& intervals) {
        vector<Interval> res;
        if(!intervals.size())
            return res;
        sort(intervals.begin(), intervals.end(), cmp);
        res.push_back(intervals[0]);
        
        for(int i = 1; i < intervals.size(); i++)
        {
            if(intervals[i].start <= res.back().end)
            {
                Interval temp(res.back().start, max(res.back().end, intervals[i].end));
                res.pop_back();
                res.push_back(temp);
            }
            else
                res.push_back(intervals[i]);
        }
        return res;
    }
};