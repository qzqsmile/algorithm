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
    int minDepth(TreeNode* root) {
        if(!root)   return 0;
        int level = 0;
        
        queue<TreeNode*> q;
        q.push(root);
        
        while(q.size())
        {
            int num = q.size();
            level++;
            while(num--)
            {
                TreeNode *cur = q.front();
                q.pop();
                if(cur->left) q.push(cur->left);
                if(cur->right) q.push(cur->right);
                if((cur->left == NULL) && (cur->right == NULL)) return level;
            }
        }
    }
};