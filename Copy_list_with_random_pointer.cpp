/*Defintion for singly-linked list with a random pointer.
 * struct RandomListNode {
 *     int label;
 *     RandomListNode *next, *random;
 *     RandomListNode(int x) : label(x), next(NULL), random(NULL) {}
 * };
 */
class Solution {
	public:
		RandomListNode *copyRandomList(RandomListNode *head) {

			RandomListNode *res = NULL, *cur = head;
			if(!cur)
				return NULL;

			//insert one node betweent two nodes
			while(cur)
			{
				RandomListNode *temp = cur->next;
				cur->next = new RandomListNode (cur->label);
				cur->next->next = temp;
				cur = temp;
			}

			cur = head;

			//copy random pointer
			while(cur)
			{
				if((cur->random))
					cur->next->random = cur->random->next;
				else
					cur->next->random = NULL;
				cur = cur->next->next;
			}

			cur = head;
			res = cur->next;

			//remove duplicate ndoes
			while(cur)
			{
				RandomListNode *cur1 = cur->next;
				cur->next = cur->next->next;
				cur = cur->next;
				if(cur)
					cur1->next = cur->next;
				else
					cur1->next = NULL;
			}

			return res;
		}
};
