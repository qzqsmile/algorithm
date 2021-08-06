/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    bool isValidBST(TreeNode* root) {
        return isValidBST(root,(long int)INT_MIN - 1, (long int)INT_MAX + 1);
    }
    
    bool isValidBST(TreeNode *root, long int min, long int max)
    {
        if(!root)   return true;
        if((root->val <= min) || (root->val >= max)) return false;
        return isValidBST(root->left, min, root->val) && isValidBST(root->right, root->val, max);
    }
    
    /*
    另一种解法，这种解法利用判断minNode和maxNode指针是否为NULL来确定是否初始化过，从而
    解决了INT_MAX INT_MIN的溢出问题
    bool isValidBST(TreeNode* root) {
    return isValidBST(root, NULL, NULL);
}

    bool isValidBST(TreeNode* root, TreeNode* minNode, TreeNode* maxNode) {
        if(!root) return true;
        if(minNode && root->val <= minNode->val || maxNode && root->val >= maxNode->val)
            return false;
    return isValidBST(root->left, minNode, root) && isValidBST(root->right, root, maxNode);
}
    */
};