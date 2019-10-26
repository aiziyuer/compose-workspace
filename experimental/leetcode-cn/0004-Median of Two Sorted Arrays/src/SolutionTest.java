import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

class SolutionTest {

	@Test
	void test() {
		//
		assertEquals(2.0, new Solution().findMedianSortedArrays(new int[] { 1, 3 }, new int[] { 2 }));
		assertEquals(2.5, new Solution().findMedianSortedArrays(new int[] { 1, 2 }, new int[] { 2, 3, 4, 5 }));
		assertEquals(3.0, new Solution().findMedianSortedArrays(new int[] { 1, 2 }, new int[] { 3, 4, 5 }));
		assertEquals(3.0, new Solution().findMedianSortedArrays(new int[] { 3, 4, 5 }, new int[] { 1, 2 }));
	}

}
