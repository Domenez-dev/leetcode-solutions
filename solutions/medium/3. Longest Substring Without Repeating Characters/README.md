Problem: Longest Substring Without Repeating Characters  
Link: https://leetcode.com/problems/longest-substring-without-repeating-characters/

Description:  
Given a string `s`, find the length of the longest substring without repeating characters.

Complexity:

- Time: O(n^2), where n is the length of the string. The `noRepeatCheck` function performs nested iterations, making the solution less optimal.
- Space: O(k), where k is the length of the substring being checked.

- Runtime: 56ms (Beats 18%)
