/**
 * 
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 * 
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 * 
 * You may assume nums1 and nums2 cannot be both empty.
 * 
 * Example 1:
 * 
 * nums1 = [1, 3] nums2 = [2]
 * 
 * The median is 2.0
 * 
 * Example 2:
 * 
 * nums1 = [1, 2] nums2 = [3, 4]
 * 
 * The median is (2 + 3)/2 = 2.5
 * 
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 * 
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 * 
 * 你可以假设 nums1 和 nums2 不会同时为空。
 * 
 * 示例 1:
 * 
 * nums1 = [1, 3] nums2 = [2]
 * 
 * 则中位数是 2.0
 * 
 * 示例 2:
 * 
 * nums1 = [1, 2] nums2 = [3, 4]
 * 
 * 则中位数是 (2 + 3)/2 = 2.5
 * 
 */

class Solution {
	public double findMedianSortedArrays(int[] nums1, int[] nums2) {

		// 设置边界
		int[] fakeNums1 = new int[nums1.length + 1];
		int[] fakeNums2 = new int[nums2.length + 1];
		for (int i = 0; i < nums1.length; i++)
			fakeNums1[i] = nums1[i];
		for (int i = 0; i < nums2.length; i++)
			fakeNums2[i] = nums2[i];
		fakeNums1[nums1.length] = fakeNums2[nums2.length] = Integer.MAX_VALUE;

		// 分别记录中位数的位置
		int i1, i2;
		i1 = i2 = 0;

		// 定义中位数
		int mid = -1;

		// 取第n小的数字
		for (int c = 0; c <= (nums1.length + nums2.length - 1) / 2; c++) {

			if (fakeNums1[i1] <= fakeNums2[i2]) {
				mid = fakeNums1[i1];
				if (i1 < fakeNums1.length)
					i1++;
			} else {
				mid = fakeNums2[i2];
				if (i2 < fakeNums2.length)
					i2++;
			}

		}

		// 奇数中位数就一个, 偶数需要再往后看一位数字然后算下均值
		return ((nums1.length + nums2.length) % 2 != 0) ? mid : (mid + Math.min(fakeNums1[i1], fakeNums2[i2])) / 2.0;

	}
}