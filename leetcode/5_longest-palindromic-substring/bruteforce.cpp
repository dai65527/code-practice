// TLE :(
class Solution {
public:
    string longestPalindrome(string s) {
        string ans = "";
        for (int i = 0; i < s.size(); i++) {
            for (int j = i + 1; j <= s.size(); j++) {
                if (j - i <= ans.size()) {
                    continue;
                }
                string sub = s.substr(i, j - i);
                if (isPalindrome(sub)) {
                    ans = sub;
                }
            }
        }
        return ans;
    }

    bool isPalindrome(string& s) {
        int i = 0;
        int j = s.size() - 1;
        while (i < j) {
            if (s[i] != s[j]) {
                return false;
            }
            i++;
            j--;
        }
        return true;
    }
};
