// https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/980424573/
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int ans = 0;
        unordered_map<char, int> charIdxMap;
        int left = 0;
        for (int right = 0; right < s.size(); right++) {
            if (!charIdxMap.count(s[right]) || charIdxMap[s[right]] < left) {
                ans = max(ans, right - left + 1);
            } else {
                left = charIdxMap[s[right]] + 1;
            }
            charIdxMap[s[right]] = right;
        }
        return ans;
    }
};
