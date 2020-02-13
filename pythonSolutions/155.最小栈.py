#
# @lc app=leetcode.cn id=155 lang=python
#
# [155] 最小栈
#
# https://leetcode-cn.com/problems/min-stack/description/
#
# algorithms
# Easy (51.31%)
# Likes:    376
# Dislikes: 0
# Total Accepted:    71.5K
# Total Submissions: 139.1K
# Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
#
# 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
# 
# 
# push(x) -- 将元素 x 推入栈中。
# pop() -- 删除栈顶的元素。
# top() -- 获取栈顶元素。
# getMin() -- 检索栈中的最小元素。
# 
# 
# 示例:
# 
# MinStack minStack = new MinStack();
# minStack.push(-2);
# minStack.push(0);
# minStack.push(-3);
# minStack.getMin();   --> 返回 -3.
# minStack.pop();
# minStack.top();      --> 返回 0.
# minStack.getMin();   --> 返回 -2.
# 
# 
#

# @lc code=start
class MinStack(object):

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.data = []
        self.helper = []
        

    def push(self, x):
        """
        :type x: int
        :rtype: None
        """
        self.data.append(x)
        if len(self.helper) == 0 or x < self.helper[-1]:
            self.helper.append(x)
        else:
            self.helper.append(self.helper[-1])
        

    def pop(self):
        """
        :rtype: None
        """
        if self.data:
            self.helper.pop()
            return self.data.pop()
        

    def top(self):
        """
        :rtype: int
        """
        if self.data:
            return self.data[-1]
        

    def getMin(self):
        """
        :rtype: int
        """
        if self.helper:
            return self.helper[-1]
        


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
# @lc code=end

