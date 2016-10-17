#include<vector>
#include<iostream>

using namespace std;

struct ListNode{
	int val;
	ListNode *next;
	ListNode(int x) : val(x), next(NULL){ };
};

class Solution {
	public:
		ListNode *mergeKLists(vector<ListNode *> &lists) {
			return reverList(lists, lists.size());
		}

		ListNode *reverList(vector<ListNode *> &lists, int size)
		{
			ListNode *left = lists[0], *right = lists[1];
			ListNode *res = NULL;
			size = lists.size();
			if(!size)
				return NULL;
			if(size == 1)
				return lists[0];

			if(size > 2)
			{
				left = reverList(lists, size /2);
				right = reverList(lists[size / 2], size - (size / 2));
			}

			ListNode *left_cur = left, *right_cur = right;
			ListNode *cur = left_cur;

			while(left_cur)
			{
				if((!right_cur)||(left_cur->val) < (right_cur->val))
				{
					if(!res)
						res = cur = new ListNode(left_cur->val);
					else
					{
						cur->next = new ListNode(left_cur->val);
						cur = cur->next;
					}
					left_cur = left_cur->next;
				}
				else
				{
					if(!res)
						res = cur = new ListNode(right_cur->val);
					else
					{
						cur->next = new ListNode(right_cur->val);
						cur = cur->next;
					}
					right_cur = right_cur->next;
				}
			}
			if(!right_cur)
			{
				cur->next = right_cur;
			}
			return res;
		}
};


int main(void)
{
	return 0;
}
