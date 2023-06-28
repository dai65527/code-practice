// https://leetcode.com/problems/longest-palindromic-substring/submissions/980866638/
// This is much faster than bruteforce.cpp (which was TLE) without using substr in the loop.
class Solution {
public:
    string longestPalindrome(string s) {
        int startIdx = 0;
        int length = 0;
        for (int i = 0; i < s.size(); i++) {
            for (int j = i; j < s.size(); j++) {
                if (length < j - i + 1 && isPalindrome(s, i, j)) {
                    startIdx = i;
                    length = j - i + 1;
                }
            }
        }
        return s.substr(startIdx, length);
    }
    bool isPalindrome(string &s, int startIdx, int endIdx) {
        while (startIdx < endIdx) {
            if (s[startIdx] != s[endIdx]) {
                return false;
            }
            startIdx++;
            endIdx--;
        }
        return true;
    }
};
