// https://leetcode.com/problems/3sum/submissions/983756652/
class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int> > ans;
        sort(nums.begin(), nums.end());
        int i = 0;
        while (i < nums.size()) {
            int j = i+1;
            int k = nums.size()-1;
            while (j < k) {
                if (nums[i] + nums[j] + nums[k] == 0) {
                    ans.push_back(vector<int>{nums[i], nums[j], nums[k]});
                    do {
                        j++;
                    } while (j < nums.size() && nums[j-1] == nums[j]);
                    do {
                        k--;
                    } while (k >= 0 && nums[k] == nums[k+1]);
                } else if (nums[i] + nums[j] + nums[k] < 0) {
                    j++;
                } else {
                    k--;
                }
            }
            do {
                i++;
            } while (i < nums.size() && nums[i-1] == nums[i]);
        }
        return ans;
    }
};
