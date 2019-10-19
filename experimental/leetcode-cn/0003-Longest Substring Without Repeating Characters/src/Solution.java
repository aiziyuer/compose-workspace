import java.util.BitSet;

/**
 * 
 * Given a string, find the length of the longest substring without repeating
 * characters.
 * 
 * Example 1:
 * 
 * Input: "abcabcbb" Output: 3 Explanation: The answer is "abc", with the length
 * of 3.
 * 
 * Example 2:
 * 
 * Input: "bbbbb" Output: 1 Explanation: The answer is "b", with the length of
 * 1.
 * 
 * Example 3:
 * 
 * Input: "pwwkew" Output: 3 Explanation: The answer is "wke", with the length
 * of 3. Note that the answer must be a substring, "pwke" is a subsequence and
 * not a substring.
 * 
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。<br>
 * 
 * 示例 1: <br>
 * 输入: "abcabcbb"<br>
 * 输出: 3 <br>
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * 
 * 示例 2: <br>
 * 输入: "bbbbb" <br>
 * 输出: 1 <br>
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * 
 * 示例 3: <br>
 * 输入: "pwwkew" <br>
 * 输出: 3 <br>
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 * 
 */

class Solution {
	public int lengthOfLongestSubstring(String s) {

		BitSet checker = new BitSet(256);
		char[] charArrays = s.toCharArray();
		int dumplicateNum = 0;
		for (int i = 0; i < charArrays.length; i++) {

			checker.clear();
			int j = 0;
			for (j = i; j < charArrays.length; j++) {
				char c = charArrays[j];
				if (checker.get(c))
					break;

				checker.set(c);
			}

			dumplicateNum = Math.max(dumplicateNum, j - i);
		}

		return dumplicateNum;
	}
}