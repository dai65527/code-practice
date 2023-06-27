// https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/980073579/
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int ans = 0;
        string sub = "";
        for (int i = 0; i < s.size(); i++) {
            size_t foundIndex = sub.find(s[i]);
            if (foundIndex == string::npos) {
                sub.push_back(s[i]);
                if (sub.size() > ans) {
                    ans = sub.size();
                }
            } else {
                cout << sub << ", " << foundIndex << endl;
                sub = sub.substr(foundIndex+1);
                sub.push_back(s[i]);
            }
        }
        return ans;
    }
};
