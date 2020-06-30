class Solution {
public:
    vector<int> searchRange(int A[], int n, int target) {
        int low, high;
        vector<int> res;
        res.push_back(-1);
        res.push_back(-1);
        low = 0; high = n-1;
        //
        while(low <= high)
        {
            int mid = (low + high) / 2;
            if(A[mid] > target)
            {
                high = mid - 1;
            }
            else if(A[mid] < target)
            {
                low = mid + 1;
            }
            else
            {
                if(A[low] == target)
                {
                    res.clear();
                    res.push_back(low);
                    break;
                }
                else
                    low++;
            }
        }
        
        low = 0; high = n-1;
        while(low <= high)
        {
            int mid = (low + high) / 2;
            if(A[mid] > target)
            {
                high = mid - 1;
            }
            else if(A[mid] < target)
            {
                low = mid + 1;
            }
            else
            {
                if(A[high] == target) 
                {
                    res.push_back(high);
                    break;
                }
                else
                {
                    high--;
                }
            }
        }
        
        return res;
    }
};