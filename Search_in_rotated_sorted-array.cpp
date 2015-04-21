class Solution {
public:
    int search(int A[], int n, int target) {
        int low = 0, high= n-1, lowest = 0;
        
        while(low < high)
        {
            int mid = (low + high) / 2;
            if(A[mid] > A[high])
                low = mid+1;
            else
                high = mid; 
        }
        
        lowest = low;
        low = 0; high = n-1;
        while(low <= high)
        {
            int mid = (low + high) / 2;
            int realmid = (mid+lowest) % n;
            if(A[realmid] == target) 
                return realmid;
            if(A[realmid] < target)low=mid+1;
            else high = mid-1;
        }
        
        return -1;
    }
};