/**
 * Definition for binary tree
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    vector<int> rightSideView(TreeNode *root) {
        vector<int> res;
        rightsearch(root,res,0);    
        
        return res;
    }
    
    void rightsearch(TreeNode *root, vector<int> & res, int level)
    {
        if(root == NULL)
            return ;
        if(res.size() == level)
            res.push_back(root->val);
        rightsearch(root->right, res, level+1);
        rightsearch(root->left, res, level+1);
    }
};