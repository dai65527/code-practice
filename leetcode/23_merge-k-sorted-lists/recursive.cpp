// https://leetcode.com/problems/merge-k-sorted-lists/submissions/987816787/
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* merge(ListNode* a, ListNode* b) {
        if (!a) {
            return b;
        }
        if (!b) {
            return a;
        }
        ListNode* head = a;
        ListNode* wait = b;
        if (a->val > b->val) {
            head = b;
            wait = a;
        }
        ListNode* node = head;
        while (node->next && wait) {
            while (node->next && node->next->val < wait->val) {
                node = node->next;
            }
            ListNode* tmp = node->next;
            node->next = wait;
            wait = tmp;
            node = node->next;
        }
        if (!node->next) {
            node->next = wait;
        }
        return head;
    }
    ListNode* merge2Lists(vector<ListNode*>&lists, int i, int j) {
        if (i == j) {
            return lists[i];
        }
        return merge(
            merge2Lists(lists, i, (i + j) / 2),
            merge2Lists(lists, (i + j) / 2 + 1, j));
    }
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        if (lists.empty()) {
            return nullptr;
        }
        return merge2Lists(lists, 0, lists.size() - 1);
    }
};
