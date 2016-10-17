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
   TreeNode *s1,*s2,*pre;
   
   void recover(TreeNode *root)
   {
       if(!root) return ;
       recover(root->left);
       if(pre && (pre->val > root->val))
       {
            if(s1 == NULL)
            {
                s1 = pre;
                s2 = root;
            }
            else
            {
                s2 = root;
            }
       }
       pre = root;
       recover(root->right);
   }
    
    void recoverTree(TreeNode *root) {  
        s1 = s2 = pre = NULL;
        recover(root);
        swap(s1->val, s2->val);
    }  
};