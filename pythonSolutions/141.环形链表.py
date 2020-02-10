#
# @lc app=leetcode.cn id=141 lang=python
#
# [141] 环形链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        # Rude solution
        # nodeSeen = set()
        # while head:
        #     if head in nodeSeen:
        #         return True
        #     else:
        #         nodeSeen.add(head)
        #     head = head.next
            
        # return False
        # Slow and Fast pointer solution
        if head == None or head.next == None:
            return False
        slow = head
        fast = head.next
        while slow != fast:
            if fast == None or fast.next == None:
                return False
            slow = slow.next
            fast = fast.next.next
        return True
        
# @lc code=end

