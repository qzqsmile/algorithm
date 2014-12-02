class Solution {
	public:
		ListNode *insertionSortList(ListNode *head) {
			ListNode *cur = head;
			ListNode *res = NULL;
			ListNode *res_cur = NULL;

			if(!head)
				return NULL;
			while(cur)
			{
				if(!res)
					res = res_cur = new ListNode(cur->val);
				else
				{
					ListNode *res_pre = NULL;
					ListNode *temp = NULL;
					res_cur = res;
					while((res_cur) && ((res_cur->val) < (cur->val)))
					{
						res_pre = res_cur;
						res_cur = res_cur->next;
					}

					if(!res_pre)
					{
						temp = res;
						res = new ListNode(cur->val);
						res->next = temp;
					}
					else{
						temp = res_pre->next;
						res_pre->next = new ListNode(cur->val);
						res_pre->next->next = temp;
					}
				}
				cur = cur->next;
			}
			return res;
		}
};
