#
# @lc app=leetcode.cn id=70 lang=python
#
# [70] 爬楼梯
#

# @lc code=start
class Solution(object):
    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        a =[0]
        a.append(1)
        a.append(2)
        res = 0
        if n <=2:
            return n
        for i in range(3,n+1):
            a.append(a[i-1] + a[i-2])
            res = a[i]
            
        return res
# @lc code=end

