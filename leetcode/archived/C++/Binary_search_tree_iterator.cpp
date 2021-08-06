/**
*Defition for binary tree
* struct TreeNode {
	*     int val;
	*     TreeNode *left;
	*     TreeNode *right;
	*     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
	* };
*/

class BSTIterator {
	public:
		//    int is_small_exist;
		stack<TreeNode *> track;
		BSTIterator(TreeNode *root) {
			if(root)
			{
				TreeNode *p = root;
				while(p)
				{
					track.push(p);
					p = p->left;
				}
			}
		}

		/** @return whether we have a next smallest number */
		bool hasNext() {
			return (track.size() > 0);
		}

		/** @return the next smallest number */
		int next() {
			if(track.size() == 0) throw exception();
			TreeNode *p = track.top();
			track.pop();
			TreeNode *cur = p->right;
			while(cur)
			{
				track.push(cur);
				cur = cur->left;
			}
			return p->val;
		}
};
/**
 * Your BSTIterator will be called like this:
 * BSTIterator i = BSTIterator(root);
 * while (i.hasNext()) cout << i.next();
 */
