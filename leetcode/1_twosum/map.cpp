// https://leetcode.com/problems/two-sum/submissions/980054590/
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> mp;
        for (int i = 0; i < nums.size(); i++) {
            if (mp.count(target - nums[i])) {
                return vector<int>{i, mp[target - nums[i]]};
            }
            mp[nums[i]] = i;
        }
        return vector<int>{0, 0};
    }
};
