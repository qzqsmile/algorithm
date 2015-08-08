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
    //这题关键是维护一个动态的Interval
    vector<Interval> insert(vector<Interval>& intervals, Interval newInterval) {
       vector<Interval> ret;
       auto it = intervals.begin();
       for(;it != intervals.end(); it++)
       {
            if(newInterval.end < (*it).start)
                break;
            else if(newInterval.start > (*it).end)
                ret.push_back(*it);
            else{
                newInterval.start = min((*it).start, newInterval.start);
                newInterval.end = max((*it).end, newInterval.end);
            }
       }
       
       ret.push_back(newInterval);
       for(; it != intervals.end(); it++)
       {
           ret.push_back(*it);
       }
        
       return ret;
    }
};