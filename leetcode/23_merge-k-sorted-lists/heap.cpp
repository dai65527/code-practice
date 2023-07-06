// https://leetcode.com/problems/merge-k-sorted-lists/submissions/987799843/
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
struct compare {
    bool operator()(const ListNode* a, const ListNode* b) {
        return a->val > b->val;
    }
};
class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        priority_queue<ListNode*, vector<ListNode*>, compare> q;
        for (int i = 0; i < lists.size(); i++) {
            if (lists[i]) {
                q.push(lists[i]);
            }
        }
        if (q.empty()) {
            return nullptr;
        }
        ListNode* head = q.top();
        q.pop();
        ListNode* node = head;
        while (!q.empty()) {
            while (node->next && node->next->val < q.top()->val) {
                node = node->next;
            }
            ListNode* tmp = node->next;
            node->next = q.top();
            q.pop();
            if (tmp) {
                q.push(tmp);
            }
            node = node->next;
        }
        return head;
    }
};
