#
# @lc app=leetcode.cn id=239 lang=python
#
# [239] 滑动窗口最大值
#
# https://leetcode-cn.com/problems/sliding-window-maximum/description/
#
# algorithms
# Hard (43.65%)
# Likes:    235
# Dislikes: 0
# Total Accepted:    28K
# Total Submissions: 64.1K
# Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
#
# 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
# 个数字。滑动窗口每次只向右移动一位。
# 
# 返回滑动窗口中的最大值。
# 
# 
# 
# 示例:
# 
# 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
# 输出: [3,3,5,5,6,7] 
# 解释: 
# 
# ⁠ 滑动窗口的位置                最大值
# ---------------               -----
# [1  3  -1] -3  5  3  6  7       3
# ⁠1 [3  -1  -3] 5  3  6  7       3
# ⁠1  3 [-1  -3  5] 3  6  7       5
# ⁠1  3  -1 [-3  5  3] 6  7       5
# ⁠1  3  -1  -3 [5  3  6] 7       6
# ⁠1  3  -1  -3  5 [3  6  7]      7
# 
# 
# 
# 提示：
# 
# 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
# 
# 
# 
# 进阶：
# 
# 你能在线性时间复杂度内解决此题吗？
# 
#

# @lc code=start
from collections import deque
class Solution(object):
    def maxSlidingWindow(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        # 暴露求解，时间复杂度为 o(n*k)
        # res = []
        # if len(nums) == 0:
        #     return res
        # tmp_max = max(nums[0:k])
        # for i in range(len(nums) -k+1):
        #     tmp_max = max(nums[i:i+k])
        #     res.append(tmp_max)
        # return res
        # 双端队列法，时间复杂度为on(n+k)常数级
        res = []
        deq = deque()
        if len(nums) == 0:
            return res
        if k == 1:
            return nums
        
        def clean_deq(i):
            if deq and deq[0] == i-k:
                deq.popleft()
            while deq and nums[i]>nums[deq[-1]]:
                deq.pop()
        
        max_index = 0
        for i in range(k):
            clean_deq(i)
            deq.append(i)
            if nums[i] > nums[max_index]:
                max_index = i
        res.append(nums[max_index])
        for i in range(k,len(nums)):
            clean_deq(i)
            deq.append(i)
            res.append(nums[deq[0]])
            
        return res
            
         
# @lc code=end

