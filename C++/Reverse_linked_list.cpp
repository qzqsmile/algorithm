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
//recursively version
 /*   ListNode* reverseList(ListNode* head) {
        if(head == NULL)
            return NULL;
        if(head->next == NULL)
            return head;
        ListNode *head1 = NULL, *cur = NULL;
        
        head1 = reverseList(head->next);
        cur = head1;
        while(cur->next)
        {
            cur = cur->next;
        }
        cur->next = head;
        head->next = NULL;
        
        return head1;
    }
*/
//iteratively version
    ListNode *reverseList(ListNode* head){
        ListNode *first = NULL, *cur = NULL;
        if(!head)
            return NULL;
        first = head;
        cur = head->next;
        head->next = NULL;
        while(cur)
        {
            ListNode *cur1 = cur->next;
            cur->next = first;
            first = cur;
            cur = cur1;
        }
        
        return first;
    }
 };