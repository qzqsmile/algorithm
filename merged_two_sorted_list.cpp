class Solution {
	public:
		ListNode *mergeTwoLists(ListNode *l1, ListNode *l2) {
			ListNode *res = NULL;
			if((l1 == NULL) && (l2 == NULL))
				return NULL;
			if(l1 == NULL)
				return l2;
			if(l2 == NULL)
				return l1;
			ListNode* cur1 = l1;
			ListNode* cur2 = l2;
			ListNode* res_cur = NULL;

			while(cur1)
			{
				while((cur2) &&((cur1->val) > (cur2->val)))
				{
					if(!res)
						res = res_cur = new ListNode(cur2->val);
					else
					{
						res_cur->next = new ListNode(cur2->val);
						res_cur = res_cur->next;
					}
					cur2 = cur2->next;
				}

				if(!res)
					res = res_cur = new ListNode(cur1->val);
				else
				{
					res_cur->next = new ListNode(cur1->val);
					res_cur = res_cur->next;
				}

				cur1 = cur1->next;
			}
			if(cur2)
			{
				res_cur->next = cur2;
			}


			return res;
		}
};
