/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
/*
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> res;
        
        inorder(res, root);
        return res;
    }
    
    void inorder(vector<int> &res, TreeNode *root)
    {
        if(!root)
            return ;
        inorder(res, root->left);
        res.push_back(root->val);
        inorder(res, root->right);
    }
};
*/
//stack version
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> node;
        stack<TreeNode *> stk;
        TreeNode *cur = root;
        
        while(cur || !stk.empty())
        {
            if(cur)
            {
                stk.push(cur);
                cur = cur->left;
            }
            else
            {
                cur = stk.top();
                stk.pop();
                node.push_back(cur->val);
                cur = cur->right;
            }
        }
        
        return node;
    }
};