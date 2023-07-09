class Solution:
    def isPalindrome(self, x: int) -> bool:
        x_str = str(x)
        word = ""
        for i in range(len(x_str)-1, -1, -1):
            word += x_str[i]
        return word == x_str