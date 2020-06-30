#include<iostream>
#include<string>
#include<vector>

using namespace std;

struct  ListNode{
	int val;
	ListNode *next;
	ListNode(int x) : val(x), next(NULL){}
};

class Solution {
	public:
		void reorderList(ListNode *head) {
			ListNode *cur = head;
			vector<int> res;

			if(!cur)
				return;

			while(cur)
			{
				res.push_back(cur->val);
				cur = cur->next;
			}

			for(int i = 0; i <=  (res.size()+1)/2; i++)
			{
				if(i==0)
				{
					head->val = res[i];
					cur = head;
					if(i < (res.size()-i-1))
					{
						cur->next = new ListNode(res[res.size()-1]);
						cur = cur->next;
					}
				}
				else
				{
					if(i < (res.size()-i-1))
					{
						cur->next = new ListNode(res[i]);
						cur = cur->next;
						cur->next = new ListNode(res[res.size()-i]);
						cur = cur->next;
					}
					if(i == (res.size()-i-1))
					{
						cur ->next = new ListNode(res[i]);
						cur = cur->next;
					}
				}
			}

			cur->next = NULL;
		}
};

int main(void)
{
	Solution s;
	
	ListNode *head;
	head = new ListNode(1);
	s.reorderList(head);

	cout << head->val  << head->next->val << head->next->next->val << endl; 

	return 0;
}
