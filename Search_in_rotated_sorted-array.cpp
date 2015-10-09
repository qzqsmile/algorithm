class Solution {
public:
    int search(vector<int>& nums, int target) {
        int begin = 0, end = nums.size()-1;
        
        while(begin <= end){
            int mid = (begin + end) / 2;
            if(nums[mid] == target)
                return mid;
            //if enter this case, impossible enter another case
            if(nums[mid] >= nums[begin]){
                if(nums[begin] <= target && target < nums[mid])
                    end = mid -1;
                else
                    begin = mid + 1;
            }
            
            if(nums[mid] <= nums[end]){
                if(nums[mid] < target && target <= nums[end])
                    begin = mid + 1;
                else
                    end = mid - 1;
            }
        }
        
        return -1;
    }
};