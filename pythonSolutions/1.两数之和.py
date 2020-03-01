#
# @lc app=leetcode.cn id=1 lang=python
#
# [1] 两数之和
#
# https://leetcode-cn.com/problems/two-sum/description/
#
# algorithms
# Easy (47.52%)
# Likes:    7668
# Dislikes: 0
# Total Accepted:    826.3K
# Total Submissions: 1.7M
# Testcase Example:  '[2,7,11,15]\n9'
#
# 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
#
# 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
#
# 示例:
#
# 给定 nums = [2, 7, 11, 15], target = 9
#
# 因为 nums[0] + nums[1] = 2 + 7 = 9
# 所以返回 [0, 1]
#
#
#

# @lc code=start


class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        # 暴力解法
        # tmp = []
        # size = len(nums)
        # for i in range(size-1):
        #     for j in range(i+1, size):
        #         if nums[i] + nums[j] == target:
        #             tmp.append(i)
        #             tmp.append(j)
        #             return tmp
        # 哈希表, 两遍哈希
        # def get_keys(d, value):
        #     for k, v in d.items():
        #         if v == value:
        #             return k
        # import collections
        # d = collections.defaultdict(int)
        # for i, v in enumerate(nums):
        #     d[i] = v
        # for k , v in d.items():
        #     component = target - v
        #     if component in nums and get_keys(d,component)!=k:
        #         return [k,get_keys(d,component)]
        # 一遍哈希，在生成map的时候顺便检查下
        d = {}
        for i in range(len(nums)):
            if nums[i] in d:
                return [d[nums[i]],i]
            d[target-nums[i]] = i
        return False
# @lc code=end
