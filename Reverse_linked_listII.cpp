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
    ListNode* reverseBetween(ListNode* head, int m, int n) {
        ListNode *pre = NULL, *last = NULL, *cur = head;
        ListNode *cir_pre = NULL, *cir_end = NULL;
        int index = 0;
        
        while(cur)
        {
            index++;
            if(index == (m-1))
                pre = cur;
            if(index == (n+1))
                last = cur;
            if(index == m)
                cir_pre = cir_end = cur;
            if((m < index) && (index <= n))
            {
                ListNode *t = cur->next;
                cur->next = cir_pre;
                cir_pre = cur;
                cur = t;
                continue;
            }
            cur = cur->next;
        }
        
        if(pre)
            pre->next = cir_pre;
        cir_end->next = last;
        
        if(1 == m)
            return cir_pre;
        return head;
    }
};