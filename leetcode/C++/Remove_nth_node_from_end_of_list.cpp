/*Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
	public:
		ListNode* removeNthFromEnd(ListNode* head, int n) {
			ListNode *first = head;
			ListNode *second = NULL;
			if(!head || (n==0)) return head;

			int index = 0;
			while(first)
			{
				index++;
				if(second)
					second = second->next;
				if(index == (n+1))
					second = head;
				first = first->next;
			}
			//这里主要判断n大于链表的情形
			if(!second)
				return head->next;
			second->next = second->next->next;

			return head;
		}
};
