// https://leetcode.com/problems/merge-two-sorted-lists/submissions/986540787/
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
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        if (!list1) {
            return list2;
        }
        if (!list2) {
            return list1;
        }
        ListNode* head = list1;
        ListNode* node = list1;
        ListNode* wait = list2;
        if (list1->val > list2->val) {
            head = list2;
            node = list2;
            wait = list1;
        }
        while (node->next && wait) {
            if (node->next->val > wait->val) {
                ListNode* tmp = node->next;
                node->next = wait;
                wait = tmp;
            } else {
                node = node->next;
            }
        }
        if (!node->next) {
            node->next = wait;
        }
        return head;
    }
};
