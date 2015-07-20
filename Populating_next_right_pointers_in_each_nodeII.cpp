/**
 * Definition for binary tree with next pointer.
 * struct TreeLinkNode {
 *  int val;
 *  TreeLinkNode *left, *right, *next;
 *  TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
 * };
 */
class Solution {
public:
    void connect(TreeLinkNode *root) {
        if(!root)
            return ;
        queue<TreeLinkNode *>q;
        q.push(root);
        
        while(q.size())
        {
            int num = q.size();
            TreeLinkNode * cur = NULL;
            while(num--)
            {
                if(!cur) cur = q.front();
                else cur->next = q.front();
                q.pop();
                if(cur->next) cur = cur->next;
                if(cur->left) q.push(cur->left);
                if(cur->right) q.push(cur->right);
            }
            cur->next = NULL;
        }
    }
};