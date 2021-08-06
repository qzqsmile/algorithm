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
    bool isBalanced(TreeNode* root) {
        if(!root)
            return true;
        return abs(Calheight(root->left) - Calheight(root->right)) <= 1 && isBalanced(root->left) && isBalanced(root->right);
    }
    
    int Calheight(TreeNode *node)
    {
        if(node)
            return (1 + max(Calheight(node->left), Calheight(node->right)));
        return 0;
    }
};