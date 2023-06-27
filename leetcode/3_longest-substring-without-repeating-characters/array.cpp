// https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/980433576/
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int ans = 0;
        vector<int> charIdx(127, -1);
        int left = 0;
        for (int right = 0; right < s.size(); right++) {
            if (charIdx[s[right]] < left) {
                ans = max(ans, right - left + 1);
            } else {
                left = charIdx[s[right]] + 1;
            }
            charIdx[s[right]] = right;
        }
        return ans;
    }
};
