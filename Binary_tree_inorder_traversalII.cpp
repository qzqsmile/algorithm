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
    vector<vector<int>> levelOrderBottom(TreeNode* root) {
        vector<vector<int> > res;
        queue<TreeNode *> q;
        if(!root)
            return res;
        q.push(root);
        while(!q.empty())
        {
            vector<int> level;
            int size = q.size();
            for(int i = 0; i < size; i++)
            {
                TreeNode *node;
                node = q.front();
                q.pop();
                level.push_back(node->val);
                if(node->left != NULL)
                    q.push(node->left);
                if(node->right != NULL)
                    q.push(node->right);
            }
            
            res.push_back(level);
        }
        reverse(res.begin(), res.end());
        return res;
    }
};