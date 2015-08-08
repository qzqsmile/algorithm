/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if(!head) return NULL;
        if(!(head->next)) return head;
        ListNode *cur = head;
        ListNode *curnext = head->next;
        
        while(curnext &&(cur->val) == (curnext->val)) 
            curnext = curnext->next;
        cur->next = deleteDuplicates(curnext);
        
        return cur;
    }
};