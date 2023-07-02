// https://leetcode.com/problems/remove-nth-node-from-end-of-list/submissions/984594120/
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
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode* lastNode = head;
        while (n > 0) {
            lastNode = lastNode->next;
            n--;
        }
        if (lastNode == nullptr) {
            return head->next;
        }
        lastNode = lastNode->next;
        ListNode* beforeNthNodeFromEnd = head;
        while (lastNode) {
            lastNode = lastNode->next;
            beforeNthNodeFromEnd = beforeNthNodeFromEnd->next;
        }
        beforeNthNodeFromEnd->next = beforeNthNodeFromEnd->next->next;
        return head;
    }
};
