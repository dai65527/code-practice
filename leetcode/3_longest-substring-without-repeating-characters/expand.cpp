// https://leetcode.com/problems/longest-palindromic-substring/submissions/982284904/
class Solution {
public:
    string longestPalindrome(string s) {
        int start = 0;
        int end = 0;
        int diff;

        for (int i = 0; i < s.size(); i++) {
            // find odd
            diff = expand(s, i, i);
            if (diff > end - start) {
                start = i - diff / 2;
                end = i + diff / 2;
            }
            // find even
            diff = expand(s, i, i+1);
            if (diff > end - start) {
                start = i - diff / 2;
                end = i+1 + diff / 2;
            }
        }
        return s.substr(start, end - start + 1);
    }
    int expand(const string& s, int i, int j) {
        while (i >= 0 && j < s.size() && s[i] == s[j]) {
            i--;
            j++;
        }
        return j - i - 2;
    }
};
