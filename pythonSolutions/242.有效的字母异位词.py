#
# @lc app=leetcode.cn id=242 lang=python
#
# [242] 有效的字母异位词
#
# https://leetcode-cn.com/problems/valid-anagram/description/
#
# algorithms
# Easy (58.23%)
# Likes:    160
# Dislikes: 0
# Total Accepted:    77.4K
# Total Submissions: 132.1K
# Testcase Example:  '"anagram"\n"nagaram"'
#
# 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
# 
# 示例 1:
# 
# 输入: s = "anagram", t = "nagaram"
# 输出: true
# 
# 
# 示例 2:
# 
# 输入: s = "rat", t = "car"
# 输出: false
# 
# 说明:
# 你可以假设字符串只包含小写字母。
# 
# 进阶:
# 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
# 
#

# @lc code=start
class Solution(object):
    def isAnagram(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        # 暴力求解
        # if sorted(s) == sorted(t):
        #     return True
        # else:
        #     return False
        # 哈希表
        import collections
        if len(s) != len(t):
            return False
        d = collections.defaultdict(int)
        for i in range(len(s)):
            d[s[i]] +=1
            d[t[i]] -= 1
        
        for j in d.values():
            if j != 0:
                return False
        return True
        
        
        
# @lc code=end

