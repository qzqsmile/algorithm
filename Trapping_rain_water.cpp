class Solution {
public:
    int trap(vector<int>& height) {
        int i = 0;
        int j = height.size()-1;
        
        while((i <= j) && (height[i] == 0)) i++;
        while((i <= j) && (height[j] == 0) ) j--;
        int count = 0;
        while(i < j)
        {
            if(height[i] < height[j])
            {
                int p = i;
                while((height[p] >= height[i]) && (i < j))
                {
                    count += height[p] - height[i];
                    i++;
                }
            }
            else
            {
                int p = j;
                while((height[j] <= height[p]) && (i < j))
                {
                    count += height[p] - height[j];
                    j--;
                }
            }
        }
        
        return count;
    }
};