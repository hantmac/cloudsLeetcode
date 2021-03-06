#
# @lc app=leetcode.cn id=20 lang=python
#
# [20] 有效的括号
#
# https://leetcode-cn.com/problems/valid-parentheses/description/
#
# algorithms
# Easy (40.75%)
# Likes:    1378
# Dislikes: 0
# Total Accepted:    200.8K
# Total Submissions: 492K
# Testcase Example:  '"()"'
#
# 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
# 
# 有效字符串需满足：
# 
# 
# 左括号必须用相同类型的右括号闭合。
# 左括号必须以正确的顺序闭合。
# 
# 
# 注意空字符串可被认为是有效字符串。
# 
# 示例 1:
# 
# 输入: "()"
# 输出: true
# 
# 
# 示例 2:
# 
# 输入: "()[]{}"
# 输出: true
# 
# 
# 示例 3:
# 
# 输入: "(]"
# 输出: false
# 
# 
# 示例 4:
# 
# 输入: "([)]"
# 输出: false
# 
# 
# 示例 5:
# 
# 输入: "{[]}"
# 输出: true
# 
#

# @lc code=start
class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        stack= []
        if s == "":
            return True
        if len(s)%2 !=0:
            return False
        if s[0] == "}" or s[0] == ")" or s[0] == "]":
            return False
        for i in s:
            if i == "(" or i == "{" or i == "[":
                stack.append(i)
            elif i == ")":
                tmp = stack.pop()
                if tmp != "(":
                    return False
            elif i == "}":
                tmp = stack.pop()
                if tmp != "{":
                    return False
            elif i == "]":
                tmp = stack.pop()
                if tmp !=  "[":
                    return False
        if len(stack) != 0:
            return False
        return True
        
# @lc code=end

