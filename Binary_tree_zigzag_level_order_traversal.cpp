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
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int>>res;
        if(!root) return res;
        queue<TreeNode *> q;
        q.push(root);
        
        bool flag = 1;
        while(!q.empty())
        {
            vector<int> row;
            int row_count = q.size();
            flag = flag ^ 1;
            while(row_count)
            {
                TreeNode *t = q.front();
                row.push_back(t->val);
                //if(flag)
                //{
                if(t->left) q.push(t->left);
                if(t->right) q.push(t->right);
                //}
                /*else
                {
                    if(t->right) q.push(t->right);
                    if(t->left) q.push(t->left);
                }*/
                q.pop();
                row_count--;
            }
            if(flag) 
                reverse(row.begin(), row.end());
            res.push_back(row);
        }
        
        return res;
    }
};