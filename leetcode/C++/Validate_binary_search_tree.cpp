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
    bool isValidBST(TreeNode* root) {
        return isValidBST(root,(long int)INT_MIN - 1, (long int)INT_MAX + 1);
    }
    
    bool isValidBST(TreeNode *root, long int min, long int max)
    {
        if(!root)   return true;
        if((root->val <= min) || (root->val >= max)) return false;
        return isValidBST(root->left, min, root->val) && isValidBST(root->right, root->val, max);
    }
    
    /*
    ��һ�ֽⷨ�����ֽⷨ�����ж�minNode��maxNodeָ���Ƿ�ΪNULL��ȷ���Ƿ��ʼ�������Ӷ�
    �����INT_MAX INT_MIN���������
    bool isValidBST(TreeNode* root) {
    return isValidBST(root, NULL, NULL);
}

    bool isValidBST(TreeNode* root, TreeNode* minNode, TreeNode* maxNode) {
        if(!root) return true;
        if(minNode && root->val <= minNode->val || maxNode && root->val >= maxNode->val)
            return false;
    return isValidBST(root->left, minNode, root) && isValidBST(root->right, root, maxNode);
}
    */
};