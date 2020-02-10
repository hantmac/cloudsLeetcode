#
# @lc app=leetcode.cn id=11 lang=python
#
# [11] 盛最多水的容器
#

# @lc code=start
class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        max_contain = 0
        aera = 0
        min_height = 0
        i = 0
        j = len(height)-1
        while i < j:
            if  height[i] < height[j]:
                min_height = height[i]
                i +=1
            else:
                min_height = height[j]
                j -=1
            aera = (j -i +1)*min_height
            max_contain = max(aera,max_contain)
                
        return max_contain
        
# @lc code=end

