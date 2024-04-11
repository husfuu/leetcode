class Solution(object):
    def plusOne(self, digits):
        """ljjkf
        :type digits: List[int]
        :rtype: List[int]
        """
        result = digits
        for idx, item in reversed(list(enumerate(digits))):
            if item != 9:
                result[idx] = item + 1
                break
            elif idx == 0 and item == 9:
                result[idx] = 0
                result.insert(0, 1)
            elif item == 9:
                    result[idx] = 0

        return result 
