#include<iostream>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};


class Solution {
	public:
		ListNode *reverseBetween(ListNode *head, int m, int n) {
			if(!head)
				return NULL;

			ListNode* rever_head = NULL, *rever_pre = NULL, *rever_cur = NULL;
			ListNode* cur = head;
			ListNode* pre = head;
			int count = 0;

			while(cur)
			{
				count++;
				if(count == (m-1))
					pre = cur;
				if(count == m)
					rever_head = rever_pre = rever_cur = cur;
				if((count > m) && (count <= n))
				{
					rever_cur = cur;
					rever_pre->next = rever_cur->next;
					rever_cur->next = rever_head;
					rever_head = rever_cur;
					cur = rever_pre;
				}
				cur = cur->next;            
			}

			if(pre != head)
				pre->next = rever_head;

			return head;
		}
};

int main(void)
{
	Solution s;
	ListNode *head = NULL;
	ListNode *res = NULL;
	head = new ListNode(1);
	head->next = new ListNode(2);
	head->next->next = new ListNode(3);

	res = s.reverseBetween(head, 2, 3);

	return 0;
}
