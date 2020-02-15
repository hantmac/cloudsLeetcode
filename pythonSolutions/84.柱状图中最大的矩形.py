#
# @lc app=leetcode.cn id=84 lang=python
#
# [84] 柱状图中最大的矩形
#
# https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
#
# algorithms
# Hard (38.59%)
# Likes:    435
# Dislikes: 0
# Total Accepted:    27.4K
# Total Submissions: 70.7K
# Testcase Example:  '[2,1,5,6,2,3]'
#
# 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
# 
# 求在该柱状图中，能够勾勒出来的矩形的最大面积。
# 
# 
# 
# 
# 
# 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
# 
# 
# 
# 
# 
# 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
# 
# 
# 
# 示例:
# 
# 输入: [2,1,5,6,2,3]
# 输出: 10
# 
#

# @lc code=start
class Solution(object):
    def largestRectangleArea(self, heights):
        """
        :type heights: List[int]
        :rtype: int
        """
        # 暴力解法
        # aera = 0
        # for i in range(len(heights)):
        #     min_height = 10000000000
        #     for j in range(i,len(heights)):
        #             min_height = min(min_height,heights[j])
        #             aera = max(aera,min_height*(j-i+1))
        # return aera
        s= []
        max_aera = 0
        s.append(-1)
        for i in range(len(heights)):
            while s[-1] != -1 and heights[s[-1]] > heights[i]:
                max_aera = max(max_aera,heights[s.pop()]*(i-s[-1]-1)) 
            s.append(i)
        while s[-1] != -1:
            max_aera = max(max_aera,heights[s.pop()]*(len(heights)-s[-1]-1))
        
        return max_aera
# @lc code=end

