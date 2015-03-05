class Solution {
public:
    void rotate(int nums[], int n, int k) {
        /*solution 1
        int *p = new int [n];
        int cur = 0;
        if(k > n)
            k = k % n;
        for(int i = 0; i != k;  i++)
        {
            p[cur++] = nums[n-k+i];
        }
        for(int i = 0; i != (n-k); i++)
        {
            p[cur++] = nums[i];
        }
        
        for(int i = 0; i != n; i++)
        {
            nums[i] = p[i];
        }
        */
        /*solution 2
        reverse(nums,nums+n);
        reverse(nums,nums+k%n);
        reverse(nums+k%n,nums+n);
        */
        
        //solution 3
        /*
		at first, nums[] is [1,2,3,4,5,6,7], n is 7, k is 3
		after first outer loop, nums[] is [4,1,2,3], n is 4, k is 3
		after second outer loop, nums[] is [4], n is 1, k is 0
		loop ends.
		*/
        for (; k %= n; n -= k)
            for (int i = 0; i < k; i++)
                swap(*nums++, nums[n - k]);
    }
};