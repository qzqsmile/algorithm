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
    ListNode *rotateRight(ListNode *head, int k) {
        int length = 0;
        ListNode *cur = head;
        ListNode *res = head;
        ListNode *pre = NULL;
        
        if(!head)
            return NULL;
        
        while(cur)
        {
            length++;
            pre = cur;
            cur = cur->next;
        }
        
        k = k % length;
        pre->next = head;
        cur = head;
        int cur_count = 0;
        
    
        while(cur)
        {
            cur_count++;
            if(cur_count == (length - k))
            {
                res = cur->next;
                cur->next = NULL;
                break;
            }
            cur = cur->next;
        }
        
        return res;
    }
};
