import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

class SolutionTest {

	@Test
	void test() {
		//
		assertArrayEquals(new int[] { 0, 1 }, new Solution().twoSum(new int[] { 2, 7, 11, 15 }, 9));

		// [3,3]
		assertArrayEquals(new int[] { 0, 1 }, new Solution().twoSum(new int[] { 3, 3 }, 6));

	}

}
