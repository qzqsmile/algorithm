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
    vector<vector<int>> pathSum(TreeNode* root, int sum) {
        vector<vector<int>>res;
        vector<int>r1;
        Sum(root, sum, res, r1);
        
        return res;
    }
    
    void Sum(TreeNode *root, int sum, vector<vector<int> > &res, vector<int>&r1)
    {
        if(!root) return ;
        r1.push_back(root->val);
        if(root->val == sum && root->left == NULL && root->right == NULL)
            res.push_back(r1);
        Sum(root->left, sum-(root->val), res, r1);
        Sum(root->right, sum-(root->val), res, r1);
        r1.pop_back();
    }
};