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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *cur1  = l1, *cur2 = l2, *cur3, *res;
        int flag, sum;
        
        if(cur1 && cur2)
        {
            sum = cur1->val + cur2->val;
            res = cur3 = new ListNode(sum % 10);
            flag = sum / 10;
            cur1 = cur1->next;
            cur2 = cur2->next;
        }
        
        while(cur1 && cur2)
        {
            sum = cur1->val + cur2->val + flag;
            cur3->next = new ListNode(sum % 10);
            flag = sum / 10;
            cur3 = cur3->next;
            cur1 = cur1->next;
            cur2 = cur2->next;
        }
        
        while(cur1)
        {
            sum = cur1->val + flag;
            cur3->next = new ListNode(sum % 10);
            flag = sum / 10;
            cur3 = cur3->next;
            cur1 = cur1->next;
        }
        
        while(cur2)
        {
            sum = cur2->val + flag;
            cur3->next = new ListNode(sum % 10);
            flag = sum / 10;
            cur3 = cur3->next;
            cur2 = cur2->next;
        }
        
        if(flag)
            cur3->next = new ListNode(flag);
        
        return res;
    }
};