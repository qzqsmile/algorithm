class Solution {
public:
    bool isPalindrome(int x) {
        int i = 0;;
        if ((x % 10 == 0 && x != 0) || x < 0) return false;
        while (i < x) {
            i = i * 10 + x % 10;
            x = x / 10;
        }
        return (i == x || i / 10 == x);        
    }
};