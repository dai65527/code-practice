// https://leetcode.com/problems/container-with-most-water/submissions/983613458/
class Solution {
public:
    int maxArea(vector<int>& height) {
        int ans = 0;
        int i = 0;
        int j = height.size() - 1;
        while (i < j) {
            if (height[i] < height[j]) {
                ans = max(ans, height[i] * (j - i));
                i++;
            } else {
                ans = max(ans, height[j] * (j - i));
                j--;
            }
        }
        return ans;
    }
};
