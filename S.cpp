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
// recursive version
/*
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int>res;
        traversal(res, root);
        return res;
    }
    
    void traversal(vector<int> &res, TreeNode *root)
    {
        if(!root)   return ;
        res.push_back(root->val);
        traversal(res, root->left);
        traversal(res, root->right);
    }*/
    
//iteratively version
    vector<int> preorderTraversal(TreeNode* root){
       vector<int>res;
       stack<TreeNode *>q;
       q.push(root);
       
       while(!q.empty())
       {
            TreeNode *s = q.top();
            q.pop();
            if(s == NULL)
                continue;
            res.push_back(s->val);
            q.push(s->right);
            q.push(s->left);
       }
       
       return res;
    }
};