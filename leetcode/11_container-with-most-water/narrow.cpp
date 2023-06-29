// https://leetcode.com/problems/container-with-most-water/submissions/982327316/
class Solution {
public:
    int maxArea(vector<int>& height) {
        int ans = 0;
        int i = 0;
        int j = height.size() - 1;
        while (i < j) {
            if (height[i] < height[j]) {
                ans = max(ans, (j - i) * height[i]);
                i++;
            } else {
                ans = max(ans, (j - i) * height[j]);
                j--;
            }
        }
        return ans;
    }
};
