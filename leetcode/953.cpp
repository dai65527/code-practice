// https://leetcode.com/problems/verifying-an-alien-dictionary/
class Solution {
public:
    bool isAlienSorted(vector<string>& words, string order) {
        int dict[128];
        int l = words.size();
        
        if (l <= 1) {
            return true;
        }
        
        // make dictionaly
        for (int i = 0; i < 26; i++) {
            dict[order[i]] = i;
        }
        
        // for all words
        for (int i = 1; i < l; i++) {
            if (alienStrcmp(words[i - 1], words[i], dict) > 0) {
                return false;
            }
        }
        return true;
    }
    
    int alienStrcmp(const string& s1, const string& s2, int* dict) {
        int l1 = s1.size();
        int l2 = s2.size();
        for (int i = 0; i < l1 && i < l2; i++) {
            if (s1[i] != s2[i]) {
                return dict[s1[i]] - dict[s2[i]];
            }
        }
        return l1 - l2;
    }
};
