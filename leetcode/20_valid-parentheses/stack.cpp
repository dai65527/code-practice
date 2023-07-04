// https://leetcode.com/problems/valid-parentheses/submissions/986517956/
class Solution {
 public:
  bool isValid(string s) {
    stack<char> expects;
    for (int i = 0; i < s.size(); i++) {
      if (s[i] == '(') {
        expects.push(')');
      } else if (s[i] == '[') {
        expects.push(']');
      } else if (s[i] == '{') {
        expects.push('}');
      } else if (!expects.empty() && s[i] == expects.top()) {
        expects.pop();
      } else {
        return false;
      }
    }
    return expects.empty();
  }
};
