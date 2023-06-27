// https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int ans = 0;
        unordered_set<char> char_set;
        int left = 0;
        for (int right = 0; right < s.size(); right++) {
            if (!char_set.count(s[right])) {
                char_set.insert(s[right]);
                ans = max(ans, right - left + 1);
            } else {
                while (s[left] != s[right]) {
                    char_set.erase(s[left]);
                    left++;
                }
                left++;
            }
        }
        return ans;
    }
};
