class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        if(!matrix.size()) return false;
        
        int begin = 0, end = matrix.size()-1;
        
        while(begin <= end)
        {
            int mid = (begin + end) / 2;
            if(matrix[mid][0] > target)
                end = mid-1;
            else if(matrix[mid][0] < target)
                begin = mid+1;
            else
                return true;
        }
        
        int index = max(0,min(begin, end));
        end = matrix[index].size() - 1;
        begin = 0;
        while(begin <= end)
        {
            int mid = (begin + end) / 2;
            if(matrix[index][mid] < target)
                begin = mid + 1;
            else if(matrix[index][mid] > target)
                end = mid-1;
            else 
                return true;
        }
        
        return false;
    }
};