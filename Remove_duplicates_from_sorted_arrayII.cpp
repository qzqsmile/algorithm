class Solution {
public:
    int removeDuplicates(int A[], int n) {
        int *cur, *find_pre, *find_next;
        cur = find_pre = find_next = A;
        
        while((find_next-A) < n)
        {
            int count = 0;
            while((*find_next == *find_pre) && ((find_next-A) < n))
            {
                find_next++;
                count++;
            }
            *cur++ = *find_pre;
            if(count >= 2)
                *cur++ = *find_pre;
            find_pre = find_next;
        }
        
        return (cur - A);
    }
};