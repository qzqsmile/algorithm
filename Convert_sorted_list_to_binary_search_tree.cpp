/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
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
    TreeNode* sortedListToBST(ListNode* head) {
        TreeNode *root = NULL;
        ListNode *cur = head;
        int len = 0;
        
        while(cur)
        {
            len++;
            cur = cur->next;
        }
        root = build(head, 1, len);
        
        return root;
    }
    
    TreeNode *build(ListNode *head, int start, int end)
    {
        if(start > end) 
            return NULL;
        TreeNode *root = NULL;
        int mid = start + (end - start) / 2 ;
        ListNode *cur = head;
        int tmp = start;
        while(cur)
        {
            if(tmp == mid)
            {
                root = new TreeNode(cur->val);
                break;
            }
            tmp++;
            cur = cur->next;
        }
        root->left = build(head, start, mid-1);
        if(cur)
            root->right = build(cur->next, mid+1, end);
        
        return root;
    }
};