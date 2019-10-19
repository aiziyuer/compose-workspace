import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

class SolutionTest {

	@Test
	void test() {
		//
		assertEquals(3, new Solution().lengthOfLongestSubstring("abcabcbb"));
		assertEquals(1, new Solution().lengthOfLongestSubstring("bbbbb"));
		assertEquals(3, new Solution().lengthOfLongestSubstring("pwwkew"));

	}

}
