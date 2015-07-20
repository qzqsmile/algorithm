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
    ListNode* swapPairs(ListNode* head) {
        if(!head)
            return NULL;
        if(!(head->next))
            return head;
        ListNode *node1 = NULL, *node2 = NULL, *temp = NULL;
        node1 = head;
        node2 = head->next;
        
        temp = node2->next;
        node2->next = node1;
        node1->next = swapPairs(temp);
        
        return node2;
    }
};