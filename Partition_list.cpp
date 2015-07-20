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
    ListNode* partition(ListNode* head, int x) {
        ListNode *lowbegin = NULL, *highbegin = NULL,*low = NULL, *high = NULL, *cur = head;
        
        while(cur)
        {
            if(cur->val < x)
            {
                if (low == NULL)
                    lowbegin = low = new ListNode(cur->val);
                else
                {
                    low->next = new ListNode(cur->val);
                    low = low->next;
                }
            }
            else
            {
                if(high == NULL)
                    highbegin = high = new ListNode(cur->val);
                else
                {
                    high->next = new ListNode(cur->val);
                    high = high->next;
                }
            }
            cur = cur->next;
        }
        if(low == NULL || high == NULL)
            return head;
        else
        {
            low->next = highbegin;
            return lowbegin;
        }
    }
};