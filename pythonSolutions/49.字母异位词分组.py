#
# @lc app=leetcode.cn id=49 lang=python
#
# [49] 字母异位词分组
#
# https://leetcode-cn.com/problems/group-anagrams/description/
#
# algorithms
# Medium (60.53%)
# Likes:    275
# Dislikes: 0
# Total Accepted:    53K
# Total Submissions: 86.9K
# Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
#
# 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
# 
# 示例:
# 
# 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
# 输出:
# [
# ⁠ ["ate","eat","tea"],
# ⁠ ["nat","tan"],
# ⁠ ["bat"]
# ]
# 
# 说明：
# 
# 
# 所有输入均为小写字母。
# 不考虑答案输出的顺序。
# 
# 
#

# @lc code=start
class Solution(object):
    def groupAnagrams(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """
        # 暴力解法
        # K:v k = tuple(), v= list[]
        import collections
        res = collections.defaultdict(list)
        for s in strs:
            res[tuple(sorted(s))].append(s)
            
        return res.values()
# @lc code=end

